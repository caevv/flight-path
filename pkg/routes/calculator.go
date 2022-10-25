package routes

import (
	"errors"
)

func Calculate(flightRoutes []FlightRoute) (FlightRoute, error) {
	if len(flightRoutes) == 1 {
		if flightRoutes[0][0] == flightRoutes[0][1] {
			return nil, errors.New("same path")
		}
		return flightRoutes[0], nil
	}

	from := make(map[string]struct{})
	to := make(map[string]struct{})

	// set all flights origin and destinations
	for _, flightRoute := range flightRoutes {
		from[flightRoute[0]] = struct{}{}
		to[flightRoute[1]] = struct{}{}
	}

	var foundFrom string
	var foundTo string

	// search for a missing origin match on destination
	for fromAirport := range from {
		if _, ok := to[fromAirport]; !ok {
			foundFrom = fromAirport
			break
		}
	}

	// search for a missing destination match on the origin
	for toAirport := range to {
		if _, ok := from[toAirport]; !ok {
			foundTo = toAirport
			break
		}
	}

	// if missing destination or origin match is not found, this is an invalid request
	if foundFrom == "" || foundTo == "" {
		return nil, errors.New("invalid path")
	}

	return FlightRoute{foundFrom, foundTo}, nil
}
