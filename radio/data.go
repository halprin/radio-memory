package radio

type Memory struct {
	FrequencyRx float64
	FrequencyTx float64
	Tag         string
}

type Radio interface {
	ReadMemories() ([]Memory, error)
	WriteMemories([]Memory) error
}
