package types

type Suggestion struct {
	// value is the thing that we're matching on and the thing that will be submitted if you select the suggestion
	Value string
	// label is what is actually displayed so it can e.g. contain color
	Label string
}
