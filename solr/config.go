package solr

type Config struct {
	URL      string `json:"url"`
	Core     string `json:"core"`
	User     string `json:"user"`
	Password string `json:"password"`
}
