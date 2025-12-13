package handlers

import (
	"html"
	"strings"

	"github.com/thejokers69/thejokers69/internal/models"
)

// RenderNavigation generates HTML for navigation items
func RenderNavigation(site *models.Site, currentPath string, navItems []models.NavigationItem) string {
	var sb strings.Builder
	for _, item := range navItems {
		href := html.EscapeString(site.BaseURL + item.Link)
		name := html.EscapeString(item.Name)
		if currentPath == item.Link {
			sb.WriteString(`<a class="page-link hover-underline" href="`)
			sb.WriteString(href)
			sb.WriteString(`" aria-current="page">`)
			sb.WriteString(name)
			sb.WriteString(`</a>`)
		} else {
			sb.WriteString(`<a class="page-link hover-underline" href="`)
			sb.WriteString(href)
			sb.WriteString(`">`)
			sb.WriteString(name)
			sb.WriteString(`</a>`)
		}
	}
	return sb.String()
}

