package web

//Copyright Carlos Peña
//Modulo de negociación WEB
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ipostelcore/ipostel/sys/web/api"
)

//Variables de Control
var (
	Enrutador   = mux.NewRouter()
	WsEnrutador = mux.NewRouter()
)

//Cargar los diferentes modulos del sistema
func Cargar() {
	CargarModulosPanel()
	CargarModulosWeb()
	CargarModulosSeguridad()

	WMAdminLTE()
	CargarModulosWebDevel()
}

//CargarModulosWeb Cargador de modulos web
func CargarModulosWeb() {
	var wUsuario api.WUsuario
	var ap api.API

	Enrutador.HandleFunc("/", Principal)
	Enrutador.HandleFunc("/ipostel/api/crud/{id}", wUsuario.ValidarToken(ap.Consultar)).Methods("GET")
	Enrutador.HandleFunc("/ipostel/api/oficina/listar", wUsuario.ValidarToken(ap.Oficinas)).Methods("GET")

	Enrutador.HandleFunc("/ipostel/api/crud", wUsuario.ValidarToken(ap.Crud)).Methods("POST")
	Enrutador.HandleFunc("/devel/api/crud", ap.Crud).Methods("POST")
}

//CargarModulosPanel Panel de Contencion
func CargarModulosPanel() {

}

//CargarModulosSeguridad Y cifrado
func CargarModulosSeguridad() {
	var wUsuario api.WUsuario
	Enrutador.HandleFunc("/ipostel/app/api/wusuario/login", wUsuario.Login).Methods("POST")
	Enrutador.HandleFunc("/ipostel/app/api/wusuario/login", wUsuario.Opciones).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipostel/api/wusuario/validar", wUsuario.ValidarToken(wUsuario.Autorizado)).Methods("POST")
	Enrutador.HandleFunc("/ipostel/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")

	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.Crear).Methods("POST")
	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")

	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.Crear)).Methods("POST")
	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.CambiarClave)).Methods("PUT")
	Enrutador.HandleFunc("/ipostel/api/wusuario", wUsuario.ValidarToken(wUsuario.Opciones)).Methods("OPTIONS")
	Enrutador.HandleFunc("/ipostel/api/wusuario/listar", wUsuario.ValidarToken(wUsuario.Listar)).Methods("GET")

}

//Principal Página inicial del sistema o bienvenida
func Principal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenidos al Bus Empresarial de Datos")
}

//WMAdminLTE OpenSource tema de panel de control Tecnología Bootstrap3
func WMAdminLTE() {
	fmt.Println("Módulos de AdminLTE V 3.5")
	prefix := http.StripPrefix("/ipostel", http.FileServer(http.Dir("public_web/www")))
	Enrutador.PathPrefix("/ipostel/").Handler(prefix)
}

//CargarModulosWebDevel Cargador de modulos web
func CargarModulosWebDevel() {

}
