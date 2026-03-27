package models

type User struct {
	ID int
	Name string
	Email string
	Password string
	Workouts []Workout
}

type Workout struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Division string `json:"division"`
	Exercises []WorkoutExercise `json:"exercises"`
}

type Package struct {
	ID int
	Name string
}

type Exercise struct {
	ID int
	Name string
	Muscle string
	Image string
	Description string
	Categorie string
}

type WorkoutExercise struct {
	WorkoutID int `json:"workout_id"`
	ExerciseID int `json:"exercise_id"`
	Reps int `json:"reps"`
	Sets int `json:"sets"`
}

type WorkoutExerciseDetail struct {
	ExerciseID int
	Name string
	Muscle string
	Reps int
	Sets int
}


