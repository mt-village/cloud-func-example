package example

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
)

// User JSONデコード用の構造体.
type User struct {
	Name string `json:"name"`
}

// HelloPubSub Pub/Subからトリガーされる関数.
func HelloPubSub(ctx context.Context, m *pubsub.Message) error {

	// メッセージの出力
	var user User
	if err := json.Unmarshal(m.Data, &user); err != nil {
		log.Fatal(err)
	}
	log.Println("Hello,", user.Name)

	// 属性の出力
	if len(m.Attributes) > 0 {
		for k, v := range m.Attributes {
			log.Println("key :", k)
			log.Println("value :", v)
		}
	}
	return nil
}
