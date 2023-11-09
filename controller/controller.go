package controller

import (
	"api/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = // use your database connection string
const dbName = "Netflix"
const collectionName = "watchlist"

// collection pointers
var collection *mongo.Collection

func init() {
	//client option
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection successful!")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection created!")
	fmt.Println(collection)

}

// helper method -- file to insert one movie in database

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie with id :", inserted.InsertedID)

}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated one movie with id :", result.ModifiedCount)
}

func deleteOneRecord(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted one movie with id :", result.DeletedCount)

}

func deleteAllMovies() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all movies ! with count :", result.DeletedCount)

	return result.DeletedCount
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var Movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		Movies = append(Movies, movie)
	}

	defer cursor.Close(context.Background())

	return Movies
}

/* actual controllers --- file */

func GetAllDBMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneRecord(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()

	json.NewEncoder(w).Encode(count)

}
