package main

import (
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/appengine"
)

func main() {
	// Use the application default credentials
	ctx := appengine.BackgroundContext()
	conf := &firebase.Config{ProjectID: "cloudclasstodo"}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	//initialize handler with firestore pointer
	h := NewHandler(&FireStore{client})

	http.HandleFunc("/item/add", h.AddItem)
	http.HandleFunc("/item/drop", h.DropItem)
	http.HandleFunc("/item/rename", h.RenameItem)
	http.HandleFunc("/", h.Ping)

	http.HandleFunc("/item/reprioritize", h.ReprioritizeItem)

	http.HandleFunc("/item/restatus", h.RestatusItem)

	appengine.Main() // Starts the server to receive requests

	defer client.Close()
}
