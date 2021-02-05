package travelerJson

import (
	travellerParse "github.ccs.neu.edu/CS4500-S21/Ormegland/A3/traveller-client/parse"
)

func GetUniqueTowns(roads travellerParse.RoadArray) []string {
	var uniqueTowns []string
	//TODO: see if we can get better than exponential
	for _, road := range roads {
		shouldAddTo := true
		shouldAddFrom := true
		for _, uniqueTown := range uniqueTowns {
			shouldAddTo = shouldAddTo && uniqueTown != road.To
			shouldAddFrom = shouldAddFrom && uniqueTown != road.From
		}
		if shouldAddTo {
			uniqueTowns = append(uniqueTowns, road.To)
		}
		if shouldAddFrom {
			uniqueTowns = append(uniqueTowns, road.From)
		}
	}
	return uniqueTowns
}
F