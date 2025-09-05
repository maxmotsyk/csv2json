package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func NewLogger() error {
	f, err := os.OpenFile("slog.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)

	if err != nil {
		return fmt.Errorf("open input.csv: %w", err)
	}
	logger := slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// міняємо тільки атрибут "time"
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				// формат у стилі "2006-01-02 15:04:05"
				return slog.Attr{
					Key:   slog.TimeKey,
					Value: slog.StringValue(t.Format("2006-01-02 15:04:05")),
				}
			}
			return a
		},
	}))

	slog.SetDefault(logger)
	return nil
}
