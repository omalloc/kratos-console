package data_test

import (
	"context"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/kratos-console/internal/conf"
	"github.com/omalloc/kratos-console/internal/data"
	"testing"
	"time"
)

func loadConfig() *conf.Bootstrap {
	c := config.New(
		config.WithSource(
			file.NewSource("../../configs"),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	return &bc
}

func TestNewAppRuntimeRepo(t *testing.T) {
	c := loadConfig()
	ds, cleanup, _ := data.NewData(c.Data, log.GetLogger())
	defer cleanup()

	repo := data.NewAppRuntimeRepo(ds)

	t.Run("Test repo.SelectOne", func(tt *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		one, err := repo.SelectOne(ctx, 1)
		if err != nil {
			tt.Error(err)
		}
		tt.Logf("SelectOne Data: %#+v", one)
	})

}
