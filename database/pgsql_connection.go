// Package database handles connection to the database
package database

import (
	"fmt"
	"os"
	// pq is required for postgreSQL driver but isn't used in code
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const databaseENV = "DATABASE_URL"

// PGSQLConnection represents a wrapper around a PostgreSQL connection
type PGSQLConnection struct {
	connection *sqlx.DB
}

// NewConnection creates a new PGSQLConnection by using the URL found in the
// DATABASE_URL environment variable
func NewConnection() (*PGSQLConnection, error) {
	url, ok := os.LookupEnv(databaseENV)
	if !ok {
		return nil, fmt.Errorf("missing ENV %s", databaseENV)
	}

	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PGSQLConnection{
		connection: db,
	}, nil
}

// GetCluster retrieves the cluster associated with clusterName
func (p PGSQLConnection) GetCluster(clusterName string) (*PGSQLCluster, error) {
	cluster := new(PGSQLCluster)
	if err := p.connection.Select(cluster, "SELECT * FROM clusters WHERE cluster_name=$1", clusterName); err != nil {
		return nil, err
	}

	return cluster, nil
}

// GetAllClusters retrieves all clusters in database
func (p PGSQLConnection) GetAllClusters() ([]PGSQLCluster, error) {
	clusters := []PGSQLCluster{}
	if err := p.connection.Select(&clusters, "SELECT * FROM clusters"); err != nil {
		return nil, err
	}

	return clusters, nil
}

// CreateCluster inserts a new cluster into the database
func (p PGSQLConnection) CreateCluster(cluster *PGSQLCluster) error {
	tx, err := p.connection.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO cluster (cluster_name, color) VALUES (:cluster_name, :color)", cluster)
	if err != nil {
		return err
	}

	return tx.Commit()
}
