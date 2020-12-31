package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type stats struct {
	StartTime           time.Time       `json:"-"`
	Activated           bool            `json:"-"`
	ExecDuration        time.Duration   `json:"exec_duration,omitempty"`
	TotalOperations     uint            `json:"total_operations,omitempty"`
	TimePerOperation    time.Duration   `json:"time_per_operation,omitempty"`
	Operations          map[opcode]uint `json:"operations,omitempty"`
	TotalMemoryAccesses uint            `json:"total_memory_accesses"`
	MemoryAccesses      map[string]uint `json:"memory_accesses"`
}

func (s *stats) String() string {
	textByte, _ := json.MarshalIndent(s, "", "  ")
	text := string(textByte)
	text = strings.ReplaceAll(text, "\"", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, "_", " ")
	return text
}

func newStats() stats {
	return stats{
		Activated:      true,
		Operations:     map[opcode]uint{},
		MemoryAccesses: map[string]uint{},
	}
}

// start the statistic measurements.
func (s *stats) start() {
	s.StartTime = time.Now()
}

// stop the statistic measurements and calculate summary values.
func (s *stats) stop() {
	s.ExecDuration = time.Since(s.StartTime)
	// Count operations
	for _, value := range s.Operations {
		s.TotalOperations += value
	}

	// Count total memory accesses
	for _, value := range s.MemoryAccesses {
		s.TotalMemoryAccesses += value
	}

	// Calculate time per operations
	nanosPerOp := float64(s.ExecDuration.Nanoseconds()) / float64(s.TotalOperations)
	s.TimePerOperation, _ = time.ParseDuration(strconv.FormatFloat(nanosPerOp, 'f', -1, 64) + "ns")

}

func (s *stats) MarshalJSON() ([]byte, error) {
	s.stop()
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
