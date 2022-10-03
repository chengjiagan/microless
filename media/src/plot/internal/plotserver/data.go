package plotserver

import "go.mongodb.org/mongo-driver/bson/primitive"

type Plot struct {
	PlotOid primitive.ObjectID `json:"plot_id" bson:"_id,omitempty"`
	Plot    string             `json:"plot" bson:"plot"`
}
