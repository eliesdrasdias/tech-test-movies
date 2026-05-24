package ports

import (
	"context"
	"tech-test-movies/movie-service/internal/core/domain"
)

type MovieRepository interface {
	Get(ctx context.Context, id int64) (*domain.Movie, error)
	List(ctx context.Context) ([]*domain.Movie, error)
	Create(ctx context.Context, movie *domain.Movie) error
	Delete(ctx context.Context, id int64) error
}
