package main

import (
	"log"
	"net/http"
)

/*
  Package needed is
  log		-> reference : https://golang.org/pkg/log/
  net/http	-> reference : https://golang.org/pkg/net/http/

*/
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*
			root page like (index.php/index.html) is set with /
			fungsi ResponseWriter di aliaskan dengan variabel w
			fungsi Request di aliaskan dengan vairabel r

			The function signature of func(w http.ResponseWriter, r *http.Request) is a common way of handling HTTP requests throughout the Go standard library
		*/
		//fungsi w (ResponseWriter) membuat konten html
		w.Write([]byte(`
		<html>
               <head>
			   <title> Simple Web with GO</title>
			   </head>
			   <body>
                      <p>Lets Chat</p>
			   </body>
		</html>
		`))
	})
	//start listen port
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("listenAndServe:", err)
	}
}

/**
output :
127.0.0.1:8080
________________
Lets Chat
**/
