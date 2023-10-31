package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
	"github.com/mzhutikov/banner-rotation/pkg/models"
)

func (sr *BannerRepo) CreateBanner(ctx context.Context, s models.Banner) (int64, error) {
	query := "INSERT INTO banners(description) VALUES($1) RETURNING id"
	var id int64
	err := sr.db.QueryRow(query, s.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (sr *BannerRepo) GetBanner(ctx context.Context, id int64) (models.Banner, error) {
	var banner models.Banner
	query := "SELECT id, description FROM banners WHERE id = $1"
	err := sr.db.QueryRowContext(ctx, query, id).Scan(&banner.ID, &banner.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return banner, fmt.Errorf("banner with ID %d not found", id)
		}
		return banner, err
	}
	return banner, nil
}

func (sr *BannerRepo) UpdateBanner(ctx context.Context, banner models.Banner) error {
	query := "UPDATE banners SET description = $1 WHERE id = $2"
	res, err := sr.db.ExecContext(ctx, query, banner.Description, banner.ID)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("banner with ID %d not found", banner.ID)
	}

	return nil
}

func (sr *BannerRepo) DeleteBanner(ctx context.Context, id int64) error {
	query := "DELETE FROM banners WHERE id = $1"
	res, err := sr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("banner with ID %d not found", id)
	}

	return nil
}
