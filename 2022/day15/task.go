package main

type Task struct {
	Pairs []Pair
}

func (t Task) Clone() *Task {
	pairs := make([]Pair, 0, len(t.Pairs))
	for _, x := range t.Pairs {
		pairs = append(pairs, x.Clone())
	}

	return &Task{
		Pairs: pairs,
	}
}

type Pair struct {
	Sensor Point
	Beacon Point
}

func (p Pair) Clone() Pair {
	return Pair{
		Sensor: p.Sensor,
		Beacon: p.Beacon,
	}
}

type Point struct {
	X, Y int
}
