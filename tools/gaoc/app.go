package main

import (
	"context"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

type App struct {
}

func (a *App) Run() {
	dir, err := embedfs.ReadDir("template")
	if err != nil {
		a.Fatal(err)
	}

	names := make(map[string]struct{}, len(dir))
	for _, f := range dir {
		names[f.Name()] = struct{}{}
	}

	fmt.Println("supported templates:", names)

	cfg := InitConfig()
	if _, ok := names[cfg.Template]; !ok {
		a.Fatal(fmt.Errorf("template %s not found", cfg.Template))
	}

	if cfg.Output == "" {
		cfg.Output = fmt.Sprintf("day%d", cfg.Day)
	}

	fmt.Println(cfg)

	if cfg.SessionToken != "" {
		cfg.Input = a.getInput(cfg.Day, cfg.Year, cfg.SessionToken)
	}

	a.write(embedfs, cfg)
}

func (a *App) Fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func (a *App) write(f embed.FS, cfg Config) {
	err := os.Mkdir(cfg.Output, 0755)
	if err != nil {
		a.Fatal(err)
	}

	err = os.Chdir(cfg.Output)
	if err != nil {
		a.Fatal(err)
	}

	prefix := path.Join("template", cfg.Template)
	storedFiles := make([]string, 0)
	err = fs.WalkDir(f, prefix, func(virtPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if virtPath == prefix {
			return nil
		}

		realPath := virtPath[len(prefix)+1:] // +1 to remove the slash

		if realPath == "" {
			return nil
		}

		if d.IsDir() {
			return os.Mkdir(realPath, 0755)
		}

		// Create the file
		file, err := os.Create(realPath)
		if err != nil {
			return err
		}
		defer file.Close()

		embedFile, err := f.Open(virtPath)
		if err != nil {
			return err
		}
		defer embedFile.Close()

		_, err = io.Copy(file, embedFile)
		if err != nil {
			return err
		}

		storedFiles = append(storedFiles, realPath)

		return nil
	})
	if err != nil {
		a.Fatal(err)
	}

	fmt.Println("created template files:", storedFiles)

	onlyTmpl := func(path string) bool {
		return strings.HasSuffix(path, ".tmpl")
	}

	storedFiles = filter(storedFiles, onlyTmpl)

	tmplts, err := template.New("aoc").ParseFiles(storedFiles...)
	if err != nil {
		a.Fatal(err)
	}

	for _, tmpl := range tmplts.Templates() {
		if tmpl.Name() == "aoc" {
			continue
		}

		out, err := os.Create(strings.TrimSuffix(tmpl.Name(), ".tmpl"))
		if err != nil {
			a.Fatal(err)
		}
		defer out.Close()

		err = tmpl.Execute(out, cfg)
		if err != nil {
			a.Fatal(err)
		}
	}

	for _, f := range storedFiles {
		err = os.Remove(f)
		if err != nil {
			a.Fatal(err)
		}
	}
}

func (a *App) getInput(day, year int, token string) string {
	timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(timeout, http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		a.Fatal(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: token,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		a.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		a.Fatal(fmt.Errorf("unexpected status code: %d", resp.StatusCode))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		a.Fatal(err)
	}

	return string(data)
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
