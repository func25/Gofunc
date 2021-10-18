package mongoquery

import "go.mongodb.org/mongo-driver/bson"

const (
	INC  = "$inc"
	SET  = "$set"
	PUSH = "$push"
)

// INC
func DIncInt(pairs ...PairSetterInt) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}
	return bson.D{{Key: INC, Value: updated}}
}

func DIncInt64(pairs ...PairSetterInt64) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}
	return bson.D{{Key: INC, Value: updated}}
}

// SET
func DSet(pairs ...PairSetter) bson.D {
	return bson.D{{Key: SET, Value: dPair(pairs)}}
}

// PUSH
func DPush(pairs ...PairSetter) bson.D {
	return bson.D{{Key: PUSH, Value: dPair(pairs)}}
}

func dPair(pairs []PairSetter) bson.D {
	pairLen := len(pairs)
	updated := bson.D{}
	for i := 0; i < pairLen; i++ {
		updated = append(updated, bson.E{Key: pairs[i].FieldName, Value: pairs[i].Value})
	}

	return updated
}
