package sys

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ipostelcore/ipostel/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Mongo Database
type Mongo struct {
	DataBase *mongo.Database
	Ctx      context.Context
}

//Conectar Establecer conexión
func (m *Mongo) Conectar(Ctx context.Context, mapa map[string]CadenaDeConexion) (DB *mongo.Database) {
	c := mapa["mongodb"]
	uri := "mongodb://" + c.Host + ":" + c.Puerto
	Client, err := mongo.Connect(Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(Ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	DB = Client.Database(c.Basedatos)
	util.Error(err)
	fmt.Println("[Host: " + c.Host + " Base De Datos: " + c.Basedatos + "  OK...] ")
	return
}

//Buscar un registro
func (m *Mongo) Buscar(coleccion string, filtro interface{}) (j []byte, err error) {
	var doc interface{}
	c := m.DataBase.Collection(coleccion)
	//.Select(bson.M{"clave": false})
	err = c.FindOne(Contexto, bson.M{"cedula": ""}).Decode(&doc)
	j, _ = json.Marshal(doc)
	return
}

//Salvar documentos
func (m *Mongo) Salvar(mgo interface{}, coleccion string) (err error) {

	// c := MGOSession.DB(CBASE).C(coleccion)
	// err = c.Insert(mgo)

	return
}
