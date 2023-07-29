package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	pkgErrors "github.com/pkg/errors"
	"sd/internal/apperrors"
	"sd/internal/domain/entities"
	"sd/pkg/client/postgresql"
)

type PurchaseRepo struct {
	client postgresql.Client
}

func NewPurchaseRepo(client postgresql.Client) *PurchaseRepo {
	return &PurchaseRepo{
		client: client,
	}
}

func (r *PurchaseRepo) GetById(ctx context.Context, id int) (*entities.Purchase, error) {
	q := `
		SELECT id, name, frequency, cost, last_date, animal_id
		FROM purchases
		WHERE id = $1
	`
	var pur entities.Purchase
	err := r.client.QueryRow(ctx, q, id).Scan(&pur.Id, &pur.Name, &pur.Frequency, &pur.Cost, &pur.LastDate, &pur.AnimalId)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return &pur, nil
}

func (r *PurchaseRepo) GetAll(ctx context.Context) (entities.Purchases, error) {
	q := `
		SELECT id, name, frequency, cost, last_date, animal_id
		FROM purchases
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	purchases := make(entities.Purchases, 0)
	for rows.Next() {
		var pur entities.Purchase

		err = rows.Scan(&pur.Id, &pur.Name, &pur.Frequency, &pur.Cost, &pur.LastDate, &pur.AnimalId)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
		}

		purchases = append(purchases, &pur)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return purchases, nil
}

func (r *PurchaseRepo) Create(ctx context.Context, purchase *entities.Purchase) (int, error) {
	q := `
		INSERT INTO purchases
		    (name, frequency, cost, last_date, animal_id) 
		VALUES 
		    ($1, $2, $3, $4, $5) 
		RETURNING id
	`
	var id int
	err := r.client.QueryRow(ctx, q, purchase.Name, purchase.Frequency, purchase.Cost, purchase.LastDate, purchase.AnimalId).Scan(&id)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return id, nil
}

func (r *PurchaseRepo) Update(ctx context.Context, purchase *entities.Purchase) error {
	q := `
		UPDATE purchases
		SET 
			name = $1,
		    frequency = $2,
		    cost = $3,
			last_date = $4,
		    animal_id = $5
		WHERE id = $6
	`
	commandTag, err := r.client.Exec(ctx, q, purchase.Name, purchase.Frequency, purchase.Cost, purchase.LastDate, purchase.AnimalId, purchase.Id)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
	}

	return nil
}

func (r *PurchaseRepo) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM purchases
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
