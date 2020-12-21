package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Stats struct {
	ExecDuration     time.Duration   `json:"exec_duration"`
	TotalOperations  uint            `json:"total_operations"`
	TimePerOperation time.Duration   `json:"time_per_operation"`
	Operations       map[Opcode]uint `json:"operations"`
}

func (s *Stats) String() string {
	text, _ := json.MarshalIndent(s, "", "  ")
	return string(text)
}

func NewStats(operationCount map[Opcode]uint, duration time.Duration) *Stats {
	s := Stats{
		ExecDuration: duration,
		Operations:   operationCount,
	}
	for _, value := range operationCount {
		s.TotalOperations += value
	}
	nanosPerOp := float64(s.ExecDuration.Nanoseconds()) / float64(s.TotalOperations)
	s.TimePerOperation, _ = time.ParseDuration(strconv.FormatFloat(nanosPerOp, 'f', -1, 64) + "ns")
	return &s
}

func (p *Program) CpuInfo() (string, uint) {
	stringCpu := make([]string, len(p.OperationCount))
	var sum uint = 0
	i := 0
	for key, value := range p.OperationCount {
		stringCpu[i] = fmt.Sprintf("%3dx %s", value, key.String())
		sum += value
		i++
	}
	return strings.Join(stringCpu, "\n"), sum
}

func (s *Stats) MarshalJSON() ([]byte, error) {
	operationsString := map[string]uint{}
	for key, value := range s.Operations {
		operationsString[key.String()] = value
	}

	type Alias Stats
	return json.Marshal(&struct {
		ExecDuration     string `json:"exec_duration"`
		TimePerOperation string `json:"time_per_operation"`
		*Alias
		Operations map[string]uint `json:"operations"`
	}{
		ExecDuration:     s.ExecDuration.String(),
		TimePerOperation: s.TimePerOperation.String(),
		Alias:            (*Alias)(s),
		Operations:       operationsString,
	})
}
