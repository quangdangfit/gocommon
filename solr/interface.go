package solr

// ISolr interface
type ISolr interface {
	Add(data []byte) error
	Update(data []byte) error
}
