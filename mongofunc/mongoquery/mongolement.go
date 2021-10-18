package mongoquery

import "go.mongodb.org/mongo-driver/bson"

func MExistOrDefault(fieldName string, exist bool, value interface{}) bson.M {
	return MOr(MEqual(fieldName, value), MExist(fieldName, exist))
}

func MExist(fieldName string, exist bool) bson.M {
	return bson.M{fieldName: bson.M{"$exist": exist}}
}
