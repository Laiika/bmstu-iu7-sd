package purchase

import "context"

type IPurchaseRepo interface {
	GetById(ctx context.Context, id int) (*Purchase, error)
	GetAll(ctx context.Context) (Purchases, error)
	Create(ctx context.Context, dto *CreatePurchase) (*Purchase, error)
	Update(ctx context.Context, id int, dto *UpdatePurchase) (*Purchase, error)
	Delete(ctx context.Context, id int) error
}
