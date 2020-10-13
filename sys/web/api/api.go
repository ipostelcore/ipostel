package api

import (
	"net/http"

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
