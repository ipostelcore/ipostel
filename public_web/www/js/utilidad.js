
class Utilidad{
    constructor(){

    }
    /**
     * 
     * @param {*} mensaje string contenido 
     * @param {*} clase  string bg-danger, bg-info, bg-success, bg-warning
     * @param {*} titulo string
     * @param {*} subtitulo string
     */
    Mensaje(mensaje, clase, titulo, subtitulo){
          $(document).Toasts('create', {
            class: clase, 
            title: titulo,
            autohide: true,
            delay: 2500,
            subtitle: subtitulo,
            body: mensaje,
            icon: 'warning',
            position: 'bottomRight'
          })
       
    }

    ProcesoExitoso(mensaje){
        Toast.fire({
            icon: 'success',
            title: mensaje
        })
    }
  

}