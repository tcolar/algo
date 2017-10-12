package main

import "strings"

// destinations is a map of direct shipments that originates from a city (city is the key)
// this allows for quickly finding a possible destination from a given location
type destinations map[string][]shipment

// Bundler is used to find a minimum number of bundles from a shipment list
// See README.md for more details
type Bundler struct {
	// destinationsByDay keeps a map of destinations keyed by day of the week for fast lookup
	// this allows for quickly finding destinations on a given day
	destinationsByDay map[weekDay]destinations
}

// NewBundler initializes a Bundler from a shipment list
func NewBundler(shipments []shipment) (*Bundler, error) {
	// Initialize our destinationsByDay map for quick lookup
	destsByDay := map[weekDay]destinations{}
	for _, shipment := range shipments {
		if _, found := destsByDay[shipment.day]; !found {
			destsByDay[shipment.day] = destinations{}
		}
		destsByDay[shipment.day][shipment.from] = append(destsByDay[shipment.day][shipment.from], shipment)
	}
	return &Bundler{
		destinationsByDay: destsByDay,
	}, nil
}

// FindBundles Greedily find bundles among all the possible destinations
// returns a list of bundle strings, as specified in README.md
func (b *Bundler) FindBundles() []string {
	bundles := []string{}
	for day := Monday; day != EOD; day = day.nextDay() {
		for _, routes := range b.destinationsByDay[day] {
			for _, shipment := range routes {
				bundle := b.findBundleGreedily("", shipment.day, shipment.from)
				if len(bundle) > 0 {
					bundles = append(bundles, strings.Trim(bundle, " "))
				}
			}
		}
	}
	return bundles
}

// findBundleGreedily greedily expands a bundle from a given (day/from) location
// this is called recursively until no shipment is found to expand a bundle any further
// Note: As a shipment is added to a bundle, it is consumed (removed from the shipment list).
func (b *Bundler) findBundleGreedily(bundle string, day weekDay, from string) string {
	shipmentFrom, found := b.destinationsByDay[day][from]
	if !found || len(shipmentFrom) == 0 {
		return bundle // No shipment available to grow our bundle from here
	}
	// Greedily pick any shipment that allows us to grow our bundle
	firstShipment := shipmentFrom[0]
	// "consume" the shipment by removing it from the shipment list
	b.destinationsByDay[day][from] = b.destinationsByDay[day][from][1:]
	// add the shipment to our bundle
	bundle += " " + firstShipment.id
	// and recurse to try and continue growing our bundle
	bundle = b.findBundleGreedily(bundle, day.nextDay(), firstShipment.to)
	return bundle
}
