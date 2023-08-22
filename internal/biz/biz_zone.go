package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/omalloc/kratos-console/api/console/resource"
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
	Status     int    `json:"status" gorm:"column:status;type:tinyint(1);default:1;comment:'1:正常;2:禁用'"`

	DBModel
}

type ZoneRepo interface {
	GetZoneList(ctx context.Context, query *QueryPager) ([]*Zone, int64, error)
	GetZoneByID(context.Context, int64) (*Zone, error)
	CreateZone(context.Context, *Zone) error
	UpdateZone(context.Context, *Zone) error
	DisableZone(context.Context, int64) error
	DeleteZone(context.Context, int64) error
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

func (uc *ZoneUsecase) GetZoneList(ctx context.Context, query *QueryPager) ([]*Zone, int64, error) {
	return uc.repo.GetZoneList(ctx, query)
}

func (uc *ZoneUsecase) GetZone(ctx context.Context, ID int64) (*Zone, error) {
	return uc.repo.GetZoneByID(ctx, ID)
}

func (uc *ZoneUsecase) CreateZone(ctx context.Context, info *pb.CreateZoneRequest) (*Zone, error) {
	zone := &Zone{
		Name:       info.Name,
		Code:       info.Code,
		RegionName: info.RegionName,
		RegionCode: info.RegionCode,
		Env:        info.Env,
	}
	if err := uc.repo.CreateZone(ctx, zone); err != nil {
		return nil, err
	}

	return zone, nil
}

func (uc *ZoneUsecase) DeleteZone(ctx context.Context, ID int64) error {
	return uc.repo.DeleteZone(ctx, ID)
}
