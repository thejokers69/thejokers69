package models

// Page represents a page with its metadata and content
type Page struct {
	Title       string
	Description string
	Permalink   string
	Content     string // HTML content after markdown rendering
	Keywords    string
}

// PageFrontmatter represents the frontmatter from markdown files
type PageFrontmatter struct {
	Layout      string `yaml:"layout"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Permalink   string `yaml:"permalink"`
	Keywords    string `yaml:"keywords"`
}

