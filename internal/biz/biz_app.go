package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/omalloc/kratos-console/api/types"
)

var (
	ErrAppNotFound  = errors.New("应用不存在")
	ErrInvalidAppId = errors.New("无效的应用ID")
)

type Port struct {
	Port     uint   `json:"port"`
	Protocol string `json:"protocol,omitempty"`
	Remark   string `json:"remark,omitempty"`
}

type App struct {
	ID          int64    `json:"id" gorm:"column:id;primaryKey;autoIncrement;comment:应用ID"`
	Name        string   `json:"name" gorm:"column:name;type:varchar(64);not null;comment:应用名称"`
	Alias       string   `json:"alias" gorm:"column:alias;type:varchar(64);comment:应用别名,有别名优先展示别名"`
	Type        int      `json:"type" gorm:"column:type;type:int(11);not null;comment:应用类型(1-系统应用,2-WEB应用,3-API应用,4-子服务,0,5-未定义)"`
	Ports       []Port   `json:"ports" gorm:"column:ports;type:json;serializer:json;comment:应用预期端口,使用0端口表示随机分配"`
	Users       []string `json:"users" gorm:"column:users;type:json;serializer:json;comment:应用负责人"`
	NamespaceID int64    `json:"namespace_id" gorm:"column:namespace_id;type:bigint(20);not null;comment:应用所属命名空间ID"`
	Repos       []string `json:"repos" gorm:"-"`
	Description string   `json:"description" gorm:"column:description;type:varchar(256);comment:应用描述"`
	Icon        string   `json:"icon" gorm:"column:icon;type:varchar(256);default null;comment:应用图标"`

	types.DBModel
}

type NamespaceApp struct {
	App

	NamespaceName  string `json:"namespace_name" gorm:"column:namespace_name;type:varchar(32);comment:命名空间唯一名称"`
	NamespaceAlias string `json:"namespace_alias" gorm:"column:namespace_alias;type:varchar(32);comment:命名空间别称"`
}

type AppRepo interface {
	List(ctx context.Context, pagination *types.Pagination) ([]*App, error)
	SelectByNamespace(context.Context, string) ([]*NamespaceApp, error)
	Create(ctx context.Context, app *App) error
	Update(ctx context.Context, app *App) error
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

func (uc *AppUsecase) GetAppList(ctx context.Context, pagination *types.Pagination) ([]*App, error) {
	return uc.repo.List(ctx, pagination)
}

func (uc *AppUsecase) GetSimpleList(ctx context.Context, namespace string) ([]*NamespaceApp, error) {
	return uc.repo.SelectByNamespace(ctx, namespace)
}

func (uc *AppUsecase) CreateApp(ctx context.Context, app *App) error {
	return uc.repo.Create(ctx, app)
}

func (uc *AppUsecase) UpdateApp(ctx context.Context, app *App) error {
	return uc.repo.Update(ctx, app)
}
