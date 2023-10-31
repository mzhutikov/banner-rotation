package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mzhutikov/banner-rotation/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestCreateGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &SocDemGroupRepo{
		db: db,
	}

	mock.ExpectQuery("INSERT INTO user_groups").WithArgs("TestUserGroup").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := r.CreateGroup(context.Background(), models.SocDemGroup{Description: "TestUserGroup"})
	require.NoError(t, err)
	require.Equal(t, int64(1), id)
}

func TestGetGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &SocDemGroupRepo{
		db: db,
	}

	expectedGroup := models.SocDemGroup{ID: 1, Description: "TestUserGroup"}
	rows := sqlmock.NewRows([]string{"id", "description"}).
		AddRow(expectedGroup.ID, expectedGroup.Description)

	mock.ExpectQuery("SELECT id, description FROM user_groups WHERE id = \\$1").WithArgs(1).WillReturnRows(rows)

	group, err := r.GetGroup(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, expectedGroup, group)
}

func TestUpdateGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &SocDemGroupRepo{
		db: db,
	}

	groupToUpdate := models.SocDemGroup{ID: 1, Description: "UpdatedGroupDescription"}

	mock.ExpectExec("UPDATE user_groups SET description = \\$1 WHERE id = \\$2").
		WithArgs(groupToUpdate.Description, groupToUpdate.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = r.UpdateGroup(context.Background(), groupToUpdate)
	require.NoError(t, err)
}

func TestDeleteGroup(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &SocDemGroupRepo{
		db: db,
	}

	groupIDToDelete := int64(1)

	mock.ExpectExec("DELETE FROM user_groups WHERE id = \\$1").
		WithArgs(groupIDToDelete).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = r.DeleteGroup(context.Background(), groupIDToDelete)
	require.NoError(t, err)
}
