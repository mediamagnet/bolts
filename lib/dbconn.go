package lib

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var err error

// RoleMeListen blah
type RoleMeListen struct {
	GuildID   string `bson:"GuildID,omitempty"`
	ChannelID string `bson:"ChannelID,omitempty"`
	RoleID    string `bson:"RoleID,omitempty"`
	IgnoreID  string `bson:"IgnoreID,omitempty"`
	Phrase    string `bson:"Phrase,omitempty"`
}

// GetClient blah
func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// MonListen mongoDB connection stuff
func MonListen(dbase string, collect string, listens RoleMeListen) {
	// Connecting to mongoDB
	client := GetClient()
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	collection := client.Database(dbase).Collection(collect)
	insertResult, err := collection.InsertOne(context.TODO(), listens)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted:", insertResult.InsertedID)
}

// MonUpdateListen blah
func MonUpdateListen(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {
	collection := client.Database("bolts").Collection("roleme")
	update := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("Error updating player", err)
	}
	return updatedResult.ModifiedCount
}

// MonReturnOneListen blah
func MonReturnOneListen(client *mongo.Client, filter bson.M) RoleMeListen {
	var phrase RoleMeListen
	collection := client.Database("bolts").Collection("listens")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&phrase)
	return phrase
}

// MonReturnAllListen blah
func MonReturnAllListen(client *mongo.Client, filter bson.M) []*RoleMeListen {
	var roles []*RoleMeListen
	collection := client.Database("bolts").Collection("listens")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error finding all the things", err)
	}
	for cur.Next(context.TODO()) {
		var role RoleMeListen
		err = cur.Decode(&role)
		if err != nil {
			log.Fatal("Error Decoding :( ", err)
		}
		roles = append(roles, &role)
	}
	return roles
}
