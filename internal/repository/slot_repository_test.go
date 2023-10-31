package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mzhutikov/banner-rotation/pkg/models"
	"github.com/stretchr/testify/require"
)


func TestCreateSlot(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &SlotRepo{
		db: db,
	}

	mock.ExpectQuery("INSERT INTO slots").WithArgs("TestSlot").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := r.CreateSlot(context.Background(), models.Slot{Description: "TestSlot"})
	require.NoError(t, err)
	require.Equal(t, int64(1), id)
}


func TestGetSlot(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &SlotRepo{
		db: db,
	}

	// Определите, что будет возвращаться при вызове запроса.
	expectedSlot := models.Slot{ID: 1, Description: "TestSlot"}
	rows := sqlmock.NewRows([]string{"id", "description"}).
		AddRow(expectedSlot.ID, expectedSlot.Description)

	mock.ExpectQuery("SELECT id, description FROM slots WHERE id = \\$1").WithArgs(1).WillReturnRows(rows)

	// Вызываем метод GetSlot и проверяем возвращаемое значение.
	slot, err := r.GetSlot(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, expectedSlot, slot)
}

func TestUpdateSlot(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &SlotRepo{
		db: db,
	}

	// Определите слот, который вы хотите обновить.
	slotToUpdate := models.Slot{ID: 1, Description: "UpdatedSlotDescription"}

	// Настраиваем мок на ожидание определенного запроса с определенными аргументами.
	mock.ExpectExec("UPDATE slots SET description = \\$1 WHERE id = \\$2").
		WithArgs(slotToUpdate.Description, slotToUpdate.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))  // тут мы говорим, что 1 строка была обновлена.

	// Вызываем метод UpdateSlot.
	err = r.UpdateSlot(context.Background(), slotToUpdate)
	require.NoError(t, err)
}

func TestDeleteSlot(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	r := &SlotRepo{
		db: db,
	}

	// ID слота, который вы хотите удалить.
	slotIDToDelete := int64(1)

	// Настраиваем мок на ожидание определенного запроса с определенными аргументами.
	mock.ExpectExec("DELETE FROM slots WHERE id = \\$1").
		WithArgs(slotIDToDelete).
		WillReturnResult(sqlmock.NewResult(1, 1))  // тут мы говорим, что 1 строка была удалена.

	// Вызываем метод DeleteSlot.
	err = r.DeleteSlot(context.Background(), slotIDToDelete)
	require.NoError(t, err)
}



