package integrational

import (
	"context"
	"sd/pkg/client/postgresql"
)

func TruncateTables(client postgresql.Client, ctx context.Context) error {
	q := `
		TRUNCATE animals, shelters, curators, purchases, curators_animals
	`
	_, err := client.Exec(ctx, q)
	if err != nil {
		return err
	}

	return nil
}
