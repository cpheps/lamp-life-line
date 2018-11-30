package database

import (
	"reflect"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_PGSQLConnection_GetCluster(t *testing.T) {
	conn, mock := CreateMockSQL(t)

	rows := sqlmock.NewRows([]string{"cluster_name", "color"}).AddRow("myCluster", 1234)
	mock.ExpectQuery(`SELECT \* FROM clusters WHERE cluster_name=myCluster`).WillReturnRows(rows)

	expected := &ClusterModel{
		Name:  "myCluster",
		Color: 1234,
	}

	out, err := conn.GetCluster(expected.Name)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Expected %+v got %+v", expected, out)
	}
}

func Test_PGSQLConnection_GetAllClusters(t *testing.T) {
	conn, mock := CreateMockSQL(t)

	rows := sqlmock.NewRows([]string{"cluster_name", "color"}).
		AddRow("myCluster", 1234).
		AddRow("yourCluster", 4567)
	mock.ExpectQuery(`SELECT \* FROM clusters`).WillReturnRows(rows)

	expected := []ClusterModel{
		{
			Name:  "myCluster",
			Color: 1234,
		},
		{
			Name:  "yourCluster",
			Color: 4567,
		},
	}

	out, err := conn.GetAllClusters()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Expected %+v got %+v", expected, out)
	}
}

func Test_PGSQLConnection_UpdateCluster(t *testing.T) {
	conn, mock := CreateMockSQL(t)

	input := &ClusterModel{
		Name:  "myCluster",
		Color: 1234,
	}

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE clusters.*").
		WithArgs(1234, "myCluster").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := conn.UpdateCluster(input)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
}

func Test_PGSQLConnection_CreateCluster(t *testing.T) {
	conn, mock := CreateMockSQL(t)

	input := &ClusterModel{
		Name:  "myCluster",
		Color: 1234,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO clusters \(cluster_name, color\).*`).
		WithArgs("myCluster", 1234).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := conn.CreateCluster(input)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
}

func Test_PGSQLConnection_DeleteCluster(t *testing.T) {
	conn, mock := CreateMockSQL(t)

	input := &ClusterModel{
		Name:  "myCluster",
		Color: 1234,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`^DELETE FROM clusters.*`).
		WithArgs("myCluster").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := conn.DeleteCluster(input)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
}
