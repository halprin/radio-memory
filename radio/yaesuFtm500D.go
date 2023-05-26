package radio

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

type YaesuFtm500D struct {
	SdCardMemoryPath string
}

/*
Memories start at 0x202 byte.  Each memory is 0x10 bytes long.
Rx Frequency = 0x00

Tags start at 0x5C00 byte.  Each tag is 0x10 bytes long.
*/

func (receiver YaesuFtm500D) ReadMemories() ([]Memory, error) {
	openFile, err := os.Open(receiver.SdCardMemoryPath)
	if err != nil {
		return nil, err
	}
	defer openFile.Close()

	_, err = openFile.Seek(0x202, 0)
	if err != nil {
		return nil, err
	}

	memories := make([]Memory, 0)

	for {
		memoryBytes := make([]byte, 0x10)
		_, err := openFile.Read(memoryBytes)
		if err != nil {
			return nil, err
		}

		rxFrequencyBytes := memoryBytes[0:4]
		mhz := hex.EncodeToString(rxFrequencyBytes[0:2])
		mhzDecimal := hex.EncodeToString(rxFrequencyBytes[2:4])
		frequencyString := fmt.Sprintf("%s.%s", mhz, mhzDecimal)
		rxFrequency, err := strconv.ParseFloat(frequencyString, 64)
		if err != nil {
			return nil, err
		}

		if rxFrequency == 144.000 {
			break
		}

		memories = append(memories, Memory{
			FrequencyRx: rxFrequency,
		})

//		break
	}

	return memories, nil
}

func (receiver YaesuFtm500D) WriteMemories(memories []Memory) error {
	return nil
}
