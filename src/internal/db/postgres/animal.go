package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	pkgErrors "github.com/pkg/errors"
	"sd/internal/apperrors"
	"sd/internal/domain/entities"
	"sd/pkg/client/postgresql"
)

type AnimalRepo struct {
	client postgresql.Client
}

func NewAnimalRepo(client postgresql.Client) *AnimalRepo {
	return &AnimalRepo{
		client: client,
	}
}

func (r *AnimalRepo) GetById(ctx context.Context, id int) (*entities.Animal, error) {
	q := `
		SELECT id, name, age, height, weight, shelter_id, type, gender
		FROM animals
		WHERE id = $1
	`
	var an entities.Animal
	err := r.client.QueryRow(ctx, q, id).Scan(&an.Id, &an.Name, &an.Age, &an.Height, &an.Weight, &an.ShelterId, &an.Type, &an.Gender)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return &an, nil
}

func (r *AnimalRepo) GetAll(ctx context.Context) (entities.Animals, error) {
	q := `
		SELECT id, name, age, height, weight, shelter_id, type, gender
		FROM animals
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	animals := make(entities.Animals, 0)
	for rows.Next() {
		var an entities.Animal

		err = rows.Scan(&an.Id, &an.Name, &an.Age, &an.Height, &an.Weight, &an.ShelterId, &an.Type, &an.Gender)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
		}

		animals = append(animals, &an)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return animals, nil
}

func (r *AnimalRepo) GetCrtrAll(ctx context.Context, crtrId int) (entities.Animals, error) {
	q := `
		SELECT a.id, a.name, a.age, a.height, a.weight, a.shelter_id, a.type, a.gender
		FROM animals a
		JOIN (SELECT animal_id
			  FROM curators_animals
			  WHERE curator_id = $1) AS ca
		ON ca.animal_id = a.id
	`
	rows, err := r.client.Query(ctx, q, crtrId)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	animals := make(entities.Animals, 0)
	for rows.Next() {
		var an entities.Animal

		err = rows.Scan(&an.Id, &an.Name, &an.Age, &an.Height, &an.Weight, &an.ShelterId, &an.Type, &an.Gender)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
		}

		animals = append(animals, &an)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return animals, nil
}

func (r *AnimalRepo) Create(ctx context.Context, animal *entities.Animal) (int, error) {
	q := `
		INSERT INTO animals
		    (name, age, height, weight, shelter_id, type, gender) 
		VALUES 
		    ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id
	`
	var id int
	err := r.client.QueryRow(ctx, q, animal.Name, animal.Age, animal.Height, animal.Weight, animal.ShelterId, animal.Type, animal.Gender).Scan(&id)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return id, nil
}

func (r *AnimalRepo) Update(ctx context.Context, animal *entities.Animal) error {
	q := `
		UPDATE animals
		SET 
		    name = $1,
		    age = $2,
		    height = $3, 
		    weight = $4,
		    shelter_id = $5,
		    type = $6,
		    gender = $7
		WHERE id = $8
	`
	commandTag, err := r.client.Exec(ctx, q, animal.Name, animal.Age, animal.Height, animal.Weight, animal.ShelterId, animal.Type, animal.Gender, animal.Id)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
	}

	return nil
}

func (r *AnimalRepo) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM animals
		WHERE id = $1
	`
	commandTag, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
	}

	return nil
}
