package pagecss

import "strings"

func NormalizeCSS(css string) string {
	lines := strings.Split(strings.TrimSpace(css), "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return strings.Join(lines, "\n")
}
