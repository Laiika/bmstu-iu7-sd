package purchase

import "context"

type PurchaseService struct {
	repo IPurchaseRepo
}

func NewPurchaseService(repo IPurchaseRepo) *PurchaseService {
	return &PurchaseService{
		repo: repo,
	}
}

func (r *PurchaseService) GetById(ctx context.Context, id int) (*Purchase, error) {
	return r.repo.GetById(ctx, id)
}
func (r *PurchaseService) GetAll(ctx context.Context) (Purchases, error) {
	return r.repo.GetAll(ctx)
}
func (r *PurchaseService) Create(ctx context.Context, dto *CreatePurchase) (*Purchase, error) {
	return r.repo.Create(ctx, dto)
}
func (r *PurchaseService) Update(ctx context.Context, id int, dto *UpdatePurchase) (*Purchase, error) {
	return r.repo.Update(ctx, id, dto)
}
func (r *PurchaseService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
