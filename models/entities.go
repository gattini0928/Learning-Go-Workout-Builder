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
	Division string
	Exercices []Exercice
}

// Admin adiciona. saudades django-admin
type Exercice struct {
	ID int
	Muscle string
	Image string
	Description string
}

