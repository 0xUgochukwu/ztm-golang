//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) Average() Bytes {
	total := Bytes(0)

	for _, amount := range b.amount {
		total += amount
	}

	return total / Bytes(len(b.amount))
}

type CpuTemp struct {
	temp []Celcius
}

func (c *CpuTemp) Average() Celcius {
	total := Celcius(0)
	for _, temp := range c.temp {
		total += temp
	}
	return total / Celcius(len(c.temp))
}

type MemoryUsage struct {
	amount []Bytes
}

func (m *MemoryUsage) Average() Bytes {
	total := Bytes(0)
	for _, amount := range m.amount {
		total += amount
	}
	return total / Bytes(len(m.amount))
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (d *Dashboard) PrintAverages() {
	fmt.Printf("Bandwidth: %v\n", d.BandwidthUsage.Average())
	fmt.Printf("CPU Temp: %v\n", d.CpuTemp.Average())
	fmt.Printf("Memory Usage: %v\n", d.MemoryUsage.Average())
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{bandwidth, temp, memory}
	dashboard.PrintAverages()
}
