// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewWorkout struct {
	Distance  float64 `json:"distance"`
	Duration  float64 `json:"duration"`
	StartDate float64 `json:"startDate"`
	EndDate   float64 `json:"endDate"`
}

type Workout struct {
	Distance   float64 `json:"distance"`
	Duration   float64 `json:"duration"`
	StartDate  float64 `json:"startDate"`
	EndDate    float64 `json:"endDate"`
	TestNumber float64 `json:"testNumber"`
}
