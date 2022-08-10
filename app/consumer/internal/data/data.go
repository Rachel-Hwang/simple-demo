package data

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	provider "helloworld/api/provider/service/v1"
	"helloworld/app/consumer/internal/biz"
	"helloworld/app/consumer/internal/conf"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	// 1. init mysql driver
	//_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProviderClient)

// Data .
type Data struct {
	//db *ent.Client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger,db *gorm.DB) (*Data, func(), error) {
	return &Data{db: db}, func() {
		log.NewHelper(logger).Info("close the connect")
	},nil
}

// NewDB gorm client
func NewDB(c *conf.Data) (*gorm.DB) {
	db,err := gorm.Open(mysql.New(mysql.Config{
		DSN: c.Database.Source,
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}),&gorm.Config{
		DisableForeignKeyConstraintWhenMigrating:true,
	})
	if err != nil {
		panic("fail to connect mysql")
	}
	return db
}

// NewDB ent Client
//func NewDB(c *conf.Data, logger log.Logger) (*Data, func(), error) {
//	log := log.NewHelper(logger)
//	drv,err := sql.Open(c.Database.Driver,c.Database.Source)
//	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
//		log.WithContext(ctx).Info(i...)
//
//		tracer := otel.Tracer("ent.")
//		kind := trace.SpanKindServer
//		_,span := tracer.Start(
//			ctx,
//			"Query",
//			trace.WithAttributes(attribute.String("sql",fmt.Sprint(i...))),
//			trace.WithSpanKind(kind),
//			)
//		span.End()
//	})
//	client := ent.NewClient(ent.Driver(sqlDrv))
//	if err != nil {
//		log.Errorf("fail open conn",err)
//		return nil, nil, err
//	}
//	if err := client.Schema.Create(context.Background());err != nil {
//		log.Errorf("failed creating schema resources: %v", err)
//		return nil, nil, err
//	}
//	d := &Data{
//		db: client,
//	}
//	return d, func() {
//		log.Info("message", "closing the data resources")
//		if err := d.db.Close();err != nil {
//			log.Error(err)
//		}
//	}, nil
//}

// NewProviderClient .
func NewProviderClient(data *Data, logger log.Logger) biz.ConsumerRepo {
	// new consul client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	reg := consul.New(client)

	conn,err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint("discovery://provider"),
		grpc.WithDiscovery(reg))

	return &consumerRepo{
		data: data,
		log:  log.NewHelper(logger),
		providerSrv: provider.NewProviderClient(conn),
	}
}