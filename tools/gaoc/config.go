package main

import (
	"flag"
	"fmt"
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

func InitConfig() Config {
	year := flag.Int("y", time.Now().Year(), "year of the AoC challenge")
	day := flag.Int("d", time.Now().Day(), "day of the AoC challenge")
	template := flag.String("t", "go", "template to use for generated code")
	output := flag.String("o", "", "output directory")
	st := flag.String("st", "", "AoC session token")

	flag.Parse()

	return Config{
		Day:          *day,
		Year:         *year,
		Template:     *template,
		Output:       *output,
		SessionToken: *st,
	}
}
