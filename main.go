package main

import (
	"encoding/json"
	//"encoding/json"
	"fmt"
	"github.com/josuepr2/aws-tutorial/app/core"
	"log"
	"net/http"
)

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world :)")
}

func main() {
	fmt.Println("Running...")

	s3ClientConn, errS3 := core.NewS3Connection()

	if errS3 != nil{
		fmt.Println("error")
	}

	http.HandleFunc("/", HomeEndpoint)
	http.HandleFunc("/s3-list", func(w http.ResponseWriter, r *http.Request){

		res, listErr := s3ClientConn.ListS3Buckets()

		if listErr != nil {
			w.Write([]byte("Error listing buckets"))
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			json.NewEncoder(w).Encode(&res)
		}
	})


	if err := http.ListenAndServe(":3666", nil); err != nil {
		log.Fatal(err)
	}
}
