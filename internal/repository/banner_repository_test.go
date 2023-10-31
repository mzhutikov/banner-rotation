package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mzhutikov/banner-rotation/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestCreateBanner(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &BannerRepo{
		db: db,
	}

	mock.ExpectQuery("INSERT INTO banners").WithArgs("TestBanner").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := r.CreateBanner(context.Background(), models.Banner{Description: "TestBanner"})
	require.NoError(t, err)
	require.Equal(t, int64(1), id)
}

func TestGetBanner(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &BannerRepo{
		db: db,
	}

	expectedSlot := models.Banner{ID: 1, Description: "TestBanner"}
	rows := sqlmock.NewRows([]string{"id", "description"}).
		AddRow(expectedSlot.ID, expectedSlot.Description)

	mock.ExpectQuery("SELECT id, description FROM banners WHERE id = \\$1").WithArgs(1).WillReturnRows(rows)

	slot, err := r.GetBanner(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, expectedSlot, slot)
}

func TestUpdateBanner(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &BannerRepo{
		db: db,
	}

	bannerToUpdate := models.Banner{ID: 1, Description: "UpdatedBannerDescription"}

	mock.ExpectExec("UPDATE banners SET description = \\$1 WHERE id = \\$2").
		WithArgs(bannerToUpdate.Description, bannerToUpdate.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = r.UpdateBanner(context.Background(), bannerToUpdate)
	require.NoError(t, err)
}

func TestDeleteBanner(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &BannerRepo{
		db: db,
	}

	bannerIDToDelete := int64(1)

	mock.ExpectExec("DELETE FROM banners WHERE id = \\$1").
		WithArgs(bannerIDToDelete).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = r.DeleteBanner(context.Background(), bannerIDToDelete)
	require.NoError(t, err)
}
