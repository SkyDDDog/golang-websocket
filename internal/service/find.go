package service

import (
	"context"
	"demo04/config"
	"demo04/internal/model/ws"
	"demo04/internal/mongo"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     uint   `json:"read"`
	CreateAt int64  `json:"create_at"`
}

func InsertMsg(id, content string, read uint, expire int64) error {
	// 插入到mongoDb中
	collection := mongo.MongoDBClient.Database(config.MongoDatabase).Collection(id)
	comment := ws.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err := collection.InsertOne(context.TODO(), comment)
	return err
}

func FindMany(sendId, id string, time int64, pageSize int) (results []ws.Result, err error) {
	var resultsMe []ws.Trainer
	var resultsYou []ws.Trainer
	sendIdCollection := mongo.MongoDBClient.Database(config.MongoDatabase).Collection(sendId)
	idCollection := mongo.MongoDBClient.Database(config.MongoDatabase).Collection(id)
	sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(), bson.D{{}},
		options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	idTimeCursor, err := idCollection.Find(context.TODO(), bson.D{{}},
		options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	err = sendIdTimeCursor.All(context.TODO(), &resultsYou) // sendId 对面发过来的
	err = idTimeCursor.All(context.TODO(), &resultsMe)      // Id 发给对面的
	log.Println("resultsYou", resultsYou)
	log.Println("resultsMe", resultsMe)
	results, _ = appendAndSort(resultsMe, resultsYou)
	return
}

func appendAndSort(resultMe, resultYou []ws.Trainer) (results []ws.Result, err error) {
	for _, r := range resultMe {
		sendSort := SendSortMsg{ // 构造返回的msg
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := ws.Result{ // 构造返回所有的内容, 包括发送者
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "me",
		}
		results = append(results, result)
	}

	for _, r := range resultYou {
		sendSort := SendSortMsg{ // 构造返回的msg
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		result := ws.Result{ // 构造返回所有的内容, 包括发送者
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "you",
		}
		results = append(results, result)
	}
	return results, err
}
