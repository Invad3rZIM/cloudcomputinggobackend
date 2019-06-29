package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) RestatusItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "activity", "status"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for readability
	activity := requestBody["activity"].(string)
	status := requestBody["status"].(string)

	//modify status
	if item, hasItem := h.AllItems[activity]; hasItem == true {
		if status == "incomplete" {
			item.Status = "incomplete"
			h.FireStore.UncompleteItem(item)
		} else if status == "complete" {
			item.Status = "complete"
			h.FireStore.CompleteItem(item)
		}
	}

	json.NewEncoder(w).Encode(&h.AllItems)
	w.WriteHeader(http.StatusOK)
}
