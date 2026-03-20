package models

type User struct {
	ID int
	Name string
	Email string
	Password string
	Age int
	Workouts []Workout
}

type Workout struct {
	ID int
	UserID int
	Name string
	Division string
	Exercices []Exercice
}

type Package struct {
	ID int
	Name string
}

type Exercice struct {
	ID int
	Name string
	Muscle string
	Image string
	Description string
	Categorie string
}

type WorkoutExercise struct {
	WorkoutID int
	ExerciseID int
	Reps int
	Sets int
}

type WorkoutExerciseDetail struct {
	ExerciseID int
	Name string
	Muscle string
	Reps int
	Sets int
}


