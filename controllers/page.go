package controllers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"job/models"
	"job/utils"
	"log"
	"math/rand"
	"strconv"
)

type PageController struct {
	MainController
}

var clientOptions *options.ClientOptions
var client *mongo.Client
var ctx = context.TODO()

func init() {
	clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	// Connect to MongoDB
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func (c *PageController) Post() {
	v := utils.Session.Get("user").(*models.User)
	filter := bson.D{{"mobile", v.Mobile}}
	var pbo []models.PositionCompanyBO
	findOptions := options.Find()
	collection := client.Database("recruit").Collection("position")
	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem models.PositionCompanyBO
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		pbo = append(pbo, elem)
	}

	pagestr := c.GetString(":id")
	page, _ := strconv.Atoi(pagestr)
	page = 1

	if cap(pbo) <= 6*page {
		for i := 0; i < 30; i++ {
			pc, _ := models.ListPosCompany(rand.Int() % 1000)
			pc.SalaryDown = pc.SalaryDown
			pc.SalaryUp = pc.SalaryUp
			pbo = append(pbo, pc)

		}
	}
	pbo = pbo[page*6-6 : page*6]

	if pbo == nil {
		fmt.Println("error occurs")
	}
	out := make(map[string]interface{})
	out["posInfo"] = pbo
	out["user"] = v
	c.Data["json"] = out
	c.ServeJSON()
}

func (c *PageController) GenMGO() {
	user := utils.Session.Get("user").(*models.User)
	var ctx = context.TODO()
	// Connect to MongoDB
	filter := bson.D{{"mobile", user.Mobile}}
	findOptions := options.Find()
	collection := client.Database("recruit").Collection("position")
	if cur, _ := collection.Find(ctx, filter, findOptions); cur != nil {
		if !cur.TryNext(context.TODO()) {
			posinfo := models.RecPosition(*user, 1, 40)
			for _, v := range posinfo {
				v.Mobile = user.Mobile
				insertOne, err := collection.InsertOne(ctx, v)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
			}
		}

	}

	c.Data["json"] = 1
	c.ServeJSON()
}

func (c *PageController) UpdateMGO() {
	user := utils.Session.Get("user").(*models.User)
	var ctx = context.TODO()
	// Connect to MongoDB
	filter := bson.D{{"mobile", user.Mobile}}
	delOptions := options.Delete()
	collection := client.Database("recruit").Collection("position")
	if cur, _ := collection.DeleteMany(ctx,filter,delOptions); cur != nil {
		posinfo := models.RecPosition(*user, 1, 40)
		for _, v := range posinfo {
			v.Mobile = user.Mobile
			insertOne, err := collection.InsertOne(ctx, v)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)
		}
	}

	c.Data["json"] = 1
	c.ServeJSON()
}