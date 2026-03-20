package models

import "github.com/gattini0928/Build-your-workout-with-Go/db"

func UpdateWorkoutExerciceRepsAndSets(workoutID int, exerciseID int, reps int, sets int) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec(`
		UPDATE workout_exercises
		SET reps = $1, sets = $2
		WHERE workout_id = $3 AND exercise_id = $4
	`, reps, sets, workoutID, exerciseID)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

