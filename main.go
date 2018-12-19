package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/josuepr2/aws-tutorial/app/core"
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

		s3ClientConn.ListS3Buckets()

		fmt.Fprintln(w, "Listing ")
	})


	if err := http.ListenAndServe(":3666", nil); err != nil {
		log.Fatal(err)
	}
}
