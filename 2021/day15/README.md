# Day15

[Challenge link](https://adventofcode.com/2021/day/15).

How to run:

Put your data into 'data.txt' (in the same directory as this readme file).

```sh
go build -o app .
./app
```

## Explanation
Task1/solver is using simple version of Dijkstra's algorithm with O(n^2) complexity.

Task2/solver is optimized with min-heap.

Even though Task1/solver is not optimized, I tested it with input for Part2 - it solved it on Ryzen 3700x in about an hour.