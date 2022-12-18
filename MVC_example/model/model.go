package model

//model.go : db에 접속해 데이터를 핸들링, 결과 전달
import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client     *mongo.Client
	colPersons *mongo.Collection
}

type Person struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
	Pnum string `bson:"pnum,omitempty"`
}

func NewModel() (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := "mongodb://127.0.0.1:27017"
	// Connect return *mongo.Client
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("go-mvc")
		r.colPersons = db.Collection("tPerson")
	}

	return r, nil
}

func (p *Model) GetPerson() []Person {

	coll := p.client.Database("go-mvc").Collection("tPerson")
	filter := bson.D{}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var pers []Person
	if err = cursor.All(context.TODO(), &pers); err != nil {
		panic(err)
	}

	for _, result := range pers {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	return pers
}
