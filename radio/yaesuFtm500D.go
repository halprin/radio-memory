package radio

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Memories start at 0x202 byte.  Each memory is 0x10 bytes long.
Rx Frequency = 0x00

Tags start at 0x5C00 byte.  Each tag is 0x10 bytes long.
*/

const memoryWidth = 0x10
const emptyMemoryFrequency = 144.000 //144.000 seems to be the default frequency for empty frequencies

type YaesuFtm500D struct {
	SdCardMemoryPath string
}

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
		memoryBytes := make([]byte, memoryWidth)
		_, err := openFile.Read(memoryBytes)
		if err != nil {
			return nil, err
		}

		rxFrequencyBytes := memoryBytes[0:4]
		mhz := hex.EncodeToString(rxFrequencyBytes[0:2])
		mhz = mhz[1:4] //trim off the first character because the frequency begins at the second character of this first hex byte
		mhzDecimal := hex.EncodeToString(rxFrequencyBytes[2:4])
		mhzDecimal = mhzDecimal[0:3] //trim off the last character because the frequency ends at the first character of this last hex byte
		log.Printf("Reading %s.%s", mhz, mhzDecimal)
		frequencyString := fmt.Sprintf("%s.%s", mhz, mhzDecimal)
		rxFrequency, err := strconv.ParseFloat(frequencyString, 64)
		if err != nil {
			return nil, err
		}

		if rxFrequency == emptyMemoryFrequency {
			//no more freqencies to read
			break
		}

		memories = append(memories, Memory{
			FrequencyRx: rxFrequency,
		})
	}

	return memories, nil
}

func (receiver YaesuFtm500D) WriteMemories(memories []Memory) error {
	return nil
}
