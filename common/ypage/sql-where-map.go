package ypage

import (
	"go.mongodb.org/mongo-driver/bson"
)

type SqlWhereMap struct {
}

func (r *SqlWhereMap) GetWhere(a bson.M) (exp string, params []interface{}) {
	d := r.M2D(a)

	//
	return new(SqlWhere).GetWhere(d)
}

func (r *SqlWhereMap) M2D(a bson.M) bson.D {
	if len(a) == 0 {
		return nil
	}
	d := bson.D{}
	for k, v := range a {
		e := bson.E{
			Key:   k,
			Value: v,
		}
		d = append(d, e)
	}
	return d
}
