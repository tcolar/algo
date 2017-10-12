package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// shipment represents an individual shipment as parsed from a shipemnt file
type shipment struct {
	id   string
	from string
	to   string
	day  weekDay
}

// newShipments reads a list of shipments from a shipment file
func newShipments(filePath string) ([]shipment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	shipments := []shipment{}

	ids := map[string]bool{} // used to check for duplicate ids

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		shipment, err := newShipment(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("Error at line %d : %s", lineNumber, err)
		}
		if _, duplicated := ids[shipment.id]; duplicated {
			return nil, fmt.Errorf("Error at line %d : Duplicated id '%s'", lineNumber, shipment.id)
		}
		ids[shipment.id] = true
		shipments = append(shipments, *shipment)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return shipments, nil
}

// newShipment parses a shipment from a single shipment line
func newShipment(shipmentLine string) (*shipment, error) {
	fields := strings.Fields(shipmentLine)
	if len(fields) != 4 {
		return nil, fmt.Errorf("Expected 4 space separated fields in '%s'", shipmentLine)
	}
	day, err := parseWeekDay(fields[3])
	if err != nil {
		return nil, err
	}
	return &shipment{
		id:   fields[0],
		from: fields[1],
		to:   fields[2],
		day:  day,
	}, nil
}
