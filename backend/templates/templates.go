package templates

import (
	"embed"
	"fmt"
)

var templateFS embed.FS

func GetTemplateByType(pageType string) string {
	file := fmt.Sprintf("html/%s.svelte", pageType)
	data, err := templateFS.ReadFile(file)
	if err != nil {
		return ""
	}
	return string(data)
}
