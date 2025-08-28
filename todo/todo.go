package todo

import (
	"encoding/json"
	"os"
	"strconv"
	"time"
)

type Item struct {
	Text      string
	Priority  int
	Position  int
	Done      bool
	CreatedAt time.Time
}
type ByPri []Item

func (s ByPri) Len() int {
	return len(s)
}
func (s ByPri) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	if s[i].Priority != s[j].Priority {
		return s[i].Priority < s[j].Priority
	}
	return s[i].Position < s[j].Position
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2

	}
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}
	return " "
}
func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	} else {
		return ""
	}
}

func (i *Item) PrettyCreated() string {
	return i.CreatedAt.Format("2006-01-02 15:04")
}
func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}
func SaveItems(filename string, Items []Item) error {
	b, err := json.Marshal(Items)
	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	for i, _ := range items {
		items[i].Position = i + 1
	}
	return items, nil

}
