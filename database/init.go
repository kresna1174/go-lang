package database

var connection string

func init() {
	connection = "MYSQL"
}

func GetConn() string {
	return connection
}