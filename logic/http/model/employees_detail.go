package model

import (
	"github.com/jinzhu/gorm"
	"echo-framework/lib/db"
)


// EmployeesDetail struct is a row record of the employees_detail table in the jz_ybs database
type EmployeesDetail struct {

	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] employees_id                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EmployeesID uint `gorm:"column:employees_id;type:int;default:0;index" json:"employees_id"` // 员工id

}

func (m EmployeesDetail) Model() *gorm.DB {
	return db.Mysql((&m).Connection()).Model(&m).Table(m.TableName())
}

func (m EmployeesDetail) ModelMaster() *gorm.DB {
	return db.ModelMaster((&m).Connection()).Model(&m).Table(m.TableName())
}

func (*EmployeesDetail) Connection() string {
	return "jz_ybs"
}

// TableName sets the insert table name for this struct type
func (e *EmployeesDetail) TableName() string {
	return "employees_detail"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *EmployeesDetail) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *EmployeesDetail) Prepare() {
}

