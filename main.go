package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>Sample Go App</title>
				</head>
				<body>
					<p>Sample Go App</p>
				</body>
			</html>
		`))
	})

	log.Println("Golang application starting on http://localhost:8080")
	log.Println("Ctrl-C to shutdown server")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
