// Package database handles connection to the database
package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	// pq is required for postgreSQL driver but isn't used in code
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
func (p PGSQLConnection) GetCluster(clusterName string) (*ClusterModel, error) {
	cluster := new(ClusterModel)
	if err := p.connection.Get(cluster, fmt.Sprintf("SELECT * FROM clusters WHERE cluster_name=%s", clusterName)); err != nil {
		return nil, err
	}

	return cluster, nil
}

// UpdateCluster updates a clusters color in the database
func (p PGSQLConnection) UpdateCluster(cluster *ClusterModel) error {
	tx, err := p.connection.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("UPDATE clusters SET color = :color WHERE cluster_name = :cluster_name", cluster)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// GetAllClusters retrieves all clusters in database
func (p PGSQLConnection) GetAllClusters() ([]ClusterModel, error) {
	clusters := []ClusterModel{}
	if err := p.connection.Select(&clusters, "SELECT * FROM clusters"); err != nil {
		return nil, err
	}

	return clusters, nil
}

// CreateCluster inserts a new cluster into the database
func (p PGSQLConnection) CreateCluster(cluster *ClusterModel) error {
	tx, err := p.connection.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.NamedExec("INSERT INTO clusters (cluster_name, color) VALUES (:cluster_name, :color)", cluster)
	if err != nil {
		return err
	}

	return tx.Commit()
}
