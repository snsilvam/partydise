package repository

import (
	"context"

	"github.com/snsilvam/partydise.git/models"
)

type OurServices interface {
	InsertOurServices(ctx context.Context, service *models.OurServices) error
}
