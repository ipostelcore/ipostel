/**
 *
 */

let Util = new Utilidad();


Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000
});


 $(function(){

    console.info('Iniciando carga del sistema');
    CargarUrl('modales', 'inc/modales');
    CargarUrl('cuerpo', 'inc/escritorio');
    Util.ProcesoExitoso('&nbsp;La página se ha cargado exitosamente');

 }


 ); //Fin de la funcion


/**
 * Carga inicial de Configuracion
 * V 1.0
 */
 function CargarConfiguracion(){
    CargarUrl('cuerpo', 'inc/configuracion/submenu');
 }
 function CargarUsuarios(){
   CargarUrl('cuerpo', 'inc/configuracion/usuarios');
   CargarUrl('cargamasiva', 'inc/configuracion/usuariosmodal');
 }
 function CargarServicios(){
   CargarUrl('cuerpo', 'inc/configuracion/servicios');
   CargarUrl('cargamasiva', 'inc/configuracion/serviciosmodal');
 }
 function CargarRoles(){
   CargarUrl('cuerpo', 'inc/configuracion/roles');
 }
 function CargarOficinas(){
   CargarUrl('cuerpo', 'inc/configuracion/oficinas');
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
  * Fin de Configuración
  */
