/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	server = flag.Bool("server", false, "run the test server")
	slow   = flag.Bool("slow", false, "simulate a slow server")
)

const reply = `{
	"id": 1,
	"category": {
		"id": 0,
		"name": "string"
	},
	"name": "sally",
	"photoUrls": ["string"],
	"tags": [{ "id": 0,"name": "string" }],
	"status":"available"
}`

func main() {
	flag.Parse()

	if *server {
		http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("url: %q\n", html.EscapeString(r.URL.Path))
			defer r.Body.Close()
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(body))
			if *slow {
				time.Sleep(10 * time.Second)
			}
			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write([]byte(reply))
			if err != nil {
				panic(err)
			}
		})

		err := http.ListenAndServe(":1234", nil)
		if err != nil {
			panic(err)
		}
	}
}
