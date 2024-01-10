package postgres

import "database/sql"

// PostgresPersister represents the PostgreSQL implementation of the Persister interface
type PostgresPersister struct {
	db *sql.DB
}

func (pp *PostgresPersister) Init() error {
	var err error
	pp.db, err = sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	return err
}

// Insert inserts a new packageSize into the database
func (pp *PostgresPersister) Insert(data int) (int, error) {
	// Insert a new user into the database
	insertUserQuery := "INSERT INTO packages (packageSize) VALUES ($1) RETURNING packageSize"
	var packageSize int
	err := pp.db.QueryRow(insertUserQuery, data).Scan(&packageSize)
	return packageSize, err
}

// Read retrieves all packageSizes from the database
func (pp *PostgresPersister) Read() ([]int, error) {
	// Retrieve all users from the database
	query := "SELECT packageSize FROM packages"
	rows, err := pp.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows, adding the packageSize to the slice
	var packageSizes []int
	for rows.Next() {
		var packageSize int
		err := rows.Scan(&packageSize)
		if err != nil {
			return nil, err
		}
		packageSizes = append(packageSizes, packageSize)
	}

	return packageSizes, nil
}

func (pp *PostgresPersister) Remove(data int) error {
	query := "DELETE FROM packages WHERE packageSize = $1"
	_, err := pp.db.Exec(query, data)
	return err
}
