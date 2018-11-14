package database

// PGSQLCluster represents how a cluster comes back from
// the database
type PGSQLCluster struct {
	Name  string `db:"cluster_name"`
	Color int    `db:"color"`
}
