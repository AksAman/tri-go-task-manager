package models

import (
	"fmt"
	"strconv"
)

type Item struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Priority int    `json:"priority"`
	Position int    `json:"position"`
	Done     bool   `json:"done"`
}

func (item *Item) PrettyItem() string {
	return fmt.Sprintf("\t%v. \t%v \t %v \t %v\n", item.Label(), item.Status(), item.Text, item.PrettyPriority())
}

func (item *Item) SetPriority(priority int) {
	switch priority {
	case 1, 3:
		item.Priority = priority
	default:
		item.Priority = 2
	}
}

func (item *Item) PrettyPriority() string {
	switch item.Priority {
	case 1:
		return "HIGH"
	case 3:
		return "LOW"
	default:
		return " "
	}
}

func (item *Item) Status() string {
	if item.Done {
		return "[x]"
	}
	return "[ ]"
}

func (item *Item) Label() string {
	return strconv.Itoa(item.ID)
}

// region ItemsByPri sort.Interface
// ItemsByPri implements sort.Interface for Item
type ItemsByPri []Item

func (s ItemsByPri) Len() int {
	return len(s)
}

func (s ItemsByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ItemsByPri) Less(i, j int) bool {

	if s[i].Done != s[j].Done {
		return s[j].Done
	}

	if s[i].Priority == s[j].Priority {
		return s[i].Position < s[j].Position
	}
	return s[i].Priority < s[j].Priority
}

// endregion DB
