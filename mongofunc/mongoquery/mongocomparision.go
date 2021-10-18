package mongoquery

import "go.mongodb.org/mongo-driver/bson"

func MIn(fieldName string, value ...interface{}) bson.M {
	return bson.M{fieldName: bson.M{"$in": value}}
}

func MNotIn(fieldName string, value ...interface{}) bson.M {
	return bson.M{fieldName: bson.M{"$nin": value}}
}

// LESS OR EQUAL
func MEqualLess(fieldName string, value interface{}) bson.M {
	return bson.M{fieldName: bson.M{"$lte": value}}
}

func MLess(fieldName string, value interface{}) bson.M {
	return bson.M{fieldName: bson.M{"$lt": value}}
}

// GREATER OR EQUAL
func MEqualGreaterInt(fieldName string, value int) bson.M {
	return bson.M{fieldName: bson.M{"$gte": value}}
}

func MEqualGreaterInt64(fieldName string, value int64) bson.M {
	return bson.M{fieldName: bson.M{"$gte": value}}
}

func MGreaterInt(fieldName string, value int) bson.M {
	return bson.M{fieldName: bson.M{"$gt": value}}
}

func MGreaterInt64(fieldName string, value int64) bson.M {
	return bson.M{fieldName: bson.M{"$gt": value}}
}

func MNotEqual(fieldName string, value interface{}) bson.M {
	return bson.M{fieldName: bson.M{"$ne": value}}
}

func MEqual(fieldName string, value interface{}) bson.M {
	return bson.M{fieldName: value}
}
