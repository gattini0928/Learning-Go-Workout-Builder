package models

import "github.com/gattini0928/Build-your-workout-with-Go/db"

func DeleteWorkout(workout_id int, user_id int) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(
		`DELETE FROM workouts WHERE id = $1 AND user_id = $2`,
		workout_id, user_id,
	)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DeleteExercice(exercise_id int) (int64, error){
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM exercises WHERE id = $1`, exercise_id)
	if err != nil {
		return 0, nil
	}
	return res.RowsAffected()
}

func DeleteExerciseFromWorkout(workout_id int, exercise_id int) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM workout_exercises WHERE workout_id = $1 and exercise_id = $2`, workout_id, exercise_id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}