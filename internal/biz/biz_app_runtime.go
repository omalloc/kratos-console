package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/contrib/kratos/orm"
	"github.com/omalloc/contrib/kratos/orm/crud"
)

type AppRuntime struct {
	ID        int64             `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:主键ID"`
	Name      string            `json:"name" gorm:"column:name;type:varchar(64);not null;comment:应用名称"`
	Hostname  string            `json:"hostname" gorm:"column:hostname;type:varchar(64);not null;comment:服务主机"`
	Key       string            `json:"key" gorm:"column:key;type:varchar(128);not null;comment:服务唯一KEY(/命名空间/服务名/服务主机)"`
	Version   string            `json:"version" gorm:"column:version;type:varchar(32);null;comment:服务运行版本"`
	Endpoints []string          `json:"endpoints" gorm:"column:endpoints;type:json;serializer:json;not null;comment:服务注册端点([]scheme://IP:Port)"`
	Metadata  map[string]string `json:"metadata" gorm:"column:metadata;type:json;serializer:json;not null;comment:服务元数据"`
	Cluster   string            `json:"cluster" gorm:"column:cluster;type:varchar(32);not null;comment:运行集群"`
	Namespace string            `json:"namespace" gorm:"column:namespace;type:varchar(32);not null;comment:运行时命名空间"`

	orm.DBModel
}

func (AppRuntime) TableName() string {
	return "apps_runtime"
}

type AppRuntimeRepo interface {
	crud.CRUD[AppRuntime]

	SelectListByNamespace(context.Context, string) ([]*AppRuntime, error)
}

type AppRuntimeUsecase struct {
	log  *log.Helper
	repo AppRuntimeRepo
}

func NewAppRuntimeUsecase(logger log.Logger, repo AppRuntimeRepo) *AppRuntimeUsecase {
	return &AppRuntimeUsecase{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (uc *AppRuntimeUsecase) Test(ctx context.Context) error {
	ret, err := uc.repo.SelectOne(ctx, 1)

	uc.log.WithContext(ctx).Infof("ret: %#+v", ret)
	return err
}
