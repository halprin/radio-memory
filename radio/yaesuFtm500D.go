package radio

type YaesuFtm500D struct {
	SdCardMemoryPath string
}

func (receiver YaesuFtm500D) ReadMemories() ([]Memory, error) {
	return []Memory{
		{FrequencyRx: 146.52},
		{FrequencyRx: 142.39},
	}, nil
}

func (receiver YaesuFtm500D) WriteMemories(memories []Memory) error {
	return nil
}
