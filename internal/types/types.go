package types

// Student represents a student
type Student struct {
	Id    int64
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=0,lte=100,required"`
}
