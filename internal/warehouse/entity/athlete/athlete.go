package athlete

// Реализация 1 ко многим со стороны атлета (1)
type Athlete struct {
	Id      int64
	Name    string
	Surname string
	Tag     string
	AdminId *int32
}
