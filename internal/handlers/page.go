package handlers

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/a-h/templ"
	"github.com/adrg/frontmatter"
	"github.com/gin-gonic/gin"
	"github.com/thejokers69/thejokers69/internal/config"
	"github.com/thejokers69/thejokers69/internal/models"
	"github.com/thejokers69/thejokers69/templates/pages"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var (
	mdParser = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // GitHub Flavored Markdown
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Allow raw HTML in markdown
		),
	)
)

// LoadPage loads a page from markdown file
func LoadPage(filePath string, site *models.Site) (*models.Page, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var fm models.PageFrontmatter
	rest, err := frontmatter.Parse(strings.NewReader(string(content)), &fm)
	if err != nil {
		return nil, err
	}

	// Convert markdown to HTML
	var buf bytes.Buffer
	if err := mdParser.Convert(rest, &buf); err != nil {
		return nil, err
	}

	// Process HTML to replace Liquid template syntax
	htmlContent := processLiquidSyntax(buf.String(), site)

	page := &models.Page{
		Title:       fm.Title,
		Description: fm.Description,
		Permalink:   fm.Permalink,
		Content:     htmlContent,
		Keywords:    fm.Keywords,
	}

	return page, nil
}

// processLiquidSyntax replaces Jekyll Liquid syntax with Go template syntax
func processLiquidSyntax(html string, site *models.Site) string {
	// Pattern to match: {{ '/path' | relative_url }} or {{ "/path" | relative_url }}
	pattern := regexp.MustCompile(`\{\{\s*['"]([^'"]+)['"]\s*\|\s*relative_url\s*\}\}`)

	html = pattern.ReplaceAllStringFunc(html, func(match string) string {
		// Extract the path from the match
		submatches := pattern.FindStringSubmatch(match)
		if len(submatches) > 1 {
			path := submatches[1]
			// Ensure path starts with /
			if !strings.HasPrefix(path, "/") {
				path = "/" + path
			}
			// Return the path with BaseURL prefix
			return site.BaseURL + path
		}
		return match
	})

	return html
}

// PageHandler handles generic page requests
func PageHandler(c *gin.Context) {
	path := c.Request.URL.Path

	// Remove baseurl prefix if present
	path = strings.TrimPrefix(path, "/thejokers69")

	// Normalize path
	path = strings.Trim(path, "/")
	if path == "" {
		path = "index"
	}

	// Map URL paths to markdown files
	fileMap := map[string]string{
		"index":    "index.md",
		"profile":  "profile.md",
		"about-me": "about-me.md",
		"skills":   "skills.md",
		"projects": "projects.md",
		"contact":  "contact.md",
	}

	fileName, ok := fileMap[path]
	if !ok {
		fileName = path + ".md"
	}

	// Detect BaseURL from request path
	requestPath := c.Request.URL.Path
	baseURL := ""
	if strings.HasPrefix(requestPath, "/thejokers69") {
		baseURL = "/thejokers69"
	}

	site := config.LoadSiteConfigWithBaseURL(baseURL)

	filePath := filepath.Join("content", "pages", fileName)

	page, err := LoadPage(filePath, site)
	if err != nil {
		c.String(404, "Page not found: %s", path)
		return
	}

	// Ensure permalink is set
	if page.Permalink == "" {
		if path == "index" {
			page.Permalink = "/"
		} else {
			page.Permalink = "/" + path + "/"
		}
	}

	// Create a component for the page content using templ.Raw for HTML
	contentComponent := templ.Raw(page.Content)

	// Ensure values are set, using site defaults if empty
	pageTitle := page.Title
	if pageTitle == "" {
		pageTitle = site.Title
	}
	pageDescription := page.Description
	if pageDescription == "" {
		pageDescription = site.Description
	}
	pagePermalink := page.Permalink
	if pagePermalink == "" {
		pagePermalink = "/"
	}

	navHTML := RenderNavigation(site, c.Request.URL.Path, models.GetNavigationItems())
	postHogScript := ""
	if site.PostHog.APIKey != "" {
		postHogScript = `<script>!(function (t, e) {
			var o, n, p, r;
			e.__SV ||
				((window.posthog = e),
				(e._i = []),
				(e.init = function (i, s, a) {
					function g(t, e) {
						var o = e.split(".");
						2 == o.length && ((t = t[o[0]]), (e = o[1])),
							(t[e] = function () {
								t.push([e].concat(Array.prototype.slice.call(arguments, 0)));
							});
					}
					((p = t.createElement("script")).type = "text/javascript"),
						(p.async = !0),
						(p.src = s.api_host + "/static/array.js"),
						(r = t.getElementsByTagName("script")[0]).parentNode.insertBefore(
							p,
							r
						);
					var u = e;
					for (
						void 0 !== a ? (u = e[a] = []) : (a = "posthog"),
							u.people = u.people || [],
							u.toString = function (t) {
								var e = "posthog";
								return (
									"posthog" !== a && (e += "." + a), t || (e += " (stub)"), e
								);
							},
							u.people.toString = function () {
								return u.toString(1) + ".people (stub)";
							},
							o =
								"capture identify alias people.set people.set_once set_config register register_once unregister opt_out_capturing has_opted_out_capturing opt_in_capturing reset isFeatureEnabled onFeatureFlags getFeatureFlag getFeatureFlagPayload reloadFeatureFlags group updateEarlyAccessFeatureEnrollment getEarlyAccessFeatures getActiveMatchingSurveys getSurveys onSessionId".split(
									" "
								),
							n = 0;
						n < o.length;
						n++
					)
						g(u, o[n]);
					e._i.push([i, s, a]);
				}),
				(e.__SV = 1));
		})(document, window.posthog || []);
		posthog.init("` + site.PostHog.APIKey + `", {
			api_host: "` + site.PostHog.Host + `",
			session_recording: { maskAllInputs: true },
		});</script>`
	}
	pages.Base(pageTitle, pageDescription, pagePermalink, site, navHTML, contentComponent, postHogScript).Render(c.Request.Context(), c.Writer)
}
