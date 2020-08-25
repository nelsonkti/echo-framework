package model

import (
	"github.com/jinzhu/gorm"
	"echo-framework/lib/db"
)


// EmployeesBase struct is a row record of the employees_base table in the jz_ybs database
type EmployeesBase struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 5] name                                           varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	Name string `gorm:"column:name;type:varchar;size:64;" json:"name"` // 姓名


	EmployeesDetail EmployeesDetail `gorm:"foreignkey:EmployeesID;;association_foreignkey:ID"`

}

func (m EmployeesBase) Model() *gorm.DB {
	return db.Mysql((&m).Connection()).Model(&m).Table(m.TableName())
}

func (m EmployeesBase) ModelMaster() *gorm.DB {
	return db.ModelMaster((&m).Connection()).Model(&m).Table(m.TableName())
}

func (*EmployeesBase) Connection() string {
	return "jz_ybs"
}

func (*EmployeesBase) TableName() string {
	return "employees_base"
}
