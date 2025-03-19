package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/omalloc/contrib/kratos/orm"
	"github.com/omalloc/contrib/kratos/orm/crud"
	"github.com/omalloc/contrib/protobuf"
)

type Role struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"column:name;type:varchar(64);comment:角色名"`
	Describe string `json:"describe" gorm:"column:describe;type:varchar(255);comment:描述"`
	Status   int64  `json:"status" gorm:"column:status;type:int;comment:状态"`

	orm.DBModel
}

func (Role) TableName() string {
	return "roles"
}

type RolePermission struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	RoleID     int64     `json:"role_id" gorm:"column:role_id;type:int;comment:角色ID"`
	PermID     int64     `json:"perm_id" gorm:"column:perm_id;type:int;comment:权限ID"`
	Actions    []*Action `json:"actions" gorm:"column:actions;type:json;serializer:json;comment:操作"`
	DataAccess []*Action `json:"data_access,omitempty" gorm:"column:data_access;type:json;serializer:json;comment:数据权限"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;type:datetime;comment:创建时间"`
}

func (RolePermission) TableName() string {
	return "roles_bind_permission"
}

type RoleRepo interface {
	crud.CRUD[Role]

	SelectFilterList(ctx context.Context, pagination *protobuf.Pagination) ([]*Role, error)
	BindPermission(ctx context.Context, roleID int64, permissionID int64, actions []*Action, dataAccess []*Action) error
	UnbindPermission(ctx context.Context, roleID int64, permissionID int64) error
}

type RoleUsecase struct {
	log      *log.Helper
	txm      orm.Transaction
	roleRepo RoleRepo
}

func NewRoleUsecase(repo RoleRepo, txm orm.Transaction, logger log.Logger) *RoleUsecase {
	return &RoleUsecase{
		log:      log.NewHelper(logger),
		txm:      txm,
		roleRepo: repo,
	}
}

func (uc *RoleUsecase) CreateRole(ctx context.Context, role *Role) error {
	return uc.txm.Transaction(ctx, func(ctx context.Context) error {
		return uc.roleRepo.Create(ctx, role)
	})
}

func (uc *RoleUsecase) UpdateRole(ctx context.Context, role *Role) error {
	return uc.txm.Transaction(ctx, func(ctx context.Context) error {
		return uc.roleRepo.Update(ctx, role.ID, role)
	})
}

func (uc *RoleUsecase) DeleteRole(ctx context.Context, id int64) error {
	return uc.txm.Transaction(ctx, func(ctx context.Context) error {
		return uc.roleRepo.Delete(ctx, id)
	})
}

func (uc *RoleUsecase) ListRole(ctx context.Context, pagination *protobuf.Pagination) ([]*Role, error) {
	return uc.roleRepo.SelectFilterList(ctx, pagination)
}

func (uc *RoleUsecase) BindPermission(ctx context.Context, roleID int64, permissionID int64, actions []*Action, dataAccess []*Action) error {
	return uc.txm.Transaction(ctx, func(ctx context.Context) error {
		return uc.roleRepo.BindPermission(ctx, roleID, permissionID, actions, dataAccess)
	})
}

func (uc *RoleUsecase) UnbindPermission(ctx context.Context, roleID int64, permissionID int64) error {
	return uc.txm.Transaction(ctx, func(ctx context.Context) error {
		return uc.roleRepo.UnbindPermission(ctx, roleID, permissionID)
	})
}
