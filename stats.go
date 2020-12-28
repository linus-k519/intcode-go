package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type stats struct {
	ExecDuration     time.Duration   `json:"exec_duration,omitempty"`
	TotalOperations  uint            `json:"total_operations,omitempty"`
	TimePerOperation time.Duration   `json:"time_per_operation,omitempty"`
	Operations       map[opcode]uint `json:"operations,omitempty"`
}

func (s *stats) String() string {
	textByte, _ := json.MarshalIndent(s, "", "")
	text := string(textByte)
	text = strings.ReplaceAll(text, "{\n", "")
	text = strings.ReplaceAll(text, "\n}", "")
	text = strings.ReplaceAll(text, "\"", "")
	text = strings.ReplaceAll(text, ",", "")
	return text
}

func newStats() *stats {
	return &stats{
		Operations: map[opcode]uint{},
	}
}

func (s *stats) calculate() {
	// Count operations
	for _, value := range s.Operations {
		s.TotalOperations += value
	}

	// Calculate time per operations
	nanosPerOp := float64(s.ExecDuration.Nanoseconds()) / float64(s.TotalOperations)
	s.TimePerOperation, _ = time.ParseDuration(strconv.FormatFloat(nanosPerOp, 'f', -1, 64) + "ns")
}

func (s *stats) MarshalJSON() ([]byte, error) {
	s.calculate()

	// Convert operations to string
	operationsString := map[string]uint{}
	for key, value := range s.Operations {
		operationsString[key.String()] = value
	}

	type Alias stats
	return json.Marshal(&struct {
		ExecDuration     string `json:"exec_duration"`
		TimePerOperation string `json:"time_per_operation"`
		*Alias
		Operations map[string]uint `json:"operations,omitempty"`
	}{
		ExecDuration:     s.ExecDuration.String(),
		TimePerOperation: s.TimePerOperation.String(),
		Alias:            (*Alias)(s),
		Operations:       operationsString,
	})
}
