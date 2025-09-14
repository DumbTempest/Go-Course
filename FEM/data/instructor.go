package data

type Instructor struct {
	Id    int
	Name  string
	Email string
	Phone string
}

func NewInstructor(Naam string, Gmail string) Instructor {
	return Instructor{
		Name:  Naam,
		Email: Gmail,
	}
}
