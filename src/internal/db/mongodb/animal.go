package mongodb

import (
	"context"
	"errors"
	pkgErrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sd/internal/apperrors"
	"sd/internal/domain/entities"
)

type AnimalRepo struct {
	collection *mongo.Collection
}

func NewAnimalRepo(db *mongo.Database) *AnimalRepo {
	return &AnimalRepo{
		collection: db.Collection("animals"),
	}
}

func (r *AnimalRepo) GetById(ctx context.Context, id int) (*entities.Animal, error) {
	filter := bson.M{"_id": id}
	result := r.collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	var an entities.Animal
	if err = result.Decode(&an); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
	}
	return &an, nil
}

func (r *AnimalRepo) GetAll(ctx context.Context) (entities.Animals, error) {
	animals := make(entities.Animals, 0)

	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return animals, nil
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	if err = cur.All(ctx, &animals); err == nil {
		return animals, nil
	}
	return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
}

func (r *AnimalRepo) GetCrtrAll(ctx context.Context, crtrId int) (entities.Animals, error) {
	animals := make(entities.Animals, 0)

	filter := bson.M{"curators": bson.M{"$in": crtrId}}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return animals, nil
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	if err = cur.All(ctx, &animals); err == nil {
		return animals, nil
	}
	return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
}

func (r *AnimalRepo) Create(ctx context.Context, animal *entities.Animal) (int, error) {
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(1)

	var an entities.Animal
	cursor, err := r.collection.Find(ctx, bson.M{}, &findOptions)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to execute query")
	}

	nID := 1
	tryCount := 3
	for tryCount >= 0 {
		if cursor.Next(ctx) {
			err = cursor.Decode(&an)
			if err != nil {
				return 0, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
			}
			nID = an.Id + 1
		} else if tryCount < 3 {
			return 0, pkgErrors.WithMessage(apperrors.ErrInternal, "duplicate key error")
		}

		tryCount--

		animal.Id = nID
		_, err = r.collection.InsertOne(ctx, animal)
		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				continue
			} else {
				return 0, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to execute query")
			}
		}
		break
	}

	return nID, nil
}

func (r *AnimalRepo) Update(ctx context.Context, animal *entities.Animal) error {
	filter := bson.M{"_id": animal.Id}

	anByte, err := bson.Marshal(animal)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, "failed to marshal document")
	}

	var updateObj bson.M
	err = bson.Unmarshal(anByte, &updateObj)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, "failed to unmarshal document")
	}

	delete(updateObj, "_id")

	update := bson.M{
		"$set": updateObj,
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}
	if result.MatchedCount == 0 {
		return pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
	}

	return nil
}

func (r *AnimalRepo) Delete(ctx context.Context, id int) error {
	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}
	if result.DeletedCount == 0 {
		return pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
	}

	return nil
}
