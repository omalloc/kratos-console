package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type App struct {
	DBModel

	ID          int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:应用ID"`
	Name        string `json:"name" gorm:"column:name;type:varchar(64);not null;comment:应用名称"`
	Description string `json:"description" gorm:"column:description;type:varchar(256);not null;comment:应用描述"`
	Icon        string `json:"icon" gorm:"column:icon;type:varchar(256);not null;comment:应用图标"`
}

type AppRepo interface {
	List(ctx context.Context, query *QueryPager) ([]*App, int64, error)
}

type AppUsecase struct {
	repo AppRepo

	log *log.Helper
}

func NewAppUsecase(logger log.Logger, repo AppRepo) *AppUsecase {
	return &AppUsecase{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (uc *AppUsecase) GetAppList(ctx context.Context, pager *QueryPager) ([]*App, int64, error) {
	return uc.repo.List(ctx, pager)
}
