package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Init(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	schema := `
    CREATE TABLE IF NOT EXISTS charge_points (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        latitude REAL NOT NULL,
        longitude REAL NOT NULL,
        status TEXT NOT NULL CHECK (status IN ('AVAILABLE', 'OCCUPIED', 'OFFLINE'))
    );
    CREATE INDEX IF NOT EXISTS idx_latitude ON charge_points(latitude);
    CREATE INDEX IF NOT EXISTS idx_longitude ON charge_points(longitude);
    `
	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
