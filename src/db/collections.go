package db

import "go.mongodb.org/mongo-driver/mongo"

var collections *collectionsObj

type collectionsObj struct {
	roleCollection    *mongo.Collection
	userCollection    *mongo.Collection
	productCollection *mongo.Collection
}

func loadCollections(client *mongo.Client) {
	collections = &collectionsObj{
		roleCollection:    client.Database("admin").Collection("roles"),
		userCollection:    client.Database("admin").Collection("users"),
		productCollection: client.Database("erp").Collection("products"),
	}
}
