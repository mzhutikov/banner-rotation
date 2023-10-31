package models

// Slot представляет слот на сайте, где отображается баннер.
type Slot struct {
	ID        int64  // уникальный идентификатор слота
	Description string // описание слота
}

// Banner представляет баннер, который отображается в слоте.
type Banner struct {
	ID          int64  // уникальный идентификатор баннера
	Description string // описание баннера
}

// SocDemGroup представляет социально-демографическую группу пользователей.
type SocDemGroup struct {
	ID          int64  // уникальный идентификатор группы
	Description string // описание группы
}
