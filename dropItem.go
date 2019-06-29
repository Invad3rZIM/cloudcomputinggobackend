package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) DropItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "activity"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for readability
	activity := requestBody["activity"].(string)

	//removes item from allitems if exists and deletes from database
	if item, hasItem := h.AllItems[activity]; hasItem == true {
		h.AllItems[item.Activity] = nil
		delete(h.AllItems, item.Activity)
		h.FireStore.RemoveFromDB(item)
	}

	json.NewEncoder(w).Encode(&h.AllItems)
	w.WriteHeader(http.StatusOK)
}
