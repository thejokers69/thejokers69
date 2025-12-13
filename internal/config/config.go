package config

import (
	"os"

	"github.com/thejokers69/thejokers69/internal/models"
)

// LoadSiteConfig loads site configuration from environment variables or defaults
func LoadSiteConfig() *models.Site {
	baseURL := os.Getenv("BASEURL")
	if baseURL == "" {
		// Default to empty for local development, set BASEURL=/thejokers69 for GitHub Pages
		baseURL = ""
	}

	return LoadSiteConfigWithBaseURL(baseURL)
}

// LoadSiteConfigWithBaseURL loads site configuration with a specific baseURL
func LoadSiteConfigWithBaseURL(baseURL string) *models.Site {

	url := os.Getenv("URL")
	if url == "" {
		url = "https://thejokers69.github.io"
	}

	postHogAPIKey := os.Getenv("POSTHOG_API_KEY")
	postHogHost := os.Getenv("POSTHOG_HOST")
	if postHogHost == "" {
		postHogHost = "https://eu.i.posthog.com"
	}

	return &models.Site{
		Title:       "Mohamed Lakssir | TheJokers69",
		Description: "Computer Engineering Student & Full-Stack Developer specializing in React, Spring Boot, and Data Visualization",
		BaseURL:     baseURL,
		URL:         url,
		Author: models.Author{
			Name:     "Mohamed Lakssir",
			Email:    "contact@mohamedlakssirt.tech",
			GitHub:   "thejokers69",
			LinkedIn: "mohamed-lakssir-b6b4301a4",
		},
		Social: models.SocialLinks{
			Name: "Mohamed Lakssir",
			Links: []string{
				"https://github.com/thejokers69",
				"https://www.linkedin.com/in/mohamed-lakssir-b6b4301a4",
			},
		},
		PostHog: models.PostHogConfig{
			APIKey: postHogAPIKey,
			Host:   postHogHost,
		},
	}
}
