package db

// Table ...
type Table interface {
	Find(result, query interface{}, args ...interface{})
	FindOne(id, result interface{})
	Insert(doc interface{})
	// InsertMany(docs []interface{}) error
	Join(t1, query string, results interface{}, args ...interface{})
	Update(doc, query interface{}, args ...interface{})
	// Delete(query interface{}) error
}

//Client ...
type Client interface {
	TableInstance(model interface{}) Table
}
