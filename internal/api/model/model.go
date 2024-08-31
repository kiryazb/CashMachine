package model


type Item struct {
	Id 		int		`gorm:"primaryKey" json:"-"`			
	Title 	string	`gorm:"size:255;not null;unique"`
	Price 	int		`gorm:"not null"`
}

func (Item) TableName() string {
	return "item"
}


