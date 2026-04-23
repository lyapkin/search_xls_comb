package state

type status string

const (
	Loading status = "loading"
	Ready   status = "ready"
	Init    status = "init"
	Error   status = "error"
)
