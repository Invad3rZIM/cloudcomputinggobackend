package main

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FireStore struct {
	client *firestore.Client
}

func (fb *FireStore) LoadAllItemsFromDB() map[string]*Item {
	allItems := make(map[string]*Item)

	for k, v := range fb.LoadCollectionFromDB("incomplete") {
		allItems[k] = v
	}

	for k, v := range fb.LoadCollectionFromDB("complete") {
		allItems[k] = v
	}

	return allItems
}

//Loads Collection into a map for Server Caching
//RequisiteStatus = incomplete or complete
func (fb *FireStore) LoadCollectionFromDB(requisiteStatus string) map[string]*Item {
	items := make(map[string]*Item)

	q := fb.client.Collection(requisiteStatus)
	iter := q.Documents(context.Background())

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		bs, err := json.Marshal(doc.Data())
		var item *Item

		json.Unmarshal(bs, &item)

		items[item.Activity] = item
	}

	return items
}

func (fb *FireStore) UpdateDB(item *Item) {
	var err error

	//Deletes the item from the Collection it isn't in..
	if item.Status == "incomplete" {
		_, err = fb.client.Collection("complete").Doc(item.Activity).Delete(context.Background())
	} else if item.Status == "complete" {
		_, err = fb.client.Collection("incomplete").Doc(item.Activity).Delete(context.Background())
	}

	if err != nil {
		log.Fatal(err)
	}

	_, err = fb.client.Collection(item.Status).Doc(item.Activity).Set(context.Background(), item)
}

//RemoveFromDB deletes an Item from FireBase if it exists
func (fb *FireStore) RemoveFromDB(item *Item) {
	_, err := fb.client.Collection(item.Status).Doc(item.Activity).Delete(context.Background())

	if err != nil {
		log.Fatal(err)
	}
}
