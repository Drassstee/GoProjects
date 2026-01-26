package migrations

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Run(ctx context.Context, db *pgx.Conn) error {
	query := `
    CREATE TABLE IF NOT EXISTS groups (
        id SERIAL PRIMARY KEY,
        groupname TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS students (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        gender CHAR(1) NOT NULL,
        birth_date DATE NOT NULL,
        group_id INTEGER NOT NULL REFERENCES groups(id)
    );
	CREATE TABLE IF NOT EXISTS instructors (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        gender CHAR(1) NOT NULL,
        birth_date DATE NOT NULL
    );
    CREATE TABLE IF NOT EXISTS subjects (
        id SERIAL PRIMARY KEY,
        subname TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS class_schedule (
        id SERIAL PRIMARY KEY,
        subject TEXT NOT NULL,
		start_time TEXT NOT NULL,
		end_time TEXT NOT NULL,
        group_id INTEGER NOT NULL REFERENCES groups(id),
		subject_id INTEGER NOT NULL REFERENCES subjects(id)
    );
    CREATE TABLE IF NOT EXISTS attendance (
        id SERIAL PRIMARY KEY,
        visit_day DATE NOT NULL,
        visit BOOLEAN NOT NULL,
        student_id INTEGER NOT NULL REFERENCES students(id),
        subject_id INTEGER NOT NULL REFERENCES subjects(id)
    );
    CREATE TABLE IF NOT EXISTS users(
        id SERIAL PRIMARY KEY,
        email TEXT UNIQUE NOT NULL,
		role TEXT NOT NULL,
        password TEXT NOT NULL
    );
	CREATE TABLE IF NOT EXISTS grades(
		id SERIAL PRIMARY KEY,
		grade NUMERIC(4,2) NOT NULL,
		credits INTEGER NOT NULL,
		student_id INTEGER NOT NULL REFERENCES students(id),
        subject_id INTEGER NOT NULL REFERENCES subjects(id)
	);
    `
	_, err := db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}
