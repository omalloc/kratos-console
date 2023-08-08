package data

import (
	"context"
	"github.com/omalloc/contrib/kratos/orm"
	"github.com/omalloc/kratos-console/internal/biz"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/omalloc/kratos-console/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewZoneRepo,
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
	logFilter := log.NewFilter(logger, log.FilterLevel(log.LevelDebug))
	dbLog := log.NewHelper(logFilter)

	db, err := orm.New(
		orm.WithDriver(mysql.Open(c.Database.Source)),
		orm.WithTracing(),
		orm.WithLogger(
			orm.WithDebug(),
			orm.WIthSlowThreshold(time.Second*2),
			orm.WithSkipCallerLookup(true),
			orm.WithSkipErrRecordNotFound(true),
			orm.WithLogHelper(logFilter),
		),
	)

	if err != nil {
		return nil, emptyCallback, err
	}

	_ = db.AutoMigrate(&biz.Zone{})

	cleanup := func() {
		dbLog.Info("closing the data resources")
	}

	return &Data{
		db: db,
	}, cleanup, nil
}

func (d *Data) Check(ctx context.Context) error {
	if c, err := d.db.DB(); err == nil {
		return c.PingContext(ctx)
	}
	return nil
}
