package repository

import (
	"context"

	"github.com/snsilvam/partydise.git/models"
)

type OurServicesRepository interface {
	InsertOurServices(ctx context.Context, service *models.OurServices) error
	GetOurServicesById(ctx context.Context, id int64) (*models.OurServices, error)
}

var implementation OurServicesRepository

func SetRepository(repository OurServicesRepository) {
	implementation = repository
}

func InsertOurServices(ctx context.Context, service *models.OurServices) error {
	return implementation.InsertOurServices(ctx, service)
}

func GetOurServicesById(ctx context.Context, id int64) (*models.OurServices, error) {
	return implementation.GetOurServicesById(ctx, id)
}
