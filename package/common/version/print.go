package version

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const labelWidth = 10

// Logger is satisfied by any structured logger (e.g. zap sugared, slog).
type Logger interface {
	Info(msg string, args ...any)
}

// extraLines renders Extra T into display rows.
// - struct/map → sorted key-value pairs
// - primitive / unserializable → single line
func extraLines(v any) []string {
	b, err := json.Marshal(v)
	if err != nil {
		return []string{fmt.Sprintf("%+v", v)}
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil || len(m) == 0 {
		s := strings.Trim(string(b), `"`)
		if s == "" || s == "null" || s == "{}" {
			return nil
		}
		return []string{s}
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	result := make([]string, 0, len(keys))
	for _, k := range keys {
		result = append(result, fmt.Sprintf("%s\x00%s", k, fmt.Sprint(m[k])))
	}
	return result
}

// String returns a bordered box with all build info.
func (i *Info[T]) String() string {
	title := i.AppName + "  " + i.Version

	allFields := [][2]string{
		{"Version", i.Version},
		{"Branch", i.Branch},
		{"Commit", i.Commit},
		{"Built", i.BuildTime},
		{"Go", i.GoVersion},
	}
	var fields [][2]string
	for _, f := range allFields {
		if f[1] != "" {
			fields = append(fields, f)
		}
	}

	extra := extraLines(i.Extra)

	// box inner width
	boxWidth := len(title) + 4
	for _, f := range fields {
		if w := 2 + labelWidth + 2 + len(f[1]) + 2; w > boxWidth {
			boxWidth = w
		}
	}
	for _, e := range extra {
		parts := strings.SplitN(e, "\x00", 2)
		val := ""
		if len(parts) == 2 {
			val = parts[1]
		} else {
			val = e
		}
		if w := 2 + labelWidth + 2 + len(val) + 2; w > boxWidth {
			boxWidth = w
		}
	}
	if boxWidth < 45 {
		boxWidth = 45
	}

	hline := func() string { return strings.Repeat("─", boxWidth) }

	centerLine := func(s string) string {
		pad := boxWidth - len(s)
		l := pad / 2
		return "│" + strings.Repeat(" ", l) + s + strings.Repeat(" ", pad-l) + "│"
	}

	rowLine := func(label, value string) string {
		inner := fmt.Sprintf("  %-*s  %s", labelWidth, label, value)
		pad := boxWidth - len(inner)
		if pad < 2 {
			pad = 2
		}
		return "│" + inner + strings.Repeat(" ", pad) + "│"
	}

	sepLine := func(label string) string {
		if label == "" {
			return "├" + hline() + "┤"
		}
		dashes := boxWidth - len(label) - 2
		l := dashes / 2
		return "├" + strings.Repeat("─", l) + " " + label + " " + strings.Repeat("─", dashes-l) + "┤"
	}

	var sb strings.Builder
	sb.WriteString("┌" + hline() + "┐\n")
	sb.WriteString(centerLine(title) + "\n")
	sb.WriteString(sepLine("") + "\n")
	for _, f := range fields {
		sb.WriteString(rowLine(f[0], f[1]) + "\n")
	}
	if len(extra) > 0 {
		sb.WriteString(sepLine("extra") + "\n")
		for _, e := range extra {
			parts := strings.SplitN(e, "\x00", 2)
			if len(parts) == 2 {
				sb.WriteString(rowLine(parts[0], parts[1]) + "\n")
			} else {
				sb.WriteString(rowLine("", e) + "\n")
			}
		}
	}
	sb.WriteString("└" + hline() + "┘\n")
	return sb.String()
}

// Print writes the build info box to stdout, centered in the terminal.
func (i *Info[T]) Print() {
	box := i.String()
	tw := termWidth()

	firstLine, _, ok := strings.Cut(box, "\n")
	if !ok {
		fmt.Print(box)
		return
	}
	boxW := len([]rune(firstLine))
	pad := (tw - boxW) / 2
	if pad <= 0 {
		fmt.Print(box)
		return
	}

	prefix := strings.Repeat(" ", pad)
	for _, line := range strings.Split(strings.TrimRight(box, "\n"), "\n") {
		fmt.Println(prefix + line)
	}
}

// Log writes the build info box via the provided logger.
func (i *Info[T]) Log(l Logger) {
	l.Info("\n" + i.String())
}
