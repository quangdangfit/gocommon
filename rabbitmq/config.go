package rabbitmq

type Config struct {
	AMQPUrl      string `json:"amqp_url"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Vhost        string `json:"vhost"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ExchangeName string `json:"exchange_name"`
	ExchangeType string `json:"exchange_type"`
	QueueName    string `json:"queue_name"`
}
