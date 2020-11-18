/**
 *
 */
let _ANOACTUAL = new Date().getUTCFullYear();

let FrmValidar = false;
let JsonMenu = {};
let Util = new Utilidad();
let conn = new Conexion();


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
    Util.DibujarMenuLateral(JsonMenu);
    Util.DibujarEscritorio(JsonMenu);

    //Util.ProcesoExitoso('&nbsp;La página se ha cargado exitosamente');

  }



 }


 ); //Fin de la funcion


/**
 * Carga inicial de Configuracion
 * V 1.0
 */
 function CargarConfiguracion(){
    //CargarUrl('cuerpo', 'inc/configuracion/submenu');
    Util.DibujarSubMenuEscritorio(JsonMenu, "Configuracion");
 }

 function CargarUsuarios(){
   CargarUrl('cuerpo', 'inc/configuracion/usuarios');
   CargarUrl('cargamasiva', 'inc/configuracion/usuariosmodal');
 }

 function CargarServicios(){
   CargarUrl('cuerpo', 'inc/configuracion/servicios');
   CargarUrl('_mdlCuerpo', 'inc/configuracion/serviciosmodal');
 }

 function CargarRoles(){
   CargarUrl('cuerpo', 'inc/configuracion/roles');
   CargarUrl('_mdlCuerpo', 'inc/configuracion/rolesmodal');
   //CargarUrl('cargamasiva', 'inc/configuracion/serviciosmodal');
 }

 function CargarOficinas(){
   CargarUrl('cuerpo', 'inc/configuracion/oficinas');
   CargarUrl('_mdlCuerpo', 'inc/configuracion/rolesmodal');
 }

 function CargarReportes(){
   CargarUrl('cuerpo', 'inc/configuracion/submenu_reportes');
  // CargarUrl('cargamasiva', 'inc/configuracion/usuariosmodal');
 }

 function CargarVentaporsermana(){
   CargarUrl('cuerpo', 'inc/configuracion/ventaporsemana');
  // CargarUrl('cargamasiva', 'inc/configuracion/usuariosmodal');
 }





/**
 * Carga inicial de Control
 * V 1.0
 */
function CargarControl(){
  Util.DibujarSubMenuEscritorio(JsonMenu, "Control");
}



/**
 * Carga inicial de Ventas
 * V 1.0
 */
function CargarVentas(){
  //CargarUrl('cuerpo', 'inc/configuracion/submenu');
  Util.DibujarSubMenuEscritorio(JsonMenu, "Ventas");
}



function CerrarSesion(){
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
   console.log("Acceder");
   ListarOficinas();
   CargarUrl('cuerpo', 'inc/configuracion/ventaporsemana');
  // CargarUrl('cargamasiva', 'inc/configuracion/usuariosmodal');
 }


 function ListarVentasPorSemanas(){
   alert("mundo")
 }

 // TRACKING

 function CargarTraking(){
    CargarUrl('cuerpo', 'inc/tracking/submenu');
    //Util.DibujarSubMenuEscritorio(JsonMenu, "Configuracion");
 }

 function CargarConsignacion(){
   CargarUrl('cuerpo', 'inc/tracking/consignacion');
   CargarUrl('cargamasiva', 'inc/tracking/consignacionmodal');
 }

 function CargarConsignacionLote(){
   CargarUrl('cuerpo', 'inc/tracking/consignacion');
   CargarUrl('cargamasiva', 'inc/tracking/consignacionmodal');
 }

 function CargarSeguimiento(){
   CargarUrl('cuerpo', 'inc/tracking/seguimiento');
 }

 function CargarDistribucion(){
   CargarUrl('cuerpo', 'inc/tracking/manejodistribucion');
 }

 function CargarConfiguracion(){
   CargarUrl('cuerpo', 'inc/tracking/submenu_distribucion');
 }

 function CargarDefRuta(){
   CargarUrl('cuerpo', 'inc/tracking/def_ruta');
   CargarUrl('cargamasiva', 'inc/tracking/def_rutamodal');
 }

 function CargarApartados(){
   CargarUrl('cuerpo', 'inc/tracking/def_apartados');
   CargarUrl('cargamasiva', 'inc/tracking/def_apartadosmodal');
 }

 function CargarRepartidores(){
   CargarUrl('cuerpo', 'inc/tracking/def_repartidores');
   CargarUrl('cargamasiva', 'inc/tracking/def_repartidoresmodal');
 }

 function CargarIncentivos(){
   CargarUrl('cuerpo', 'inc/tracking/def_incentivos');
   CargarUrl('cargamasiva', 'inc/tracking/def_incentivosmodal');
 }


 // PUNTO POSTAL VENTAS

 function CargarVentas(){
    CargarUrl('cuerpo', 'inc/ventas/submenu');
    //Util.DibujarSubMenuEscritorio(JsonMenu, "Configuracion");
 }

 function CargarServicios(){
    CargarUrl('cuerpo', 'inc/ventas/submenu_servpostales');
    //Util.DibujarSubMenuEscritorio(JsonMenu, "Configuracion");
 }

 function CargarServiciosTelegraficos(){
     CargarUrl('cuerpo', 'inc/ventas/telegramas');
    //Util.DibujarSubMenuEscritorio(JsonMenu, "Configuracion");
 }
