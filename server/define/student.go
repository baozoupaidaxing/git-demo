package define

type Student struct {
	Id     int `gorm:"primary_key"`
	Name   string
	Class  string
	Course string
	Age    int
	Grade  int
}
