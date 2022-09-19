package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IncidentsOn struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IncidentID string             `json:"incidentID,omitempty" bson:"incidentID,omitempty"`
	ClientID   string             `json:"clientID,omitempty" bson:"catchphrase,omitempty"`
	Address    string             `json:"address,omitempty" bson:"address,omitempty"`
	Status     string             `json:"status,omitempty" bson:"status,omitempty"`
}
