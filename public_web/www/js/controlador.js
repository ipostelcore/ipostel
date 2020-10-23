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
    CargarUrl('cuerpo', 'inc/escritorio');
    Util.ProcesoExitoso('&nbsp;La página se ha cargado exitosamente');

 }
 
 
 )



 function CargarConfiguracion(){
    CargarUrl('cuerpo', 'inc/configuracion/usuarios');
 }