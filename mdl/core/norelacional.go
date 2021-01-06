package core

import (
	"encoding/json"
	"fmt"

	"github.com/ipostelcore/ipostel/sys"
	"github.com/ipostelcore/ipostel/util"
	"go.mongodb.org/mongo-driver/bson"
)

//InsertNOSQL Insert, Generador de Consultas
func (C *Core) InsertNOSQL(v map[string]interface{}) (jSon []byte, err error) {

	c := sys.MongoDB.Collection(sys.APICORE)
	rs, err := c.InsertOne(sys.Contexto, v)
	if err != nil {
		fmt.Println("Error creando Query en Mongodb ", err.Error())
	}
	jSon, _ = json.Marshal(rs)
	return

}

//Listar Insert, Generador de Consultas
func (C *Core) Listar() (jSon []byte, err error) {
	c := sys.MongoDB.Collection(sys.APICORE)
	var lst []ApiCore
	rs, err := c.Find(sys.Contexto, bson.M{})
	for rs.Next(sys.Contexto) {
		var api ApiCore
		e := rs.Decode(&api)
		util.Error(e)
		lst = append(lst, api)
	}
	jSon, _ = json.Marshal(lst)
	return

}
