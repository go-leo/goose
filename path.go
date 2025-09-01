package goose

import "strings"

func URLPath(path string, pairs map[string]string) string {
	sections := strings.Split(path, "/")
	for i := 0; i < len(sections); i++ {
		section := sections[i]
		if section == "" {
			continue
		}
		if !strings.HasPrefix(section, "{") {
			continue
		}
		if !strings.HasSuffix(section, "}") {
			continue
		}
		if strings.HasSuffix(section, "...}") {
			section = section[1 : len(section)-4]
		} else {
			section = section[1 : len(section)-1]
		}
		sections[i] = pairs[section]
	}
	return strings.Join(sections, "/")
}
