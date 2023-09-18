package model

type ClassRecord struct {
	ID     uint   `gorm:"primarykey" csv:"id"`
	Time   string `csv:"time"`
	First  int    `csv:"first"`
	Second int    `csv:"second"`
	Third  int    `csv:"third"`
	Work   int    `csv:"work"`
}

func (c *ClassRecord) TableName() string {
	return database + ".class_records"
}
