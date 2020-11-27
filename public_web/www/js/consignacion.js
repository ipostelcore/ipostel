/**
 * Clase: Consignación en ventas
 * Referencia: inc/ventas/consignacion
 * Autor: Carlos Peña
 * Fecha: 27-11-2020
 * @gesaodin
 */
class Consignacion{
    constructor(){
        this.tipoenvio = '';
        this.peso = 0;
        this.tipo = '';
        this.tarifa = '';
        this.grupo = '';
        this.exonerar = '';
        this.serial = '';
        this.serialsunacrip = '';
        this.monto = 0;
        this.totaliva = 0;
        this.cobrado = 0;
        this.oficinaorigen = '';
        this.remitente = false;
        this.cedula = '';
        this.nombreapellido = '';
        this.direccion = '';
        this.telefono = '';
        this.correo = '';
        this.nombredestinatario = '';
        this.paisdestino = '';
        this.oficinadestino = '';
        this.direcciondestino = '';
        this.telefonodestino = '';
    }
    
    Obtener(){
        return this;
    }

    Crear(req){

    }
}

