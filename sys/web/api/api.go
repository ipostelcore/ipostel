package api

import (
	"encoding/json"
	"net/http"

	"github.com/ipostelcore/ipostel/mdl/core"
	"github.com/ipostelcore/ipostel/sys/seguridad"
	"github.com/ipostelcore/ipostel/util"
)

//UsuarioConectado Seguridad Informatica
var UsuarioConectado seguridad.Usuario

//API estructuras generales
type API struct {
	OID    int    `json:"oid"`
	Source string `json:"source"`
	Urls   string `json:"urls"`
	Base   string `json:"base"`
	Query  string `json:"query"`
	Autor  string `json:"autor"`
}

//Agregar una interfaz blanca
func (a *API) Agregar(w http.ResponseWriter, r *http.Request) {

}

//Consultar una interfaz blanca
func (a *API) Consultar(w http.ResponseWriter, r *http.Request) {

}

//Oficinas Listar todas las oficinas
func (a *API) Oficinas(w http.ResponseWriter, r *http.Request) {

	var c core.Core
	Cabecera(w, r)
	w.WriteHeader(http.StatusOK)
	j, _ := c.Oficinas()
	w.Write(j)
	return
}

//Crud conexion para solicitud de token
func (a *API) Crud(w http.ResponseWriter, r *http.Request) {
	var c core.Core
	Cabecera(w, r)
	var v map[string]interface{}
	e := json.NewDecoder(r.Body).Decode(&v)
	util.Error(e)

	j, _ := c.CrearQuery(v)

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
