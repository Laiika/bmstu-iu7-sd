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

type CuratorRepo struct {
	collection *mongo.Collection
}

func NewCuratorRepo(db *mongo.Database) *CuratorRepo {
	return &CuratorRepo{
		collection: db.Collection("curators"),
	}
}

func (r *CuratorRepo) GetById(ctx context.Context, id int) (*entities.Curator, error) {
	filter := bson.M{"_id": id}
	result := r.collection.FindOne(ctx, filter)
	err := result.Err()
	if err != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	var c entities.Curator
	if err = result.Decode(&c); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
	}
	return &c, nil
}

func (r *CuratorRepo) GetByChatId(ctx context.Context, chatId string) (*entities.Curator, error) {
	filter := bson.M{"chat_id": bson.M{"$eq": chatId}}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	var curator entities.Curator
	if cur.Next(ctx) {
		err = cur.Decode(&curator)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
		}
	} else {
		return nil, apperrors.ErrInternal
	}

	return &curator, nil
}

func (r *CuratorRepo) GetAll(ctx context.Context) (entities.Curators, error) {
	curators := make(entities.Curators, 0)

	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return curators, nil
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	if err = cur.All(ctx, &curators); err == nil {
		return curators, nil
	}
	return nil, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to decode document")
}

func (r *CuratorRepo) Create(ctx context.Context, curator *entities.Curator) (int, error) {
	findOptions := options.FindOptions{}
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(1)

	var c entities.Curator
	cursor, err := r.collection.Find(ctx, bson.M{}, &findOptions)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperrors.ErrInternal, "failed to execute query")
	}

	nID := 1
	tryCount := 3
	for tryCount >= 0 {
		if cursor.Next(ctx) {
			err = cursor.Decode(&c)
			if err != nil {
				return 0, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
			}
			nID = c.Id + 1
		} else if tryCount < 3 {
			return 0, pkgErrors.WithMessage(apperrors.ErrInternal, "duplicate key error")
		}

		tryCount--

		curator.Id = nID
		_, err = r.collection.InsertOne(ctx, curator)
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

func (r *CuratorRepo) Update(ctx context.Context, curator *entities.Curator) error {
	filter := bson.M{"_id": curator.Id}

	cByte, err := bson.Marshal(curator)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, "failed to marshal document")
	}

	var updateObj bson.M
	err = bson.Unmarshal(cByte, &updateObj)
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

func (r *CuratorRepo) Delete(ctx context.Context, id int) error {
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
