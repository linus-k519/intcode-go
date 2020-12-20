package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Stats struct {
	ExecDuration    time.Duration   `json:"exec_duration"`
	TotalOperations uint            `json:"total_operations"`
	Operations      map[Opcode]uint `json:"operations"`
}

func (s *Stats) String() string {
	text, _ := json.MarshalIndent(s, "", "  ")
	return string(text)
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
		ExecDuration string          `json:"exec_duration"`
		Operations   map[string]uint `json:"operations"`
		*Alias
	}{
		ExecDuration: s.ExecDuration.String(),
		Operations:   operationsString,
		Alias:        (*Alias)(s),
	})
}
