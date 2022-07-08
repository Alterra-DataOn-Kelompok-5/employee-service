package division

import (
	"context"
	"errors"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/constant"
	pkgdto "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/dto"
	res "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util/response"
)

type service struct {
	DivisionRepository repository.Division
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.DivisionResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.DivisionResponse, error)
	Store(ctx context.Context, payload *dto.CreateDivisionRequestBody) (*dto.DivisionResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateDivisionRequestBody) (*dto.DivisionResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.DivisionWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		DivisionRepository: f.DivisionRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.DivisionResponse], error) {
	divisions, info, err := s.DivisionRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.DivisionResponse

	for _, division := range divisions {
		data = append(data, dto.DivisionResponse{
			ID:   division.ID,
			Name: division.Name,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.DivisionResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}
func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.DivisionResponse, error) {
	var result dto.DivisionResponse
	data, err := s.DivisionRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.DivisionResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.DivisionResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Name = data.Name

	return &result, nil
}

func (s *service) Store(ctx context.Context, payload *dto.CreateDivisionRequestBody) (*dto.DivisionResponse, error) {
	var result dto.DivisionResponse
	isExist, err := s.DivisionRepository.ExistByName(ctx, *payload.Name)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return &result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("division already exists"))
	}

	data, err := s.DivisionRepository.Save(ctx, payload)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Name = data.Name

	return &result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateDivisionRequestBody) (*dto.DivisionResponse, error) {
	division, err := s.DivisionRepository.FindByID(ctx, *payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.DivisionResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.DivisionResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.DivisionRepository.Edit(ctx, &division, payload)
	if err != nil {
		return &dto.DivisionResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	var result dto.DivisionResponse
	result.ID = division.ID
	result.Name = division.Name

	return &result, nil
}
func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.DivisionWithCUDResponse, error) {
	division, err := s.DivisionRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.DivisionWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.DivisionWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.DivisionRepository.Destroy(ctx, &division)
	if err != nil {
		return &dto.DivisionWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.DivisionWithCUDResponse{
		DivisionResponse: dto.DivisionResponse{
			ID:   division.ID,
			Name: division.Name,
		},
		CreatedAt: division.CreatedAt,
		UpdatedAt: division.UpdatedAt,
		DeletedAt: division.DeletedAt,
	}

	return result, nil
}
