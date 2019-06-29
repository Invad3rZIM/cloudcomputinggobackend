package main

type Item struct {
	Activity string
	Status   string //status can be complete / incomplete
	Priority int    //to allow priority sorting
}

//NewItem is an itemConstructor
func NewItem(activity string, priority int) *Item {
	return &Item{Activity: activity,
		Status:   "incomplete",
		Priority: priority}
}

func (fb *FireStore) CompleteItem(item *Item) {
	item.Status = "complete"
	fb.UpdateDB(item)
}

func (fb *FireStore) UncompleteItem(item *Item) {
	item.Status = "incomplete"
	fb.UpdateDB(item)
}

func (fb *FireStore) ChangePriority(item *Item, newPriority int) {
	item.Priority = newPriority
	fb.UpdateDB(item)
}
