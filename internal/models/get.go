package models

import "github.com/gattini0928/Learning-Go-Workout-Builder/internal/db"

func GetUserByEmail(email string) (User, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return User{}, err
	}

	defer conn.Close()

	var user User

	query := conn.QueryRow(`SELECT id, email, password FROM users WHERE email = $1`, email)
	err = query.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, err
}

func GetAllWorkouts(user_id int) ([]Workout, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, user_id, name, division FROM workouts WHERE user_id = $1`, user_id)
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

func GetWorkoutByID(workout_id int) (WorkoutResponse, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return WorkoutResponse{}, err
	}
	defer conn.Close()

	var workout WorkoutResponse

	row := conn.QueryRow(`
		SELECT id, user_id, name, division
		FROM workouts
		WHERE id = $1
	`, workout_id)

	err = row.Scan(&workout.ID, &workout.UserID, &workout.Name, &workout.Division)
	if err != nil {
		return WorkoutResponse{}, err
	}

	exercises, err := GetWorkoutExercises(workout_id)
	if err != nil {
		return WorkoutResponse{}, err
	}

	workout.Exercises = exercises

	return workout, nil
}

func GetAllExercices() ([]Exercise, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, name, muscle, image, description, categorie FROM exercises`)
	if err != nil {
		return nil, err
	}

	var exercices []Exercise

	for rows.Next() {
		var exercise Exercise		
		err = rows.Scan(&exercise.ID, &exercise.Name, &exercise.Muscle, &exercise.Image, &exercise.Description,&exercise.Categorie)

		if err != nil {
			return nil, err
		}

		exercices = append(exercices, exercise)
	}

	return exercices, nil
}

func GetExercicesByCategory(category string) ([]Exercise, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
		SELECT id, name, muscle, image, description, categorie 
		FROM exercises 
		WHERE categorie = $1
		`, category)

	if err != nil {
		return nil, err
	}

	var exercices []Exercise

	for rows.Next(){
		var exercise Exercise
		err = rows.Scan(&exercise.ID, &exercise.Name, &exercise.Muscle, &exercise.Image, &exercise.Description, &exercise.Categorie)
		
		if err != nil {
			return nil, err
		}

		exercices = append(exercices, exercise)
	}

	return exercices, nil
}

func GetWorkoutExercises(workoutID int) ([]WorkoutExerciseDetail, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(`
		SELECT e.id, e.name, e.muscle, we.reps, we.sets
		FROM workout_exercises we
		JOIN exercises e ON e.id = we.exercise_id
		WHERE we.workout_id = $1
	`, workoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []WorkoutExerciseDetail

	for rows.Next() {
		var e WorkoutExerciseDetail

		err := rows.Scan(&e.ExerciseID, &e.Name, &e.Muscle, &e.Reps, &e.Sets)
		if err != nil {
			return nil, err
		}

		result = append(result, e)
	}

	return result, nil
}

func GetWorkoutExercise(workoutID int, exerciseID int) (WorkoutExerciseDetail, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return WorkoutExerciseDetail{}, err
	}
	defer conn.Close()

	var e WorkoutExerciseDetail

	err = conn.QueryRow(`
		SELECT e.id, e.name, e.muscle, we.reps, we.sets
		FROM workout_exercises we
		JOIN exercises e ON e.id = we.exercise_id
		WHERE we.workout_id = $1 AND we.exercise_id = $2
	`, workoutID, exerciseID).Scan(
		&e.ExerciseID,
		&e.Name,
		&e.Muscle,
		&e.Reps,
		&e.Sets,
	)

	if err != nil {
		return WorkoutExerciseDetail{}, err
	}

	return e, nil
}