package example

import (
	"encoding/json"
	"fmt"

	"github.com/quangdangfit/gocommon/solr"
)

func Solr() {
	var conf = solr.Config{
		URL:      "http://localhost:8983/",
		Core:     "core",
		User:     "user",
		Password: "password",
	}
	s, err := solr.New(conf)
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]interface{}{
		"key": "value",
	}

	b, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
	}

	if err := s.Add(b); err != nil {
		fmt.Println(err)
	}

}
