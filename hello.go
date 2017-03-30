package main

import (
    "fmt"
    "html"
    "log"
  	"github.com/ajstarks/svgo"
    "net/http"
)
var x = 1

func main() {

    http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
        fmt.Print("TEST")
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
        fmt.Fprint(w, x)
        x+=1
    })
    
    http.HandleFunc("/image/", func(w http.ResponseWriter, r *http.Request){
    		fmt.Print(r.URL.Path[1:])
    		http.ServeFile(w, r, r.URL.Path[1:])
    })

    http.HandleFunc("/svg/", func(w http.ResponseWriter, r *http.Request){
      width  :=500
      height :=500
    	canvas := svg.New(w)
    	canvas.Start(width, height)
    	canvas.Circle(width/x, height/x, 100)
    	x+=1
    	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
    	canvas.End()
    })
    log.Fatal(http.ListenAndServe(":8081", nil))

}