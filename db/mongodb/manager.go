package mongodb

import "gopkg.in/mgo.v2"

// CollectionManager provides utility interface methods for managing collections in a mongo database
type CollectionManager interface {
	GetCollection() *mgo.Collection
	Close()
}

// ResourceManager provides concrete implementation of methods for interacting with mongo databases
type ResourceManager struct {
	Session        *mgo.Session
	DB             *mgo.Database
	CollectionName string
}

// NewResourceManager constructs a manager object for one of the Collections in the mongo DB.
// Management of the collection is delegated to the given CollectionManager
func NewResourceManager(collectionName string) (*ResourceManager, error) {
	session, db, err := newSession()
	if err != nil {
		return nil, err
	}

	rm := ResourceManager{
		Session:        session,
		DB:             db,
		CollectionName: collectionName,
	}
	return &rm, nil
}

// GetCollection returns the collection being operated on by this manager
func (r *ResourceManager) GetCollection() *mgo.Collection {
	return r.DB.C(r.CollectionName)
}

// Close should be called before the ResourceManager is garbage collected
// The underlying resources will be released cleanly by closing the session and logging out of the database
func (r *ResourceManager) Close() {
	r.Session.Close()
	r.DB.Logout()
}
