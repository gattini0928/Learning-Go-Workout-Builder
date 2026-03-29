package models

import "github.com/gattini0928/Learning-Go-Workout-Builder/internal/db"

func DeleteWorkout(workout_id int, user_id int) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM workout_exercises WHERE workout_id = $1`, workout_id)
	if err != nil {
		return 0, err
	}

	res, err := tx.Exec(
		`DELETE FROM workouts WHERE id = $1 AND user_id = $2`,
		workout_id, user_id,
	)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
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