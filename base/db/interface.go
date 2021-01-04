package db

type Table interface {
	Find(result, query interface{}, args ...interface{})
	FindOne(id, result interface{})
	Insert(doc interface{})
	// // // InsertMany(docs []interface{}) error
	Join(t1, query string, results interface{})
	Update(doc, query interface{}, args ...interface{})
	// Delete(query interface{}) error
}

type Client interface {
	TableInstance(model interface{}) Table
}
