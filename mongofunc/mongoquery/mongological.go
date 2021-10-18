package mongoquery

import "go.mongodb.org/mongo-driver/bson"

func MOr(filters ...interface{}) bson.M {
	return bson.M{"$or": filters}
}
