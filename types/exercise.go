package types

type Exercise struct {
	Name        string `json:"name"`
	Kind        string `json:"kind"`
	MuscleGroup string `json:"muscle_group"`
}

func NewExercise(name string, kind string, muscleGroup string) *Exercise {
	return &Exercise{
		Name:        name,
		Kind:        kind,
		MuscleGroup: muscleGroup,
	}
}
