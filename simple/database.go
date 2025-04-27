package simple

type Database struct {
	Name string
}

type DatabaseMongoDB Database
type DatabasePostgreSQL Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongodb *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgreSQL, DatabaseMongoDB: mongodb}
}