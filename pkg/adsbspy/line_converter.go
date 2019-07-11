package adsbspy

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

type Message struct {
	Message          []uint8
	Timestamp        uint32
	TimingResolution uint8
	RSSI             uint16
}

func ReadLine(line string) (*Message, error) {

	if len(line) == 0 {
		return nil, errors.New("line is empty")
	}

	// Example: *8D4BAB4558AB031C446849B72535;1D5D32D0;0A;32AB;
	parts := strings.Split(line, ";")
	if len(parts) < 4 {
		return nil, errors.New("line must have 4 parts")
	}

	if len(parts[0]) < 1 || parts[0][0] != '*' {
		return nil, errors.New("first part of the line must start with '*'")
	}

	decodedPart0, err := hex.DecodeString(parts[0][1:])
	if err != nil {
		return nil, fmt.Errorf("unable to read first part of the line due to: %v", err)
	}

	decodedPart1, err := hex.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("unable to read second part of the line due to: %v", err)
	}

	decodedPart2, err := hex.DecodeString(parts[2])
	if err != nil {
		return nil, fmt.Errorf("unable to read third part of the line due to: %v", err)
	}

	decodedPart3, err := hex.DecodeString(parts[3])
	if err != nil {
		return nil, fmt.Errorf("unable to read fourth part of the line due to: %v", err)
	}

	timestamp := uint32(decodedPart1[0])<<24 + uint32(decodedPart1[1])<<16 + uint32(decodedPart1[2])<<8 + uint32(decodedPart1[3])

	rssi := uint16(decodedPart3[0])<<8 + uint16(decodedPart3[1])

	return &Message{
		Message:          decodedPart0,
		Timestamp:        timestamp,
		TimingResolution: decodedPart2[0],
		RSSI:             rssi,
	}, nil
}
