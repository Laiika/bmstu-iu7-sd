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

type PurchaseRepo struct {
	collection *mongo.Collection
}

func NewPurchaseRepo(db *mongo.Database) *PurchaseRepo {
	return &PurchaseRepo{
		collection: db.Collection("purchases"),
	}
}

func (r *PurchaseRepo) GetById(ctx context.Context, id int) (*entities.Purchase, error) {
	filter := bson.M{"_id": id}
	result := r.collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	var p entities.Purchase
	if err = result.Decode(&p); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
	}
	return &p, nil
}

func (r *PurchaseRepo) GetAll(ctx context.Context) (entities.Purchases, error) {
	purchases := make(entities.Purchases, 0)

	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return purchases, nil
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	if err = cur.All(ctx, &purchases); err == nil {
		return purchases, nil
	}
	return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
}

func (r *PurchaseRepo) Create(ctx context.Context, purchase *entities.Purchase) (int, error) {
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(1)

	var p entities.Purchase
	cursor, err := r.collection.Find(ctx, bson.M{}, &findOptions)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to execute query")
	}

	nID := 1
	tryCount := 3
	for tryCount >= 0 {
		if cursor.Next(ctx) {
			err = cursor.Decode(&p)
			if err != nil {
				return 0, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
			}
			nID = p.Id + 1
		} else if tryCount < 3 {
			return 0, pkgErrors.WithMessage(apperrors.ErrInternal, "duplicate key error")
		}

		tryCount--

		purchase.Id = nID
		_, err = r.collection.InsertOne(ctx, purchase)
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

func (r *PurchaseRepo) Update(ctx context.Context, purchase *entities.Purchase) error {
	filter := bson.M{"_id": purchase.Id}

	pByte, err := bson.Marshal(purchase)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, "failed to marshal document")
	}

	var updateObj bson.M
	err = bson.Unmarshal(pByte, &updateObj)
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

func (r *PurchaseRepo) Delete(ctx context.Context, id int) error {
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
