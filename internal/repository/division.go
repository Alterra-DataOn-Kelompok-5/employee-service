package repository

import (
	"context"
	"strings"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/model"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/dto"
	"gorm.io/gorm"
)

type Division interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Division, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.Division, error)
	Save(ctx context.Context, division *dto.CreateDivisionRequestBody) (model.Division, error)
	Edit(ctx context.Context, oldEmployee *model.Division, updateData *dto.UpdateDivisionRequestBody) (*model.Division, error)
	Destroy(ctx context.Context, division *model.Division) (*model.Division, error)
	ExistByName(ctx context.Context, name string) (bool, error)
}

type division struct {
	Db *gorm.DB
}

func NewDivisionRepository(db *gorm.DB) *division {
	return &division{
		db,
	}
}

func (r *division) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Division, *pkgdto.PaginationInfo, error) {
	var divisions []model.Division
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.Division{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&divisions).Error

	return divisions, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *division) FindByID(ctx context.Context, id uint) (model.Division, error) {
	var division model.Division
	if err := r.Db.WithContext(ctx).Model(&model.Division{}).Where("id = ?", id).First(&division).Error; err != nil {
		return division, err
	}
	return division, nil
}

func (r *division) Save(ctx context.Context, division *dto.CreateDivisionRequestBody) (model.Division, error) {
	newDivision := model.Division{
		Name: *division.Name,
	}
	if err := r.Db.WithContext(ctx).Save(&newDivision).Error; err != nil {
		return newDivision, err
	}
	return newDivision, nil
}

func (r *division) Edit(ctx context.Context, oldDivision *model.Division, updateData *dto.UpdateDivisionRequestBody) (*model.Division, error) {
	if updateData.Name != nil {
		oldDivision.Name = *updateData.Name
	}

	if err := r.Db.WithContext(ctx).Save(oldDivision).Find(oldDivision).Error; err != nil {
		return nil, err
	}

	return oldDivision, nil
}

func (r *division) Destroy(ctx context.Context, division *model.Division) (*model.Division, error) {
	if err := r.Db.WithContext(ctx).Delete(division).Error; err != nil {
		return nil, err
	}
	return division, nil
}

func (r *division) ExistByName(ctx context.Context, name string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Division{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}
