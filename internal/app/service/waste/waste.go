package waste

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Msik/h-microservice/internal/app/repository"
)

type WasteService struct {
	wasteRepository *repository.WasteRepository
}

func New(db *sqlx.DB) *WasteService {
	return &WasteService{wasteRepository: repository.NewWasteRepository(db)}
}

func (ws *WasteService) Add(ctx context.Context, amount float32, category_id uint64) (uint64, error) {
	id, err := ws.wasteRepository.Store(ctx, amount, category_id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ws *WasteService) Delete(ctx context.Context, waste_id uint64) error {
	if err := ws.wasteRepository.Delete(ctx, waste_id); err != nil {
		return err
	}

	return nil
}
