package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	pkgErrors "github.com/pkg/errors"
	"sd/internal/apperrors"
	"sd/internal/domain/entities"
	"sd/pkg/client/postgresql"
)

type CuratorRepo struct {
	client postgresql.Client
}

func NewCuratorRepo(client postgresql.Client) *CuratorRepo {
	return &CuratorRepo{
		client: client,
	}
}

func (r *CuratorRepo) GetById(ctx context.Context, id int) (*entities.Curator, error) {
	q := `
		SELECT id, chat_id, name, surname, phone_number
		FROM curators
		WHERE id = $1
	`
	var cur entities.Curator
	err := r.client.QueryRow(ctx, q, id).Scan(&cur.Id, &cur.ChatId, &cur.Name, &cur.Surname, &cur.PhoneNumber)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return &cur, nil
}

func (r *CuratorRepo) GetByChatId(ctx context.Context, chatId string) (*entities.Curator, error) {
	q := `
		SELECT id, chat_id, name, surname, phone_number
		FROM curators
		WHERE chat_id = $1
	`
	var cur entities.Curator
	err := r.client.QueryRow(ctx, q, chatId).Scan(&cur.Id, &cur.ChatId, &cur.Name, &cur.Surname, &cur.PhoneNumber)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
		}
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return &cur, nil
}

func (r *CuratorRepo) GetAll(ctx context.Context) (entities.Curators, error) {
	q := `
		SELECT id, chat_id, name, surname, phone_number
		FROM curators
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	curators := make(entities.Curators, 0)
	for rows.Next() {
		var cur entities.Curator

		err = rows.Scan(&cur.Id, &cur.ChatId, &cur.Name, &cur.Surname, &cur.PhoneNumber)
		if err != nil {
			return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
		}

		curators = append(curators, &cur)
	}

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return curators, nil
}

func (r *CuratorRepo) Create(ctx context.Context, curator *entities.Curator) (int, error) {
	_, err := r.GetByChatId(ctx, curator.ChatId)
	if err == nil {
		return 0, apperrors.ErrEntityExists
	}

	q := `
		INSERT INTO curators
		    (chat_id, name, surname, phone_number) 
		VALUES 
		    ($1, $2, $3, $4) 
		RETURNING id
	`
	var id int
	err = r.client.QueryRow(ctx, q, curator.ChatId, curator.Name, curator.Surname, curator.PhoneNumber).Scan(&id)
	if err != nil {
		return 0, pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}

	return id, nil
}

func (r *CuratorRepo) Update(ctx context.Context, curator *entities.Curator) error {
	q := `
		UPDATE curators
		SET 
			chat_id = $1,
		    name = $2,
		    surname = $3,
		    phone_number = $4
		WHERE id = $5
	`
	commandTag, err := r.client.Exec(ctx, q, curator.ChatId, curator.Name, curator.Surname, curator.PhoneNumber, curator.Id)
	if err != nil {
		return pkgErrors.WithMessage(apperrors.ErrInternal, err.Error())
	}
	if commandTag.RowsAffected() != 1 {
		return pkgErrors.WithMessage(apperrors.ErrEntityNotFound, err.Error())
	}

	return nil
}

func (r *CuratorRepo) Delete(ctx context.Context, id int) error {
	q := `
		DELETE FROM curators
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
