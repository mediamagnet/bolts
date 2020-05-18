package lib

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
var err error

type RoleMeListen struct {
	GuildID 	string				`bson:"GuildID,omitempty"`
	ChannelID	string				`bson:"ChannelID,omitempty"`
	RoleID		string				`bson:"RoleID,omitempty"`
	Phrase 		string				`bson:"Phrase,omitempty"`
}

// GetClient
func GetClient() *mongo.Client {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb+srv://mediamagnet:a287593A@cluster0-hjehy.gcp.mongodb.net/test?retryWrites=true&w=majority")
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

// mongoDB connection stuff
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

func MonUpdateListen(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {
	collection := client.Database("bolts").Collection("roleme")
	update := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("Error updating player", err)
	}
	return updatedResult.ModifiedCount
}

func MonReturnAllListen(client *mongo.Client, filter bson.M) []*RoleMeListen {

	var roles []*RoleMeListen
	collection := client.Database("bolts").Collection("listens")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Could not find document ", err)
	}
	for cur.Next(context.TODO()) {
		var roleme RoleMeListen
		err = cur.Decode(&roles)
		if err != nil {
			log.Fatal("Decode Error ", err)
		}
		roles = append(roles, &roleme)
	}
	return roles
}


