package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"math"
	"net/http"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

//go:embed templates
var templateFS embed.FS

//go:embed static
var staticFS embed.FS

var pageTemplates map[string]*template.Template

var funcMap = template.FuncMap{
	"formatDuration": formatDuration,
	"formatTime":     formatTime,
	"formatDate":     formatDate,
	"sub":            func(a, b int) int { return a - b },
	"add":            func(a, b int) int { return a + b },
	"seq":            func(n int) []int { s := make([]int, n); for i := range s { s[i] = i }; return s },
	"dict": func(pairs ...any) map[string]any {
		m := make(map[string]any, len(pairs)/2)
		for i := 0; i < len(pairs)-1; i += 2 {
			key, _ := pairs[i].(string)
			m[key] = pairs[i+1]
		}
		return m
	},
	"contains": func(slice []string, item string) bool {
		for _, s := range slice {
			if s == item {
				return true
			}
		}
		return false
	},
	"eq": func(a, b string) bool { return a == b },
	"localTime": func(dt types.DateTime) string {
		if dt.IsZero() {
			return ""
		}
		return dt.Time().Format("2006-01-02T15:04")
	},
}

func initTemplates() {
	pageTemplates = make(map[string]*template.Template)

	partialFiles := []string{
		"templates/timer/timer_controls.html",
		"templates/timer/timings_list.html",
		"templates/timer/timing_item.html",
		"templates/timer/timing_edit.html",
		"templates/search/search_bar.html",
		"templates/search/filter_bar.html",
	}

	// Build a base template set that includes base layout + all partials
	baseWithPartials := template.Must(
		template.New("").Funcs(funcMap).ParseFS(
			templateFS,
			append([]string{"templates/base.html"}, partialFiles...)...,
		),
	)

	// For each page, clone the base+partials and add page-specific template
	pages := map[string]string{
		"login":    "templates/auth/login.html",
		"register": "templates/auth/register.html",
		"timer":    "templates/timer/timer_page.html",
	}

	for name, page := range pages {
		t := template.Must(template.Must(baseWithPartials.Clone()).ParseFS(templateFS, page))
		pageTemplates[name] = t
	}

	// Partial-only template set for HTMX partial responses
	partials := template.Must(
		template.New("").Funcs(funcMap).ParseFS(templateFS, partialFiles...),
	)
	pageTemplates["timer_controls"] = partials
	pageTemplates["timings_list"] = partials
	pageTemplates["timing_item"] = partials
}

func renderPage(e *core.RequestEvent, name string, data any) error {
	t, ok := pageTemplates[name]
	if !ok {
		return e.InternalServerError("Template not found: "+name, nil)
	}

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "base", data); err != nil {
		return e.InternalServerError("Template error: "+err.Error(), nil)
	}

	e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	e.Response.WriteHeader(http.StatusOK)
	_, err := buf.WriteTo(e.Response)
	return err
}

func renderPartial(e *core.RequestEvent, name, block string, data any) error {
	t, ok := pageTemplates[name]
	if !ok {
		return e.InternalServerError("Template not found: "+name, nil)
	}

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, block, data); err != nil {
		return e.InternalServerError("Template error: "+err.Error(), nil)
	}

	e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	e.Response.WriteHeader(http.StatusOK)
	_, err := buf.WriteTo(e.Response)
	return err
}

// renderToString renders a named template block to a string
func renderToString(name, block string, data any) (string, error) {
	t, ok := pageTemplates[name]
	if !ok {
		return "", fmt.Errorf("template not found: %s", name)
	}
	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, block, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func getStaticFS() fs.FS {
	sub, _ := fs.Sub(staticFS, "static")
	return sub
}

func formatDuration(start, stop types.DateTime) string {
	if start.IsZero() {
		return "0s"
	}
	end := stop.Time()
	if stop.IsZero() {
		end = time.Now()
	}
	d := end.Sub(start.Time())
	totalSec := int(math.Floor(d.Seconds()))
	h := totalSec / 3600
	m := (totalSec % 3600) / 60
	s := totalSec % 60

	if h > 0 {
		return fmt.Sprintf("%dt %02dm %02ds", h, m, s)
	}
	if m > 0 {
		return fmt.Sprintf("%dm %02ds", m, s)
	}
	return fmt.Sprintf("%ds", s)
}

func formatTime(dt types.DateTime) string {
	if dt.IsZero() {
		return "-"
	}
	return dt.Time().Format("15:04:05")
}

var norwegianMonths = [...]string{
	"januar", "februar", "mars", "april", "mai", "juni",
	"juli", "august", "september", "oktober", "november", "desember",
}

func formatDate(dt types.DateTime) string {
	if dt.IsZero() {
		return "-"
	}
	t := dt.Time()
	return fmt.Sprintf("%d. %s %d", t.Day(), norwegianMonths[t.Month()-1], t.Year())
}
