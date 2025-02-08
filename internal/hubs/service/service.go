package hService

import (
	"awesomeProject/Project/WMS/internal/domain"
	"awesomeProject/Project/WMS/internal/hubs/repository"
	"awesomeProject/Project/WMS/internal/hubs/responses"

	//"awesomeProject/Project/WMS/internal/hubs/requests"
	//"awesomeProject/Project/WMS/internal/hubs/responses"
	"context"
	error2 "github.com/omniful/go_commons/error"
)

type Service struct {
	repo hub_repository.HubRepository
}

type HubService interface {
	CreateHub(ctx context.Context, hub domain.Hub) (*responses.CreateHubResponse, error2.CustomError)
	GetHubByID(ctx context.Context, id uint) (domain.Hub, error2.CustomError)
	DeleteHub(ctx context.Context, id int) error
}

func NewHubService(hubRepo hub_repository.HubRepository) *Service {
	return &Service{
		repo: hubRepo,
	}
}

func (s *Service) CreateHub(ctx context.Context, hub domain.Hub) (domain.Hub, error2.CustomError) {
	return s.repo.CreateHub(ctx, &hub), error2.CustomError{}
}

func (s *Service) GetHubByID(ctx context.Context, id int) (domain.Hub, error) {
	hub, err := s.repo.GetHubByID(ctx, id)
	if err.Exists() {
		return domain.Hub{}, err
	}
	return hub, nil
}

func (s *Service) DeleteHub(ctx context.Context, id int) error {
	return s.repo.DeleteHub(ctx, id)
}
