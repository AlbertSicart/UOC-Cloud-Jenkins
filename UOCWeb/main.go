package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// HTML con mensaje e imagen
	html := `
	<html>
	<head>
	<title>Alumno UOC</title>
	</head>
	<body style="font-family: Arial; text-align:center; margin-top: 50px;">
	<h1>Soy Alumno de la UOC</h1>
	<img src="/static/uoc.png" alt="Logo UOC"/>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

func main() {
	//Servir los archivos estáticos (como la imagen)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Ruta principal
	http.HandleFunc("/", handler)

	fmt.Println("Servidor web en ejecución en http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
