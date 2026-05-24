package services

import (
	"context"
	"testing"

	"movie-service/internal/core/domain"
)

type mockMovieRepository struct{}

func (m *mockMovieRepository) Get(ctx context.Context, id int64) (*domain.Movie, error) {
	return &domain.Movie{
		Title: "Mocked Movie",
		Year:  "2026",
	}, nil
}

func (m *mockMovieRepository) List(ctx context.Context) ([]*domain.Movie, error)     { return nil, nil }
func (m *mockMovieRepository) Create(ctx context.Context, movie *domain.Movie) error { return nil }
func (m *mockMovieRepository) Delete(ctx context.Context, id int64) error            { return nil }

func TestMovieService_Get(t *testing.T) {
	mockRepo := &mockMovieRepository{}

	service := NewMovieService(mockRepo)

	ctx := context.Background()
	movie, err := service.GetMovie(ctx, 1)

	if err != nil {
		t.Fatalf("Não era esperado erro, recebido: %v", err)
	}

	if movie.Title != "Mocked Movie" {
		t.Errorf("Esperado 'Mocked Movie', recebido '%s'", movie.Title)
	}
}
