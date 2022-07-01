package employee

import (
	"context"

	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/dto"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/factory"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/internal/repository"
	"github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/constant"
	res "github.com/Alterra-DataOn-Kelompok-5/employee-service/pkg/util/response"
)

type service struct {
	EmployeeRepository repository.Employee
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[dto.EmployeeResponse], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*dto.EmployeeResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		EmployeeRepository: f.EmployeeRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[dto.EmployeeResponse], error) {

	employees, info, err := s.EmployeeRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.EmployeeResponse

	for _, employee := range employees {
		data = append(data, dto.EmployeeResponse{
			ID:       employee.ID,
			Fullname: employee.Fullname,
			Email:    employee.Email,
		})

	}

	result := new(dto.SearchGetResponse[dto.EmployeeResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*dto.EmployeeResponse, error) {
	var result *dto.EmployeeResponse

	data, err := s.EmployeeRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.EmployeeResponse{
		ID:       data.ID,
		Fullname: data.Fullname,
		Email:    data.Email,
	}

	return result, nil
}
