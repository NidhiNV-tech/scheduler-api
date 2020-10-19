package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Meeting struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StartTime    string             `json:"starttime,omitempty" bson:"starttime,omitempty"`
	EndTime      string             `json:"endtime" bson:"endtime,omitempty"`
	Timestamp    string             `json:"timestamp" bson:"timestamp,omitempty"`
	Participants *Participants      `json:"participants" bson:"participants,omitempty"`
}

type Participant struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Rsvp  string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}
