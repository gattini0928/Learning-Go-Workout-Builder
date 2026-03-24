package models

import "github.com/gattini0928/Learning-Go-Workout-Builder/db"

func InsertUser(user User) (int, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0 , err
	}

	defer conn.Close()

	var id int
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func InsertWorkout(workout Workout) (int, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	var id int

	query := `INSERT INTO workouts (user_id, name, division) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(query, workout.UserID, workout.Name, workout.Division).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func InsertExercice(exercise Exercise) (int, error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	var id int

	query := `INSERT INTO exercices (muscle, image, description, categorie) VALUES ($1, $2, $3, $4) RETURNING id`
	err = conn.QueryRow(query, exercise.Muscle, exercise.Image, exercise.Description, exercise.Categorie).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}