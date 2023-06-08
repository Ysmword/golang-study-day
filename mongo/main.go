package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

type Student struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func initDB() error {
	var err error
	option := options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(2 * time.Second)
	client, err = mongo.Connect(context.Background(), option)
	if err != nil {
		return fmt.Errorf("connect mongodb failed: %v", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("ping mongodb failed: %v", err)
	}
	return nil
}

func create(ctx context.Context, collection *mongo.Collection) error {
	doc := Student{Name: "小明", Age: 20}
	res, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("insert one doc failed: %v", err)
	}
	fmt.Println("insert id: %v", res.InsertedID)
	docs := []interface{}{
		Student{Name: "小明1", Age: 20},
		Student{Name: "小明2", Age: 20},
		Student{Name: "小明3", Age: 20},
		Student{Name: "小明4", Age: 20},
		Student{Name: "小明5", Age: 20},
	}
	ress, err := collection.InsertMany(ctx, docs)
	fmt.Println("insert many ids: %v", ress.InsertedIDs)
	return nil
}

func update(ctx context.Context, collection *mongo.Collection) error {
	filter := bson.D{{"name", "小明1"}}
	update := bson.D{{"$inc", bson.D{{"age", 5}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("update one failed: %v", err)
	}
	return nil
}

func query(ctx context.Context, collection *mongo.Collection) error {
	sort := bson.D{{"age", 1}, {"name", -1}}
	findOption := options.Find().SetSort(sort)
	filter := bson.D{{}}
	cursor, err := collection.Find(ctx, filter, findOption)
	if err != nil {
		return fmt.Errorf("find failed: %v", err)
	}
	for cursor.Next(ctx) {
		var doc Student
		err := cursor.Decode(&doc)
		if err != nil {
			return fmt.Errorf("decode doc failed: %v", err)
		}
		fmt.Printf("%+v\n", doc)
	}
	return nil
}

func delete(ctx context.Context, collection *mongo.Collection) error {
	filter := bson.D{{"name", "小明1"}}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete one doc failed: %v", err)
	}
	fmt.Println(res)
	return nil
}

func main() {
	if err := initDB(); err != nil {
		fmt.Println("init db failed", err)
		return
	}
	collection := client.Database("school").Collection("students")
	// err := create(context.Background(), collection)
	// if err != nil {
	// 	fmt.Println("create failed:", err)
	// 	return
	// }

	// err := update(context.Background(), collection)
	// if err != nil {
	// 	fmt.Println("update failed:", err)
	// 	return
	// }

	err := query(context.Background(), collection)
	if err != nil {
		fmt.Println("query failed:", err)
		return
	}

	err = delete(context.Background(), collection)
	if err != nil {
		fmt.Println("delete failed:", err)
		return
	}

	err = query(context.Background(), collection)
	if err != nil {
		fmt.Println("query failed:", err)
		return
	}
}
