package controller

import (
	"fmt"
	"log"
	"context"
	
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"

)

const connectionString = "mongodb+srv://abhijit:rk110295@cluster0.j88jka3.mongodb.net/sensor_data"
const dbName  = "netflix"
const colName = "watchlist"
var collection *mongo.Collection

func init(){
	clientOption := option.Client().ApplyURI(connectionString)

	client, err = mongo.Connect(context.TODO(), clientOption)

	if err != nil{
		log.Fatal(err)
     }

	fmt.Println("Mongodb Connected")

	collection = client.Database(dbName).collection(colName)

	fmt.Println("Collection instance is ready")
}

// MongoDB hepler
func insertOneMovie(movie model.Netflix){
	inserted , err = collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inerted 1 movie db with id : ", inserted.insertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M("$set": bson.M{"watched":true})

	result, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count : ", result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id, _  = primitive.ObjectIDFromHex(movieId)
	filter = bson.M{"_id": id}
	deleteCount, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete eith delete count : ", deleteCount)
}

func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
    
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of Movie delete : ", deleteResult.DetetedCount)
	return deleteResult.DetetedCount
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies [] primitive.M

	for cur.next(context.Background()){
		var movie bson.M
		err := cur.Decode({&movie})
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
} 

// Controllers

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(allMovies)
}


func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewEncoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}


func MarksAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
    
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie() {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
    
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies() {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
    
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}