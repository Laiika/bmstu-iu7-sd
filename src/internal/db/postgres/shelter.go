package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	pkgErrors "github.com/pkg/errors"
	"sd/internal/apperrors"
	"sd/internal/domain/entities"
	"sd/pkg/client/postgresql"
)

type ShelterRepo struct {
	client postgresql.Client
}

func NewShelterRepo(client postgresql.Client) *ShelterRepo {
	return &ShelterRepo{
		client: client,
	}
}

func (r *ShelterRepo) GetById(ctx context.Context, id int) (*entities.Shelter, error) {
	q := `
		SELECT id, street, house
		FROM shelters
		WHERE id = $1
	`
	var sh entities.Shelter
	err := r.client.QueryRow(ctx, q, id).Scan(&sh.Id, &sh.Street, &sh.House)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return &sh, nil
}

func (r *ShelterRepo) GetAll(ctx context.Context) (entities.Shelters, error) {
	q := `
		SELECT id, street, house
		FROM shelters
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	shelters := make(entities.Shelters, 0)
	for rows.Next() {
		var sh entities.Shelter

		err = rows.Scan(&sh.Id, &sh.Street, &sh.House)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
		}

		shelters = append(shelters, &sh)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return shelters, nil
}

func (r *ShelterRepo) Create(ctx context.Context, shelter *entities.Shelter) (int, error) {
	q := `
		INSERT INTO shelters
		    (street, house) 
		VALUES 
		    ($1, $2) 
		RETURNING id
	`
	var id int
	err := r.client.QueryRow(ctx, q, shelter.Street, shelter.House).Scan(&id)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return id, nil
}

func (r *ShelterRepo) Update(ctx context.Context, shelter *entities.Shelter) error {
	q := `
		UPDATE shelters
		SET 
			street = $1,
		    house = $2
		WHERE id = $3
	`
	commandTag, err := r.client.Exec(ctx, q, shelter.Street, shelter.House, shelter.Id)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
	}

	return nil
}

func (r *ShelterRepo) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM shelters
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
