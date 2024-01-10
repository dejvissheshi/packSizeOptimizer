package db

import "database/sql"

// MySQLPersister represents the MySQL implementation of the Persister interface
type MySQLPersister struct {
	db *sql.DB
}

func (mp *MySQLPersister) Init() error {
	var err error
	mp.db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	return err
}

// Insert inserts a new packageSize into the database
func (mp *MySQLPersister) Insert(data int) (int, error) {
	// Insert a new user into the database
	insertPackageQuery := "INSERT INTO packages (packageSize) VALUES (?)"
	result, err := mp.db.Exec(insertPackageQuery, data)
	if err != nil {
		return 0, err
	}

	// Retrieve the ID of the inserted user
	packageSize, err := result.LastInsertId()
	return int(packageSize), err
}

// Read retrieves all packageSizes from the database
func (mp *MySQLPersister) Read() ([]int, error) {
	// Retrieve all users from the database
	query := "SELECT packageSize FROM packages"
	rows, err := mp.db.Query(query)
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

func (mp *MySQLPersister) Remove(data int) error {
	// Remove the package from the database
	removePackageQuery := "DELETE FROM packages WHERE packageSize = ?"
	_, err := mp.db.Exec(removePackageQuery, data)
	return err
}

func (mp *MySQLPersister) Close() error {
	return mp.db.Close()
}
