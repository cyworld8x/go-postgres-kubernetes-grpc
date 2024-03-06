package social

// Store provides all functions to execute do queries and transactions
type Store struct {
	*Queries
	db DBTX
}

//Newstere creates a new Store
func NewStore(db DBTX) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
