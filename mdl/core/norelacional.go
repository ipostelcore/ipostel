package core

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ipostelcore/ipostel/sys"
	"github.com/ipostelcore/ipostel/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

//CrearNOSQL Insert, Generador de Consultas
func (C *Core) CrearNOSQL(coleccion string, query string, db *mongo.Database) (jSon []byte, err error) {
	var M util.Mensajes
	c := db.Collection(coleccion)
	var lst []bson.M
	var bsonMap bson.M
	fmt.Println(query)
	err = json.Unmarshal([]byte(query), &bsonMap)
	fmt.Println("aa ", bsonMap)
	rs, err := c.Find(sys.Contexto, bsonMap)
	if err = rs.All(sys.Contexto, &lst); err != nil {
		M.Msj = "Driver de conexión falló"
		M.Tipo = 1
		M.Fecha = time.Now()
		jSon, err = json.Marshal(M)
		return
	}
	jSon, _ = json.Marshal(lst)
	return
}
