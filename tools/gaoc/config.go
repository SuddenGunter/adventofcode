package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Day          int
	Year         int
	Template     string
	Output       string
	SessionToken string
	Input        string
}

func (c Config) String() string {
	return fmt.Sprintf("generating template. day: %d, year: %d, template: %s", c.Day, c.Year, c.Template)
}

func InitConfig() (Config, error) {
	fset := flag.NewFlagSet("gen", flag.ExitOnError)
	year := fset.Int("y", time.Now().Year(), "year of the AoC challenge")
	day := fset.Int("d", time.Now().Day(), "day of the AoC challenge")
	template := fset.String("t", "go", "template to use for generated code")
	output := fset.String("o", "", "output directory")
	st := fset.String("st", "", "AoC session token")

	if len(os.Args) < 2 || os.Args[1] != "gen" {
		return Config{}, fmt.Errorf("unsupported mode. To generate code use 'gaoc gen'")
	} else {
		fset.Parse(os.Args[2:])
	}

	return Config{
		Day:          *day,
		Year:         *year,
		Template:     *template,
		Output:       *output,
		SessionToken: *st,
	}, nil
}
