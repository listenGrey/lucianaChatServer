package model

type QA struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Chat struct {
	ChatID int64  `json:"cid" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	QAs    []QA   `json:"qa_s" bson:"qa_s"`
}

type UserChats struct {
	UserID int64 `bson:"_id"`
	Chats  []struct {
		ChatID int64  `bson:"c_id"`
		Name   string `bson:"name"`
	}
}
