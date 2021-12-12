package cave

type Cave struct {
	Size        Size
	ConnectedTo []string
}

const (
	Start = "start"
	End   = "end"
)
