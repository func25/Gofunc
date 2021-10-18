package mongoseed

import (
	"context"
	"gofunc/mathfunc"
	"gofunc/mongofunc/mongorely"
)

type Hero struct {
	mongorely.ObjectId `bson:",inline"`
	Name               string `bson:"name"`
	Damage             int    `bson:"damage"`
}

func (*Hero) GetMongoCollName() string {
	return "Heroes"
}

func Seed(ctx context.Context, num int) error {
	for i := 0; i < num; i++ {
		damage, _ := mathfunc.RandomInt(0, 2)
		err := mongorely.Create(ctx, &Hero{
			Name:   "TestHero",
			Damage: damage,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
