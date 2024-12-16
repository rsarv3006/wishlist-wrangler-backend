package coresharedrepository

import (
	"fmt"
	"gorm.io/gorm"
)

type FindOptions struct {
	Limit  int
	Offset int
	Order  string
}

type CrudRepository[T any, ID comparable] struct {
	db *gorm.DB
}

func NewCrudRepository[T any, ID comparable](db *gorm.DB) *CrudRepository[T, ID] {
	return &CrudRepository[T, ID]{db: db}
}

func (r *CrudRepository[T, ID]) Save(entity *T) (*T, error) {
	if err := r.db.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *CrudRepository[T, ID]) SaveAll(entities []*T) ([]*T, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	for _, entity := range entities {
		if err := tx.Save(entity).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	return entities, tx.Commit().Error
}

func (r *CrudRepository[T, ID]) FindById(id ID) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *CrudRepository[T, ID]) ExistsById(id ID) (bool, error) {
	var count int64
	if err := r.db.Model(new(T)).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *CrudRepository[T, ID]) FindAll(opts *FindOptions) ([]*T, error) {
	query := r.db
	if opts != nil {
		if opts.Limit > 0 {
			query = query.Limit(opts.Limit)
		}
		if opts.Offset > 0 {
			query = query.Offset(opts.Offset)
		}
		if opts.Order != "" {
			query = query.Order(opts.Order)
		}
	}

	var entities []*T
	if err := query.Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to find entities: %w", err)
	}
	return entities, nil
}

func (r *CrudRepository[T, ID]) FindAllById(ids []ID) ([]*T, error) {
	var entities []*T
	if err := r.db.Find(&entities, ids).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *CrudRepository[T, ID]) Count() (int64, error) {
	var count int64
	if err := r.db.Model(new(T)).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CrudRepository[T, ID]) DeleteById(id ID) error {
	if err := r.db.Delete(new(T), id).Error; err != nil {
		return err
	}
	return nil
}

func (r *CrudRepository[T, ID]) Delete(entity *T) error {
	if err := r.db.Delete(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r *CrudRepository[T, ID]) DeleteAllById(ids []ID) error {
	if err := r.db.Delete(new(T), ids).Error; err != nil {
		return err
	}
	return nil
}

func (r *CrudRepository[T, ID]) DeleteAll(entities []*T) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, entity := range entities {
		if err := tx.Delete(entity).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *CrudRepository[T, ID]) DeleteAllEntities() error {
	if err := r.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(new(T)).Error; err != nil {
		return err
	}
	return nil
}
