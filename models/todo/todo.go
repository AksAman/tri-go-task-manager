package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/AksAman/tri/utils"
)

type Item struct {
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

func (item Item) PrettyPriority() string {
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

func (item Item) Label() string {
	return strconv.Itoa(item.Position)
}

func ReadItems(filename string) ([]Item, error) {
	if !utils.DoesFileExists(filename) {
		return []Item{}, errors.New(filename + " doesn't exist!")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item

	if err := json.Unmarshal(data, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}

func SaveItems(filename string, items []Item) error {
	marshalled, err := json.MarshalIndent(items, "", "    ")

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, marshalled, 0644)
	if err != nil {
		return err
	}

	return nil
}

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

func ShowTridos(items []Item, filterCondition func(item Item) bool) {
	filteredItems := []Item{}
	for _, item := range items {
		if filterCondition(item) {
			filteredItems = append(filteredItems, item)
		}
	}

	if len(filteredItems) == 0 {
		fmt.Println("No TODOs found!")
	}

	sort.Sort(ItemsByPri(filteredItems))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	defer w.Flush()

	for _, item := range filteredItems {
		fmt.Fprint(w, item.PrettyItem())
	}
}
