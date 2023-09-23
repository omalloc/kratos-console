package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/omalloc/contrib/kratos/orm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/omalloc/kratos-console/internal/biz"
	"github.com/omalloc/kratos-console/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewZoneRepo,
	NewNodeRepo,
	NewAppRepo,
	NewNamespaceRepo,
	NewAppRuntimeRepo,
)

var (
	emptyCallback = func() {}
)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	db, err := orm.New(
		orm.WithDriver(mysql.Open(c.Database.Source)),
		orm.WithTracingOpts(orm.WithDatabaseName("kratos_cp")),
		orm.WithLogger(
			orm.WithDebug(),
			orm.WIthSlowThreshold(time.Second*2),
			orm.WithSkipCallerLookup(false),
			orm.WithSkipErrRecordNotFound(true),
			orm.WithLogHelper(logger),
		),
	)

	if err != nil {
		return nil, emptyCallback, err
	}

	_ = db.Session(&gorm.Session{SkipHooks: true}).
		AutoMigrate(
			&biz.Zone{},
			&biz.Node{},
			&biz.App{},
			&biz.Namespace{},
			&biz.AppRuntime{},
		)

	cleanup := func() {
		log.Info("closing the data resources")
	}

	return &Data{
		db: db,
	}, cleanup, nil
}

func (d *Data) Check(ctx context.Context) error {
	// check database connection
	// but skip gorm hooks (tracing and more...)
	tx := d.db.Session(&gorm.Session{SkipHooks: true})
	if c, err := tx.DB(); err == nil {
		return c.PingContext(ctx)
	}
	return nil
}
