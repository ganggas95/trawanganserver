package models

import (
	"github.com/revel/revel"
	"time"
)

type Person struct {
	IdPerson     int64     `gorm:"column:person_id;primary_key;AUTO_INCREMENT"`
	Nama         string    `gorm:"column:nama;type:varchar(45)"`
	Alamat       string    `gorm:"column:alamat;type:varchar(255)"`
	TempatLahir  string    `gorm:"column:tempat_lahir;type:varchar(45)"`
	TanggalLahir time.Time `gorm:"column:tanggal_lahir;"`
	Pekerjaan    string    `gorm:"column:pekerjaan;type:varchar(45)"`
}

const (
	SQL_DATE_FORMAT = "2006-01-02"
)

func (person *Person) Validation(v *revel.Validation) {
	v.Check(person.Nama,
		revel.Required{},
		revel.MinSize{4},
		revel.MaxSize{45},
	)
	v.Check(person.Alamat,
		revel.Required{},
		revel.MinSize{8},
		revel.MaxSize{40},
	)
	v.Check(person.TempatLahir,
		revel.Required{},
		revel.MinSize{4},
		revel.MaxSize{20},
	)
	person.TanggalLahir.Format(SQL_DATE_FORMAT)

}
