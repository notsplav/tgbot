package internal

import (
	"context"
	"errors"
)

type Video struct {
	Path     string // путь к файлу на диске
	Size     int64  // размер в байтах
	Duration int    // длительность в секундах
	Width    int
	Height   int
}

var (
	ErrUnsupported = errors.New("ссылка не поддерживается")
	ErrNotFound    = errors.New("видео удалено или приватное")
	ErrTooLarge    = errors.New("видео больше 50 МБ")
)

type Downloader interface {
	Download(ctx context.Context, url string) (*Video, error)
}
