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
    Util.CargarJsonMenu();
    Util.DibujarEscritorio(JsonMenu);
    Util.DibujarMenuLateral(JsonMenu);
    Util.DibujarEscritorio(JsonMenu);
    ObtenerOficinas();

  }



 }


 ); //Fin de la funcion


/**
 * Carga de Módulo CONFIGURACIÓN
 * V 1.0
 */
 function CargarConfiguracion(){
    Util.DibujarSubMenuEscritorio(JsonMenu, "Configuracion");
 }

 function CargarUsuarios(){
  //Util.DibujarSubMenuEscritorio(JsonMenu, "Usuarios");
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
 * Carga Módulo CONTROL
 * V 1.0
 */
function CargarControl(){
  Util.DibujarSubMenuEscritorio(JsonMenu, "Control");
}

function CargarSemanas(){
  CargarUrl('cuerpo', 'inc/control/semanas');
}




/**
 * Función de Cierre de Sesión
 */
function CerrarSesion(){
	sessionStorage.removeItem('ipostel');
	$(location).attr("href","../index.html");
}


/**
 * Funcion de Iniciar Sesión
 */
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

 /**
  * Carga de Módulo TRACKING
  */

 function CargarTraking(){
    Util.DibujarSubMenuEscritorio(JsonMenu, "Tracking");
 }

 function CargarConsignacion(){
   CargarUrl('cuerpo', 'inc/ventas/consignacion');
   
 }


 function CargarSeguimiento(){
   CargarUrl('cuerpo', 'inc/tracking/seguimiento');
 }

 function CargarDistribucion(){
   CargarUrl('cuerpo', 'inc/tracking/manejodistribucion');
 }

 function CargarConfiguracionTracking(){
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


/**
 * Carga de Módulo VENTAS
 * V 1.0
 */
function CargarVentas(){
  Util.DibujarSubMenuEscritorio(JsonMenu, "Ventas");
}

 function CargarServicios(){
    CargarUrl('cuerpo', 'inc/ventas/submenu_servicios');
    //Util.DibujarSubMenuEscritorio(JsonMenu, "Servicios");
 }

 function CargarServiciosPostales(){
  CargarUrl('cuerpo', 'inc/ventas/postales');
  //Util.DibujarSubMenuEscritorio(JsonMenu, "Postales");
}

 function CargarServiciosTelegraficos(){
     CargarUrl('cuerpo', 'inc/ventas/telegramas');
 }

 function CargarServiciosVarios(){
  CargarUrl('cuerpo', 'inc/ventas/varios');
 //Util.DibujarSubMenuEscritorio(JsonMenu, "Varios");
}

function CargarServiciosConsignacion(){
  CargarUrl('cuerpo', 'inc/ventas/consignacion');
 //Util.DibujarSubMenuEscritorio(JsonMenu, "EBS/EMS");
}

function CargarServiciosPortePagado(){
  CargarUrl('cuerpo', 'inc/ventas/portepagado');
 //Util.DibujarSubMenuEscritorio(JsonMenu, "Apartados");
}

function CargarServiciosOtros(){
  CargarUrl('cuerpo', 'inc/ventas/serviciosotros');
 //Util.DibujarSubMenuEscritorio(JsonMenu, "Apartados");
}


function CargarServiciosApartados(){
  CargarUrl('cuerpo', 'inc/ventas/apartados');
 //Util.DibujarSubMenuEscritorio(JsonMenu, "Apartados");
}

function CargarServiciosReportes(){
  CargarUrl('cuerpo', 'inc/ventas/reportes');
 //Util.DibujarSubMenuEscritorio(JsonMenu, "Reportes");
}

/**
 * Carga de Módulo Panel
 */

function CargarPanel(){
  Util.DibujarSubMenuEscritorio(JsonMenu, "Panel");
}
function CargarApiCore(){
  CargarUrl('cuerpo', 'inc/panel/apicore');
}