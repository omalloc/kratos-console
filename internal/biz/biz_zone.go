package biz

import (
	"context"
	"github.com/omalloc/contrib/protobuf"
	pb "github.com/omalloc/kratos-console/api/console/resource"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/omalloc/contrib/kratos/orm"
)

type EnvType string

const (
	EnvDev  EnvType = "dev"
	EnvTest EnvType = "test"
	EnvQA   EnvType = "qa"
	EnvProd EnvType = "prod"
)

type ZoneQuery struct {
	*protobuf.Pagination

	Env string `json:"env"`
}

type Zone struct {
	ID         int64  `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"column:name;type:varchar(32);comment:可用区名称"`
	Code       string `json:"code" gorm:"column:code;type:varchar(32);comment:可用区代码"`
	RegionName string `json:"region_name" gorm:"column:region_name;type:varchar(32);comment:地区名称"`
	RegionCode string `json:"region_code" gorm:"column:region_code;type:varchar(32);comment:地区代码"`
	Env        string `json:"env" gorm:"column:env;type:varchar(12);comment:环境变量dev,test,qa,prod,etc.."`
	Status     int    `json:"status" gorm:"column:status;type:tinyint(1);default:1;comment:1=正常,2=禁用"`

	orm.DBModel
}

type ZoneRepo interface {
	GetZoneList(ctx context.Context, query *ZoneQuery) ([]*Zone, error)
	GetZoneByID(context.Context, int64) (*Zone, error)
	SelectSimpleList(context.Context) ([]*Zone, error)
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

func (uc *ZoneUsecase) GetZoneList(ctx context.Context, query *ZoneQuery) ([]*Zone, error) {
	return uc.repo.GetZoneList(ctx, query)
}

func (uc *ZoneUsecase) GetSimpleList(ctx context.Context) ([]*Zone, error) {
	return uc.repo.SelectSimpleList(ctx)
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

func (uc *ZoneUsecase) UpdateZone(ctx context.Context, req *pb.UpdateZoneRequest) error {
	zone := &Zone{
		ID:         int64(req.Id),
		Name:       req.Name,
		Code:       req.Code,
		RegionName: req.RegionName,
		RegionCode: req.RegionCode,
		Env:        req.Env,
	}
	return uc.repo.UpdateZone(ctx, zone)
}

func (uc *ZoneUsecase) DeleteZone(ctx context.Context, ID int64) error {
	return uc.repo.DeleteZone(ctx, ID)
}
