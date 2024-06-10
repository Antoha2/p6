package helper

type Test struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
