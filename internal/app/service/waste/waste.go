package waste

import (
	"github.com/jmoiron/sqlx"

	"github.com/Msik/h-microservice/internal/app/repository"
)

type WasteService struct {
	wasteRepository *repository.WasteRepository
}

func New(db *sqlx.DB) *WasteService {
	return &WasteService{wasteRepository: repository.NewWasteRepository(db)}
}
