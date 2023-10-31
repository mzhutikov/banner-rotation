package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
	"github.com/mzhutikov/banner-rotation/pkg/models"
)

func (sr *SocDemGroupRepo) CreateGroup(ctx context.Context, s models.SocDemGroup) (int64, error) {
	query := "INSERT INTO user_groups(description) VALUES($1) RETURNING id"
	var id int64
	err := sr.db.QueryRow(query, s.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (sr *SocDemGroupRepo) GetGroup(ctx context.Context, id int64) (models.SocDemGroup, error) {
	var group models.SocDemGroup
	query := "SELECT id, description FROM user_groups WHERE id = $1"
	err := sr.db.QueryRowContext(ctx, query, id).Scan(&group.ID, &group.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return group, fmt.Errorf("group with ID %d not found", id)
		}
		return group, err
	}
	return group, nil
}

func (sr *SocDemGroupRepo) UpdateGroup(ctx context.Context, group models.SocDemGroup) error {
	query := "UPDATE user_groups SET description = $1 WHERE id = $2"
	res, err := sr.db.ExecContext(ctx, query, group.Description, group.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("group with ID %d not found", group.ID)
	}

	return nil
}

func (sr *SocDemGroupRepo) DeleteGroup(ctx context.Context, id int64) error {
	query := "DELETE FROM user_groups WHERE id = $1"
	res, err := sr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("group with ID %d not found", id)
	}

	return nil
}
