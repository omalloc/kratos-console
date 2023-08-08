package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type EnvType string

const (
	EnvDev  EnvType = "dev"
	EnvTest EnvType = "test"
	EnvQA   EnvType = "qa"
	EnvProd EnvType = "prod"
)

type Zone struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"column:name;type:varchar(32);"`
	Code       string `json:"code" gorm:"column:code;type:varchar(32);"`
	RegionName string `json:"region_name" gorm:"column:region_name;type:varchar(32);"`
	RegionCode string `json:"region_code" gorm:"column:region_code;type:varchar(32);"`
	Env        string `json:"env" gorm:"column:env;type:varchar(12);"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Filter struct {
	Keyword string `json:"keyword"`
}

type ZoneRepo interface {
	GetZoneList(ctx context.Context) ([]*Zone, int64, error)
}

type ZoneUsecase struct {
	repo ZoneRepo
	log  *log.Helper
}

func NewZoneUsecase(repo ZoneRepo, logger log.Logger) *ZoneUsecase {
	return &ZoneUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ZoneUsecase) GetZoneList(ctx context.Context, q *QueryPager) ([]*Zone, int64, error) {
	return uc.repo.GetZoneList(ctx)
}
