package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	pkgErrors "github.com/pkg/errors"
	"sd/internal/apperrors"
	"sd/internal/domain/entities"
	"sd/pkg/client/postgresql"
)

type DiseaseRepo struct {
	client postgresql.Client
}

func NewDiseaseRepo(client postgresql.Client) *DiseaseRepo {
	return &DiseaseRepo{
		client: client,
	}
}

func (r *DiseaseRepo) GetById(ctx context.Context, id int) (*entities.Disease, error) {
	q := `
		SELECT id, diagnosis, symptoms, cause, is_chronic, animal_id 
		FROM diseases
		WHERE id = $1
	`
	var d entities.Disease
	err := r.client.QueryRow(ctx, q, id).Scan(&d.Id, &d.Diagnosis, &d.Symptoms, &d.Cause, &d.IsChronic, &d.AnimalId)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return &d, nil
}

func (r *DiseaseRepo) GetAll(ctx context.Context) (entities.Diseases, error) {
	q := `
		SELECT id, diagnosis, symptoms, cause, is_chronic, animal_id
		FROM diseases
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	diseases := make(entities.Diseases, 0)
	for rows.Next() {
		var d entities.Disease

		err = rows.Scan(&d.Id, &d.Diagnosis, &d.Symptoms, &d.Cause, &d.IsChronic, &d.AnimalId)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
		}

		diseases = append(diseases, &d)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return diseases, nil
}

func (r *DiseaseRepo) GetAnimalAll(ctx context.Context, anId int) (entities.Diseases, error) {
	q := `
		SELECT id, diagnosis, symptoms, cause, is_chronic, animal_id
		FROM diseases
		WHERE animal_id = $1
	`
	rows, err := r.client.Query(ctx, q, anId)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	diseases := make(entities.Diseases, 0)
	for rows.Next() {
		var d entities.Disease

		err = rows.Scan(&d.Id, &d.Diagnosis, &d.Symptoms, &d.Cause, &d.IsChronic, &d.AnimalId)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
		}

		diseases = append(diseases, &d)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return diseases, nil
}

func (r *DiseaseRepo) Create(ctx context.Context, disease *entities.Disease) (int, error) {
	q := `
		INSERT INTO diseases
		    (diagnosis, symptoms, cause, is_chronic, animal_id) 
		VALUES 
		    ($1, $2, $3, $4, $5) 
		RETURNING id
	`
	var id int
	err := r.client.QueryRow(ctx, q, disease.Diagnosis, disease.Symptoms, disease.Cause, disease.IsChronic, disease.AnimalId).Scan(&id)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return id, nil
}

func (r *DiseaseRepo) Update(ctx context.Context, disease *entities.Disease) error {
	q := `
		UPDATE diseases
		SET 
		    diagnosis = $1,
		    symptoms = $2,
		    cause = $3, 
		    is_chronic = $4,
		    animal_id = $5
		WHERE id = $6
	`
	commandTag, err := r.client.Exec(ctx, q, disease.Diagnosis, disease.Symptoms, disease.Cause, disease.IsChronic, disease.AnimalId, disease.Id)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
	}

	return nil
}

func (r *DiseaseRepo) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM diseases
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
