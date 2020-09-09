package user

import "time"

// creamos estructura
type Usuario struct {
	Id			int
	Nombre		string
	FechaAlta	time.Time
	Status 		bool
}

// hace una referencia a Usuario, utilizo el puntero
func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, status bool){
	this.Id			= id
	this.Nombre 	= nombre
	this.FechaAlta 	= fechaalta
	this.Status 	= status
}