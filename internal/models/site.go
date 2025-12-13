package models

// Site represents site-wide configuration
type Site struct {
	Title       string
	Description string
	BaseURL     string
	URL         string
	Author      Author
	Social      SocialLinks
	PostHog     PostHogConfig
}

// Author represents the site author information
type Author struct {
	Name    string
	Email   string
	GitHub  string
	LinkedIn string
}

// SocialLinks represents social media links
type SocialLinks struct {
	Name string
	Links []string
}

// PostHogConfig represents PostHog analytics configuration
type PostHogConfig struct {
	APIKey string
	Host   string
}

// NavigationItem represents a navigation menu item
type NavigationItem struct {
	Link string
	Name string
}

// GetNavigationItems returns the navigation items for the site
func GetNavigationItems() []NavigationItem {
	return []NavigationItem{
		{Link: "/profile/", Name: "Profile"},
		{Link: "/about-me/", Name: "About"},
		{Link: "/skills/", Name: "Skills"},
		{Link: "/projects/", Name: "Projects"},
		{Link: "/contact/", Name: "Contact"},
	}
}

