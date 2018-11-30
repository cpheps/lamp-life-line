package database

// ClusterModel represents how a cluster comes back from
// the database
type ClusterModel struct {
	Name  string `db:"cluster_name"`
	Color int    `db:"color"`
}
