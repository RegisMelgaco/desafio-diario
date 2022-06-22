package entity

type Email struct {
	val string
}

func NewEmail(val string) Email {
	return Email{val}
}
