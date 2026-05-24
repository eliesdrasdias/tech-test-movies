package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"movie-service/internal/core/domain"
	"movie-service/internal/core/ports"
)

// mongoRepository implementa a interface MovieRepository
type mongoRepository struct {
	collection *mongo.Collection
}

// NewMongoRepository cria um novo MovieRepository
func NewMongoRepository(collection *mongo.Collection) ports.MovieRepository {
	return &mongoRepository{
		collection: collection,
	}
}

func (r *mongoRepository) Get(ctx context.Context, id int64) (*domain.Movie, error) {
	var movie domain.Movie
	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}

	return &movie, nil
}

func (r *mongoRepository) List(ctx context.Context) ([]*domain.Movie, error) {
	findOptions := options.Find().SetLimit(100) // Limita a 100 resultados para evitar sobrecarga
	cursor, err := r.collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var movies []*domain.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *mongoRepository) Create(ctx context.Context, movie *domain.Movie) error {
	_, err := r.collection.InsertOne(ctx, movie)
	return err
}

func (r *mongoRepository) Delete(ctx context.Context, id int64) error {
	filter := bson.M{"_id": id}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("movie not found to delete")
	}
	return nil
}
