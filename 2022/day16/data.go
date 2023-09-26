package main

type Data struct {
	Valves []Valve
}

func (d *Data) Clone() *Data {
	newData := Data{
		Valves: make([]Valve, len(d.Valves)),
	}

	for i, v := range d.Valves {
		newData.Valves[i].FlowRate = v.FlowRate
		newData.Valves[i].Name = v.Name
		newData.Valves[i].Tunnels = make([]string, len(v.Tunnels))

		for j, tunnel := range v.Tunnels {
			newData.Valves[i].Tunnels[j] = tunnel
		}
	}

	return &newData
}

type Valve struct {
	Name     string
	FlowRate int
	Tunnels  []string
}
