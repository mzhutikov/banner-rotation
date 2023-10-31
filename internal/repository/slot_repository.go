package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
	"github.com/mzhutikov/banner-rotation/pkg/models"
)

func (sr *SlotRepo) CreateSlot(ctx context.Context, s models.Slot) (int64, error) {
	query := "INSERT INTO slots(description) VALUES($1) RETURNING id"
	var id int64
	err := sr.db.QueryRow(query, s.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (sr *SlotRepo) GetSlot(ctx context.Context, id int64) (models.Slot, error) {
	var slot models.Slot
	query := "SELECT id, description FROM slots WHERE id = $1"
	err := sr.db.QueryRowContext(ctx, query, id).Scan(&slot.ID, &slot.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return slot, fmt.Errorf("slot with ID %d not found", id)
		}
		return slot, err
	}
	return slot, nil
}

func (sr *SlotRepo) UpdateSlot(ctx context.Context, slot models.Slot) error {
	query := "UPDATE slots SET description = $1 WHERE id = $2"
	res, err := sr.db.ExecContext(ctx, query, slot.Description, slot.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("slot with ID %d not found", slot.ID)
	}

	return nil
}

func (sr *SlotRepo) DeleteSlot(ctx context.Context, id int64) error {
	query := "DELETE FROM slots WHERE id = $1"
	res, err := sr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("slot with ID %d not found", id)
	}

	return nil
}