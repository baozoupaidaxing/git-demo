package psql

import (
	"github.com/jinzhu/gorm"
	"goProject/student/server/define"
)

func SelectById(sId int) (*define.Student, error) {
	student := &define.Student{}
	err := studentsDB.Table("students").Where("id=?", sId).Find(&student).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return student, nil
		}
		return nil, err
	}
	//fmt.Println(student)
	return student, nil
}
