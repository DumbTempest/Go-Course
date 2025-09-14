package data

import "fmt"

func (i Instructor) Print() string {
	return fmt.Sprintf("Id: %d, Name: %s, Email: %s, Phone: %s", i.Id, i.Name, i.Email, i.Phone)
}
