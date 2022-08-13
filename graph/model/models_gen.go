// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Workout struct {
	ID        string      `json:"id"`
	Distance  float64     `json:"distance"`
	Duration  float64     `json:"duration"`
	StartDate string      `json:"startDate"`
	Type      WorkoutType `json:"type"`
}

type WorkoutInput struct {
	Distance  float64     `json:"distance"`
	Duration  float64     `json:"duration"`
	StartDate string      `json:"startDate"`
	Type      WorkoutType `json:"type"`
}

type WorkoutType string

const (
	WorkoutTypeCycling WorkoutType = "cycling"
	WorkoutTypeRunning WorkoutType = "running"
)

var AllWorkoutType = []WorkoutType{
	WorkoutTypeCycling,
	WorkoutTypeRunning,
}

func (e WorkoutType) IsValid() bool {
	switch e {
	case WorkoutTypeCycling, WorkoutTypeRunning:
		return true
	}
	return false
}

func (e WorkoutType) String() string {
	return string(e)
}

func (e *WorkoutType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkoutType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkoutType", str)
	}
	return nil
}

func (e WorkoutType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
