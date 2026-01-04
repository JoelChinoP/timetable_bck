package scalars

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

// ParseTime intenta parsear formatos comunes
func ParseTime(s string) (Time, error) {
	s = strings.TrimSpace(s)

	// Permite HH:MM o HH:MM:SS
	layouts := []string{"15:04", "15:04:05"}
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return Time{Time: t}, nil
		}
	}
	return Time{}, fmt.Errorf("invalid Time format: %q", s)
}

// MarshalGQL implementa serialización para GraphQL
func (t Time) MarshalGQL(w io.Writer) {
	// GraphQL espera string JSON
	io.WriteString(w, fmt.Sprintf("%q", t.Format("15:04:05")))
}

// UnmarshalGQL implementa parseo desde input GraphQL
func (t *Time) UnmarshalGQL(v any) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("Time must be a string")
	}
	parsed, err := ParseTime(s)
	if err != nil {
		return err
	}
	*t = parsed
	return nil
}
