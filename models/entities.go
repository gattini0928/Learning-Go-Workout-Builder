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

// Admin adiciona. saudades django-admin
type Exercice struct {
	ID int
	Name string
	Muscle string
	Image string
	Description string
}

