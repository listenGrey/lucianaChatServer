package conf

var (
	KafkaServerAddress = "localhost:9092"
	GrpcServerAddress  = "localhost:8964"
	DBAddress          = "mongodb://172.17.0.2:27017"
	Database           = "luciana_chat_db"
	Collection         = "chats"
	CertFile           = "./ca/server.crt"
	KeyFile            = "./ca/server.key"
)
