package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ipostelcore/ipostel/mdl/core"
	"github.com/ipostelcore/ipostel/sys"
	"github.com/ipostelcore/ipostel/sys/seguridad"
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
	if e == nil {
		j, _ := c.OperarConsulta(v)

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("error de objeto")
		fmt.Fprintln(w, "Error en el objeto de carga")
	}

}

//Listar conexion para solicitud de token
func (a *API) Listar(w http.ResponseWriter, r *http.Request) {
	var xcore core.Core
	Cabecera(w, r)
	j, _ := xcore.Listar(sys.MGOSession)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
