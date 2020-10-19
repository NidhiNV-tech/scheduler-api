package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//Connection mongoDB with helper class
collection := helper.ConnectDB()
func main() {
	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/meeting/{id}", getMeeting).Methods("GET")
	// r.HandleFunc("/meetings", getMeetingForTime).Queries("starttime","{starttime}").Methods("GET")
	r.HandleFunc("/meetings", scheduleMeeting).Methods("POST")
	r.HandleFunc("/meetings", getMeetingForParticipant).Queries("email","{email}").Methods("GET")
	

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getMeetingForTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created meeting array
	var meetings []models.Meeting

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var meeting models.Meeting
		// & character returns the memory address of the following variable.
		err := cur.Decode(&meeting) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		meetings = append(meetings, meeting)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(meetings) // encode similar to serialize process.
}

func getMeeting(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var meeting models.Meeting
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&meeting)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(meeting)
}

func scheduleMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var meeting models.Meeting

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&meeting)

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), meeting)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

//func getMeetingForParticipant, here in filter add Meeti