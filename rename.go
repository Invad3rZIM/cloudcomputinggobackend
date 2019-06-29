package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) RenameItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "oldname", "newname"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for readability
	oldName := requestBody["oldname"].(string)
	newName := requestBody["newname"].(string)

	//rename acts as a drop, then an add for simplicity sake
	if item, hasItem := h.AllItems[oldName]; hasItem == true {
		h.AllItems[item.Activity] = nil
		delete(h.AllItems, item.Activity)
		h.FireStore.RemoveFromDB(item)

		replacement := NewItem(newName, item.Priority)
		replacement.Status = item.Status

		//adds new items to map and database
		h.AllItems[replacement.Activity] = replacement
		h.FireStore.UpdateDB(replacement)
	}

	json.NewEncoder(w).Encode(&h.AllItems)
	w.WriteHeader(http.StatusOK)
}
