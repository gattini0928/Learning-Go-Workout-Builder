package models

import "github.com/gattini0928/Build-your-workout-with-Go/db"

func GetAllWorkouts(user_id int) ([]Workout, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, user_id, division FROM workouts WHERE user_id = $1`, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []Workout

	for rows.Next() {
		var w Workout

		err := rows.Scan(&w.ID, &w.UserID, &w.Name, &w.Division)
		if err != nil {
			return nil, err
		}

		workouts = append(workouts, w)
	}

	return workouts, nil
}

func GetWorkoutByID(workout_id int) (Workout, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return Workout{}, err
	}
	defer conn.Close()

	var workout Workout

	row := conn.QueryRow(
		`SELECT id, user_id, name, division FROM workouts WHERE id = $1`,
		workout_id,
	)

	err = row.Scan(&workout.ID, &workout.UserID, &workout.Name, &workout.Division)
	if err != nil {
		return Workout{}, err
	}

	return workout, nil
}