package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// MySQLPersister represents the MySQL implementation of the Persister interface
type MySQLPersister struct {
	db *sql.DB
}

// Migrate creates the packages table and inserts the default package sizes
// TODO: As best practice migrations should be handled by a separate tool that upgrades the DB schema
// up to the latest version. For the sake of simplicity we will handle migrations here.
func (mp *MySQLPersister) Migrate() error {
	// Create the packages table
	_, err := mp.db.Exec("CREATE TABLE packages (packageSize int NOT NULL)")
	if err != nil {
		return err
	}
	_, err = mp.db.Exec("INSERT INTO packages (packageSize) VALUES (250, 500, 1000, 2000, 5000)")
	return err
}

// Init opens a connection to the database
func (mp *MySQLPersister) Init() error {
	var err error
	mp.db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	return err
}

// Add inserts a new packageSize into the database
func (mp *MySQLPersister) Add(data int) (int, error) {
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

// Remove removes a packageSize from the database
func (mp *MySQLPersister) Remove(data int) error {
	// Remove the package from the database
	removePackageQuery := "DELETE FROM packages WHERE packageSize = ?"
	_, err := mp.db.Exec(removePackageQuery, data)
	return err
}

// Close closes the connection to the database
func (mp *MySQLPersister) Close() error {
	return mp.db.Close()
}
