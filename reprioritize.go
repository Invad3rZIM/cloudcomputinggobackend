package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) ReprioritizeItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var requestBody map[string]interface{}

	//ensure json is decoded
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//ensure all requisite json components are found
	if err := h.VerifyBody(requestBody, "activity", "priority"); err != nil {
		fmt.Fprintln(w, err.Error())

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	//extract field values to variables for readability
	activity := requestBody["activity"].(string)
	newPriority := int(requestBody["priority"].(float64))

	//modify priority
	if item, hasItem := h.AllItems[activity]; hasItem == true {
		item.Priority = newPriority
		h.FireStore.UpdateDB(item)
	}

	json.NewEncoder(w).Encode(&h.AllItems)
	w.WriteHeader(http.StatusOK)
}
