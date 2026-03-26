CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(150) UNIQUE,
    password TEXT
);

CREATE TABLE IF NOT EXISTS workouts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    name VARCHAR(30),
    division VARCHAR(20),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE packages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS exercises (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    muscle VARCHAR(30),
    image TEXT,
    description VARCHAR(300),
    categorie VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS workout_exercises (
    workout_id INTEGER,
    exercise_id INTEGER,
    reps Integer,
    sets Integer,
    FOREIGN KEY (workout_id) REFERENCES workouts(id),
    FOREIGN KEY (exercise_id) REFERENCES exercises(id)
);

