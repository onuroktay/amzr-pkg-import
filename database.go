package OnurTPIjsonImporter

// DATABASE contains an interface with the method to implement in order to use de database
type DATABASE struct {
	// List of Methods to be implemented in the db struct (couchbase, elasticsearch, ...)
	accesser interface {
		SaveItem(*ITEM) error
	}
}
