# adventofcode

[Advent of code](https://adventofcode.com/) solutions.

## 2024

33/50 stars

Languages: Go, Elixir

Solved in December 2024 - January 2025

## 2022

30/50 stars

Languages: Go, Elixir

Second AoC that I've participated in. Solved in December 2022

## 2021

50/50 stars

Languages: Go, Elixir

First AoC that I've participated in. Solved in December 2021

## 2015

28/50 stars

Languages: Elixir, Go

Solved long time after event has ended, when there was no other AoC live event.

## Tools

To help with boilerplate code setup and input file creation I wrote simple tool called gaoc (Go + AoC).

```sh
 go install github.com/SuddenGunter/adventofcode/tools/gaoc@latest
```

### How to use gaoc

```sh
gaoc gen
```

Supported flags:
- -d - day of the AoC (defaults to current day)
- -y - year of the AoC (defaults to current year)
- -t - template - right now only go and demo are supported (defaults to go)
- -o - output - directory where to put generated code (defaults to `day{{.Day}}`)
- -st - AoC session token, if not empty gaoc tries to fetch input and put it into template
