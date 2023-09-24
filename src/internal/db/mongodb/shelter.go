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

type ShelterRepo struct {
	collection *mongo.Collection
}

func NewShelterRepo(db *mongo.Database) *ShelterRepo {
	return &ShelterRepo{
		collection: db.Collection("shelters"),
	}
}

func (r *ShelterRepo) GetById(ctx context.Context, id int) (*entities.Shelter, error) {
	filter := bson.M{"_id": id}
	result := r.collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	var sh entities.Shelter
	if err = result.Decode(&sh); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
	}
	return &sh, nil
}

func (r *ShelterRepo) GetAll(ctx context.Context) (entities.Shelters, error) {
	shelters := make(entities.Shelters, 0)

	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return shelters, nil
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	if err = cur.All(ctx, &shelters); err == nil {
		return shelters, nil
	}
	return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
}

func (r *ShelterRepo) Create(ctx context.Context, shelter *entities.Shelter) (int, error) {
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(1)

	var an entities.Shelter
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

		shelter.Id = nID
		_, err = r.collection.InsertOne(ctx, shelter)
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

func (r *ShelterRepo) Update(ctx context.Context, shelter *entities.Shelter) error {
	filter := bson.M{"_id": shelter.Id}

	shByte, err := bson.Marshal(shelter)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, "failed to marshal document")
	}

	var updateObj bson.M
	err = bson.Unmarshal(shByte, &updateObj)
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

func (r *ShelterRepo) Delete(ctx context.Context, id int) error {
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
