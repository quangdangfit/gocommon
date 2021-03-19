package solr

import (
	"encoding/json"
	"fmt"

	gosolr "github.com/vanng822/go-solr/solr"
)

// Solr struct
type Solr struct {
	solr *gosolr.SolrInterface
}

// New Solr object
func New(conf Config) (ISolr, error) {
	solr, err := gosolr.NewSolrInterface(conf.URL, fmt.Sprintf("solr/%s", conf.Core))
	if err != nil {
		return nil, err
	}
	solr.SetBasicAuth(conf.User, conf.Password)

	return &Solr{solr: solr}, nil
}

// Add document to Solr
func (s *Solr) Add(data []byte) error {
	docs := make([]gosolr.Document, 0, 1)

	var newDoc gosolr.Document
	err := json.Unmarshal(data, &newDoc)
	if err != nil {
		return err
	}
	docs = append(docs, newDoc)

	res, err := s.solr.Add(docs, 0, nil)
	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf("solr add: %v", res.Result)
	}

	res, err = s.solr.Commit()
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("solr commit add: %v", res.Result)
	}

	return nil
}

// Update document in Solr
func (s *Solr) Update(data []byte) error {
	docs := make([]gosolr.Document, 0, 1)

	var newDoc gosolr.Document
	err := json.Unmarshal(data, &newDoc)
	if err != nil {
		return err
	}
	docs = append(docs, newDoc)

	res, err := s.solr.Update(docs, nil)
	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf("solr update: %v", res.Result)
	}

	res, err = s.solr.Commit()
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("solr commit update: %v", res.Result)
	}

	return nil
}
