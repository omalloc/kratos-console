package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/omalloc/contrib/kratos/orm/crud"
	"gorm.io/gorm"

	pb "github.com/omalloc/kratos-console/api/console/resource"
	"github.com/omalloc/kratos-console/internal/biz"
)

type appRuntimeRepo struct {
	crud.CRUD[biz.AppRuntime]

	data *Data
}

func (r *appRuntimeRepo) SelectListByNamespace(ctx context.Context, s string) ([]*biz.AppRuntime, error) {
	var (
		list []*biz.AppRuntime
		err  error
	)
	err = r.data.db.WithContext(ctx).
		Model(&biz.AppRuntime{}).
		Where("namespace = ?", s).
		Find(&list).Error

	return list, err
}

func (r *appRuntimeRepo) SelectListByNameAndNamespace(ctx context.Context, name, ns string) (*biz.AppPair, error) {
	var (
		mainList     []*biz.AppRuntime
		childrenList []*biz.AppRuntime
		err          error
	)

	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		t2 := tx.Model(&biz.AppRuntime{}).
			Select("max(id) AS id").
			Group("namespace, name").
			Scopes(r.filters(name, ns))

		if err := tx.Model(&biz.AppRuntime{}).
			Select([]string{
				"namespace",
				"name",
				"replace(`key`, concat('/', hostname), '') AS `key`",
				"version",
				"cluster",
			}).
			Joins("JOIN (?) AS t2 ON t2.id = apps_runtime.id", t2).
			Scopes(r.filters(name, ns)).
			Order("name DESC").
			Find(&mainList).Error; err != nil {
			return err
		}

		return tx.Model(&biz.AppRuntime{}).Select([]string{
			"id",
			"key",
			"hostname",
			"name",
			"namespace",
			"metadata",
			"endpoints",
			"version",
			"cluster",
			"TIMESTAMPDIFF(SECOND, updated_at, NOW()) AS `last_healthy_sec`",
		}).Scopes(r.filters(name, ns)).Find(&childrenList).Error
	})

	return &biz.AppPair{
		MainData: mainList,
		Children: childrenList,
	}, err
}

func (r *appRuntimeRepo) filters(name, ns string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if name != "" {
			tx.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
		}
		if ns != "" {
			tx.Where("namespace = ?", ns)
		}
		return tx
	}
}

func (r *appRuntimeRepo) BatchSave(ctx context.Context, data []*biz.AppRuntime) error {
	return r.data.db.WithContext(ctx).Model(&biz.AppRuntime{}).
		Save(&data).
		Error
}

func (r *appRuntimeRepo) UpdateHang(ctx context.Context, ID int64, hang bool) (*biz.AppRuntime, error) {
	var err error
	var app biz.AppRuntime

	err = r.data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&biz.AppRuntime{}).
			Where("id = ?", ID).
			Take(&app).Error; err != nil {
			return err
		}

		if app.Name == "kratos-agent" {
			return errors.BadRequest("system-component can't hang", "kratos-agent 是系统组件,不允许挂起")
		}

		if app.Metadata == nil {
			app.Metadata = make(map[string]string, 1)
		}
		// 需要有 Metadata 常量
		app.Metadata[pb.MetadataType_hang.String()] = fmt.Sprintf("%t", hang)

		return tx.Model(&biz.AppRuntime{}).Where("id = ?", ID).Save(&app).Error
	})

	return &app, err
}

func NewAppRuntimeRepo(data *Data) biz.AppRuntimeRepo {
	return &appRuntimeRepo{
		CRUD: crud.New[biz.AppRuntime](data.db),
		data: data,
	}
}
