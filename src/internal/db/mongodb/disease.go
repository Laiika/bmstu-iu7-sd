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

type DiseaseRepo struct {
	collection *mongo.Collection
}

func NewDiseaseRepo(db *mongo.Database) *DiseaseRepo {
	return &DiseaseRepo{
		collection: db.Collection("diseases"),
	}
}

func (r *DiseaseRepo) GetById(ctx context.Context, id int) (*entities.Disease, error) {
	filter := bson.M{"_id": id}
	result := r.collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	var d entities.Disease
	if err = result.Decode(&d); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
	}
	return &d, nil
}

func (r *DiseaseRepo) GetAll(ctx context.Context) (entities.Diseases, error) {
	diseases := make(entities.Diseases, 0)

	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return diseases, nil
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	if err = cur.All(ctx, &diseases); err == nil {
		return diseases, nil
	}
	return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
}

func (r *DiseaseRepo) GetAnimalAll(ctx context.Context, anId int) (entities.Diseases, error) {
	diseases := make(entities.Diseases, 0)

	filter := bson.M{"animal_id": bson.M{"$eq": anId}}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return diseases, nil
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	if err = cur.All(ctx, &diseases); err == nil {
		return diseases, nil
	}
	return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
}

func (r *DiseaseRepo) Create(ctx context.Context, disease *entities.Disease) (int, error) {
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(1)

	var d entities.Disease
	cursor, err := r.collection.Find(ctx, bson.M{}, &findOptions)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to execute query")
	}

	nID := 1
	tryCount := 3
	for tryCount >= 0 {
		if cursor.Next(ctx) {
			err = cursor.Decode(&d)
			if err != nil {
				return 0, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
			}
			nID = d.Id + 1
		} else if tryCount < 3 {
			return 0, pkgErrors.WithMessage(apperrors.ErrInternal, "duplicate key error")
		}

		tryCount--

		disease.Id = nID
		_, err = r.collection.InsertOne(ctx, disease)
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

func (r *DiseaseRepo) Update(ctx context.Context, disease *entities.Disease) error {
	filter := bson.M{"_id": disease.Id}

	dByte, err := bson.Marshal(disease)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, "failed to marshal document")
	}

	var updateObj bson.M
	err = bson.Unmarshal(dByte, &updateObj)
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

func (r *DiseaseRepo) Delete(ctx context.Context, id int) error {
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
