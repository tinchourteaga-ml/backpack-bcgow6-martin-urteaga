package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (u *Usuario) cambiarNombre(nombre string, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) cambiarEdad(edad int) {
	u.Edad = edad
}

func (u *Usuario) cambiarCorreo(correo string) {
	u.Correo = correo
}

func (u *Usuario) cambiarContraseña(contraseña string) {
	u.Contraseña = contraseña
}

func main() {
	usuario := Usuario{Nombre: "John", Apellido: "Doe", Edad: 34, Correo: "jdoe@gmail.com", Contraseña: "abc123"}
	fmt.Println(usuario)
	usuario.cambiarNombre("Jane", "Dou")
	usuario.cambiarEdad(27)
	usuario.cambiarCorreo("jdou@yahoo.com")
	usuario.cambiarContraseña("123abc")
	fmt.Println(usuario)
}
