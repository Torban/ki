# Go html/template – rask guide

Dokumentasjon: https://pkg.go.dev/html/template@go1.26.3

## Definer og bruk en template

```html
{{define "navn"}}
  ...innhold...
{{end}}

{{template "navn" .}}
```

## Skriv ut verdier

```html
{{.Felt}}                          <!-- felt på gjeldende kontekst -->
{{.Metode "argument"}}             <!-- kall metode med argument -->
{{formatDate (.Timing.GetDateTime "startTime")}}  <!-- egendefinert funksjon -->
```

## If / else

```html
{{if .Timing.GetString "description"}}
  <span>{{.Timing.GetString "description"}}</span>
{{else}}
  <span>Ingen beskrivelse</span>
{{end}}
```

## Løkke (range)

```html
{{range .Items}}
  <li>{{.Navn}}</li>
{{end}}
```

Inne i `range` er `.` hvert element. Bruk `$` for å nå rot-konteksten:

```html
{{range .Timing.GetStringSlice "tags"}}
  <span>{{.}}</span>
  <button hx-post="/api/timings/{{$.Timing.Id}}/remove-tag">×</button>
{{end}}
```

## Variabler

```html
{{$id := .Timing.Id}}
<div id="timing-{{$id}}">...</div>
```

## Med (with)

Setter ny kontekst hvis verdien ikke er tom:

```html
{{with .Timing.GetString "description"}}
  <span>{{.}}</span>
{{end}}
```

## Unngå HTML-escaping (bruk med forsiktighet)

```html
{{.HtmlInnhold}}   <!-- escapes automatisk – trygt som standard -->
```

Alle verdier escapes automatisk. `html/template` skiller seg fra `text/template` nettopp på dette – det forhindrer XSS.

## Hvor kommer data fra?

Data sendes inn fra Go-handleren som et struct eller et objekt. I tidtaker er `.Timing` et PocketBase `Record`-objekt – én rad fra databasen.

```go
// Go-handler sender recorden til templaten
tmpl.ExecuteTemplate(w, "timing_item", map[string]any{
    "Timing": record,
})
```

```html
<!-- Templaten leser felter med typede metoder på Record -->
{{.Timing.GetBool "isActive"}}
{{.Timing.GetString "description"}}
{{.Timing.GetStringSlice "tags"}}
{{.Timing.GetDateTime "startTime"}}
```

Feltene er definert i databasemigrasjonen (`pb_migrations/timings.go`) og settes i handleren med `record.Set("isActive", true)`.

## Egendefinerte funksjoner

Registreres i Go med `FuncMap` og er tilgjengelig som vanlige kall i templates:

```html
{{formatDuration (.Timing.GetDateTime "startTime") (.Timing.GetDateTime "stopTime")}}
{{localTime (.Timing.GetDateTime "startTime")}}
```
