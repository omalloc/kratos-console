package data

import (
	"context"
	"fmt"
	"github.com/omalloc/contrib/protobuf"
	"gorm.io/gorm"

	"github.com/omalloc/kratos-console/internal/biz"
)

type appRepo struct {
	data *Data
}

// Create implements biz.AppRepo.
func (r *appRepo) Create(ctx context.Context, app *biz.App) error {
	return r.data.db.WithContext(ctx).Model(&biz.App{}).
		Create(app).Error
}

// Update implements biz.AppRepo.
func (r *appRepo) Update(ctx context.Context, app *biz.App) error {
	if app.ID <= 0 {
		return biz.ErrInvalidAppId
	}

	return r.data.db.WithContext(ctx).Model(&biz.App{}).
		Omit("id").
		Updates(&app).Error
}

func (r *appRepo) List(ctx context.Context, pagination *protobuf.Pagination, query *biz.AppQuery) ([]*biz.App, error) {
	var (
		apps []*biz.App
	)
	err := r.data.db.WithContext(ctx).Model(&biz.App{}).
		Scopes(ApplyAppQuery(query.Name, query.NamespaceID)).
		Count(pagination.Count()).
		Scopes(pagination.Paginate()).
		Find(&apps).Error
	return apps, err
}

// ApplyAppQuery applies the query conditions to the gorm.DB with Scopes.
func ApplyAppQuery(name string, nsID int64) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if name != "" {
			tx.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
		}
		if nsID > 0 {
			tx.Where("namespace_id = ?", nsID)
		}
		return tx
	}
}

func (r *appRepo) SelectByNamespace(ctx context.Context, namespace string) ([]*biz.NamespaceApp, error) {
	var (
		apps []*biz.NamespaceApp
		err  error
	)

	tx := r.data.db.WithContext(ctx).Model(&biz.App{}).
		Select("apps.*, t1.`name` AS namespace_name, t1.alias AS namespace_alias").
		Joins("LEFT JOIN namespaces AS t1 ON t1.id = apps.namespace_id")
	if namespace != "" {
		tx.Where("t1.`name` LIKE ?", fmt.Sprintf("%%%s%%", namespace))
	}

	err = tx.Find(&apps).Error
	return apps, err
}

func NewAppRepo(data *Data) biz.AppRepo {
	return &appRepo{
		data: data,
	}
}
