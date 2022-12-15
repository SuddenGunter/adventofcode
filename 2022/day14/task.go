package main

type Task struct {
	Walls     []Wall
	SandSpawn Point
}

func (t Task) Clone() *Task {
	walls := make([]Wall, 0, len(t.Walls))
	for _, x := range t.Walls {
		walls = append(walls, x.Clone())
	}

	return &Task{
		Walls:     walls,
		SandSpawn: t.SandSpawn,
	}
}
