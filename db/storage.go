package db

// start connections with database and distributed cache
func Open() {
	openDatabase()
	openCache()
}

// close connections with database and distributed cache
func Close() {
	closeDatabase()
	closeCache()
}
