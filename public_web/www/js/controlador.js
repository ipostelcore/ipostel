/**
 * 
 */
let _ANOACTUAL = new Date().getUTCFullYear();

let FrmValidar = false;
let JsonMenu = {};
let Util = new Utilidad();


Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000
});


 $(function(){
  if (sessionStorage.getItem('ipostel') == undefined ){
    $(location).attr("href","index.html");
  }else{
    console.info('Iniciando carga del sistema');
    CargarUrl('modales', 'inc/modales');
    //CargarUrl('cuerpo', 'inc/escritorio');
    Util.CargarJsonMenu(); 
    Util.DibujarEscritorio(JsonMenu);

    Util.ProcesoExitoso('&nbsp;La página se ha cargado exitosamente');

  }
    
  

 }
 ); //Fin de la funcion


/**
 * Carga inicial de Configuracion
 * V 1.0
 */
 function CargarConfiguracion(){
    CargarUrl('cuerpo', 'inc/configuracion/submenu');
    //Util.DibujarSubMenuEscritorio(JsonMenu, "Configuracion");
 }

 function CargarUsuarios(){
   CargarUrl('cuerpo', 'inc/configuracion/usuarios');
   CargarUrl('cargamasiva', 'inc/configuracion/usuariosmodal');
 }
 function CargarServicios(){
   //CargarUrl('cuerpo', 'inc/configuracion/servicios');
   CargarUrl('_mdlCuerpo', 'inc/configuracion/serviciosmodal');

 }
 function CargarRoles(){
   CargarUrl('cuerpo', 'inc/configuracion/servicios');
   CargarUrl('cargamasiva', 'inc/configuracion/serviciosmodal');
 }
 function CargarRoles(){
   CargarUrl('cuerpo', 'inc/configuracion/roles');
 }
 function CargarOficinas(){
   CargarUrl('cuerpo', 'inc/configuracion/oficinas');
 }
 /**
  * Fin de Configuración
  */


/**
 * Carga inicial de Configuracion
 * V 1.0
 */
function CargarControl(){
  Util.DibujarSubMenuEscritorio(JsonMenu, "Control");
}





function CerrarSession(){
	sessionStorage.removeItem('ipostel');
	$(location).attr("href","../index.html");
}

function IniciarSesion(){
  if (sessionStorage.getItem('ipostel') != undefined ){

    var e = sessionStorage.getItem("ipostel");
    var s = e.split(".");
    var json = JSON.parse(atob(s[1]));
    Usuario = json.Usuario;


    $("#_PerfilUsuario").html(Usuario.Perfil.descripcion);
    $("#_NombreUsuario").html(Usuario.nombre);

  }
}

 function CargarReportes(){
   CargarUrl('cuerpo', 'inc/configuracion/submenu_reportes');
  // CargarUrl('cargamasiva', 'inc/configuracion/usuariosmodal');
 }

 function CargarVentaPorSermana(){
   CargarUrl('cuerpo', 'inc/configuracion/ventaporsemana');
  // CargarUrl('cargamasiva', 'inc/configuracion/usuariosmodal');
 }


