package hub_repository

import (
	"awesomeProject/Project/WMS/internal/domain"
	error3 "awesomeProject/Project/WMS/pkg/error"
	"context"
	"errors"
	"fmt"
	"github.com/omniful/go_commons/db/sql/postgres"
	error2 "github.com/omniful/go_commons/error"
	"sync"
)

type Repository struct {
	db *postgres.DbCluster
}

type HubRepository interface {
	CreateHub(ctx context.Context, hub *domain.Hub) error2.CustomError
	GetHubByID(c context.Context, id int) (domain.Hub, error2.CustomError)
	DeleteHub(ctx context.Context, id int) error
}

var repo *Repository
var repoOnce sync.Once

func NewRepository(db *postgres.DbCluster) *Repository {
	repoOnce.Do(func() {
		repo = &Repository{
			db: db,
		}
	})
	return repo
}

func (r *Repository) CreateHub(c context.Context, hub *domain.Hub) (domain.Hub, error2.CustomError) {
	result := r.db.GetMasterDB(c).Create(&hub)
	if resultErr := result.Error; resultErr != nil {
		err := error2.NewCustomError(error3.SqlCreateError, fmt.Sprintf("Could not create hub for condition  : %+v, err: %v", hub, resultErr))
		return domain.Hub{}, err
	}
	return result, error2.CustomError{}
}

func (r *Repository) GetHubByID(ctx context.Context, id int) (domain.Hub, error) {
	var hub domain.Hub
	if id <= 0 {
		return hub, errors.New("invalid ID")
	}

	// Fetching hub with its database ID
	result := r.db.GetMasterDB(ctx).First(&hub, id)
	if result.Error != nil {
		return domain.Hub{}, result.Error
	}
	return hub, nil
}

func (r *Repository) DeleteHub(ctx context.Context, id int) error {
	result := r.db.GetMasterDB(ctx).Delete(&domain.Hub{}, id)
	return result.Error
}
