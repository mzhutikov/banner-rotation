package repository

import (
	"database/sql"
	"fmt"

	"context"

	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
	"github.com/mzhutikov/banner-rotation/pkg/models"
)

// Config определяет параметры подключения к базе данных.
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// Database представляет структуру для взаимодействия с базой данных.
type Database struct {
	conn *sql.DB
}

// SlotRepository определяет методы для работы со слотами.
type SlotRepository interface {
	CreateSlot(ctx context.Context, slot models.Slot) (int64, error)
	GetSlot(ctx context.Context, id int64) (models.Slot, error)
	UpdateSlot(ctx context.Context, slot models.Slot) error
	DeleteSlot(ctx context.Context, id int64) error
}

// BannerRepository определяет методы для работы с баннерами.
type BannerRepository interface {
	CreateBanner(ctx context.Context, banner models.Banner) (int64, error)
	GetBanner(ctx context.Context, id int64) (models.Banner, error)
	UpdateBanner(ctx context.Context, banner models.Banner) error
	DeleteBanner(ctx context.Context, id int64) error
}

// SocDemGroupRepository определяет методы для работы с социально-демографическими группами.
type SocDemGroupRepository interface {
	CreateGroup(ctx context.Context, group models.SocDemGroup) (int64, error)
	GetGroup(ctx context.Context, id int64) (models.SocDemGroup, error)
	UpdateGroup(ctx context.Context, group models.SocDemGroup) error
	DeleteGroup(ctx context.Context, id int64) error
}

type SlotRepo struct {
	db *sql.DB
}

type BannerRepo struct {
	db *sql.DB
}

type SocDemGroupRepo struct {
	db *sql.DB
}

// Repo объединяет все репозитории.
type Repo struct {
	SlotRepository
	BannerRepository
	SocDemGroupRepository
}

// NewDatabase создает и возвращает новый экземпляр Database.
func NewDatabase(cfg Config) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{conn: db}, nil
}

// Close закрывает соединение с базой данных.
func (d *Database) Close() {
	d.conn.Close()
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		SlotRepository:      &SlotRepo{db: db},
		BannerRepository:    &BannerRepo{db: db},
		SocDemGroupRepository: &SocDemGroupRepo{db: db},
	}
}

func (s *SlotRepo) CreateSlot(slot *models.Slot) error {
	query := "INSERT INTO slots(id, description) VALUES($1, $2)"
	_, err := s.db.Exec(query, slot.ID, slot.Description)
	return err
}

func (s *BannerRepo) CreateBanner(banner *models.Banner) error {
	query := "INSERT INTO banners(id, description) VALUES($1, $2)"
	_, err := s.db.Exec(query, banner.ID, banner.Description)
	return err
}

func (s *SocDemGroupRepo) CreateGroup(group *models.Group) error {
	query := "INSERT INTO user_groups(id, description) VALUES($1, $2)"
	_, err := s.db.Exec(query, group.ID, group.Description)
	return err
}
