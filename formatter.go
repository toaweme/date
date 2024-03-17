package date

import (
	"fmt"
	"strings"
	"time"
)

// Formatter is a struct that holds the time and the mapping of the date format attributes
type Formatter struct {
	time    time.Time
	mapping map[string]string
}

// NewFormatter returns a new Formatter
func NewFormatter(time time.Time, mapping map[string]string) *Formatter {
	return &Formatter{time: time, mapping: mapping}
}

// Render renders our date format string to a Go time layout
func (d *Formatter) Render(format string) (string, error) {
	var goLayout strings.Builder
	for i := 0; i < len(format); i++ {
		character := string(format[i])
		layout, ok := d.mapping[character]
		// we keep the character if it does not exist in the mapping
		if !ok {
			// err := fmt.Errorf("skipping character '%s' invalid date format attribute", string(format[i]))
			// log.Trace().Err(err).Msg("failed to apply format filter")
			goLayout.WriteString(character)
			continue
		}
		// we keep the character if it does not have a direct mapping to a date format attribute
		if layout == "" {
			// err := fmt.Errorf("character '%s' does not have a direct mapping to a date format attribute", string(format[i]))
			// log.Trace().Err(err).Msg("failed to apply format filter")
			goLayout.WriteString(character)
			continue
		}

		goLayout.WriteString(layout)
	}

	// validate if the layout is valid
	_, err := time.Parse(goLayout.String(), goLayout.String())
	if err != nil {
		return "", fmt.Errorf("failed to render date format: '%s': %w", goLayout.String(), err)
	}

	return goLayout.String(), nil
}
