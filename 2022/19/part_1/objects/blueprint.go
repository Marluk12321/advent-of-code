package objects

import (
	"strconv"
	"strings"
)

type ConstructionCosts map[Robot]Resources

type Blueprint struct {
	Id                int
	ConstructionCosts ConstructionCosts
}

func toNumber(text string) int {
	number, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return number
}

func getId(text string) int {
	fields := strings.Fields(text)
	return toNumber(fields[1])
}

func toResource(text string) Resource {
	resource, found := resourceFromName[text]
	if !found {
		panic(text)
	}
	return resource
}

func toRobot(resource Resource) Robot {
	robot, found := RobotForResource[resource]
	if !found {
		panic(resource)
	}
	return robot
}

func loadCost(fields []string) Resources {
	cost := Resources{}
	for i, word := range fields {
		if word[len(word)-1] == ',' {
			word = word[:len(word)-1]
		}
		resource, found := resourceFromName[word]
		if found {
			amount := toNumber(fields[i-1])
			cost[resource] = amount
		}
	}
	return cost
}

func loadCosts(text string) ConstructionCosts {
	costs := ConstructionCosts{}
	costDescs := strings.Split(text, ".")
	costDescs = costDescs[:len(costDescs) - 1]  // remove empty string after last "."
	for _, costDesc := range costDescs {
		fields := strings.Fields(costDesc)
		robot := toRobot(toResource(fields[1]))
		cost := loadCost(fields[4:])
		costs[robot] = cost
	}
	return costs
}

func MakeBlueprint(text string) Blueprint {
	parts := strings.Split(text, ": ")
	return Blueprint{
		Id:                getId(parts[0]),
		ConstructionCosts: loadCosts(parts[1]),
	}
}
