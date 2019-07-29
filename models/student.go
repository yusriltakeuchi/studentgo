package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id      int       `orm:"auto"`
	Name    string    `orm:"size(100)"`
	Email   string    `orm:"size(100);unique"`
	Address string    `orm:"size(200)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Upadted time.Time `orm:"auto_now;type(datetime)"`
}

//Initialize model
func init() {
	orm.RegisterModel(new(Student))
}

//GetAllStudents Get all student
func GetAllStudents() []*Student {
	o := orm.NewOrm()
	var students []*Student

	qs := o.QueryTable(new(Student))
	qs.OrderBy("name").All(&students)

	return students
}

//GetStudentByID Find students by id
func GetStudentByID(id int) *Student {
	o := orm.NewOrm()
	var student Student

	err := o.QueryTable(new(Student)).Filter("id", id).One(&student)
	if err == orm.ErrNoRows {
		return nil
	}
	return &student
}

//GetStudentByEmail Find students by email
func GetStudentByEmail(email string) *Student {
	o := orm.NewOrm()
	var student Student

	err := o.QueryTable(new(Student)).Filter("email", email).One(&student)
	if err == orm.ErrNoRows {
		return nil
	}
	return &student
}

//InsertStudent insert new student
func InsertStudent(student Student) error {
	o := orm.NewOrm()
	_, err := o.Insert(&student)
	return err
}

//UpdateStudent update student data
func UpdateStudent(student Student, id int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Student)).Filter("id", id).Update(orm.Params{
		"name":    student.Name,
		"email":   student.Email,
		"address": student.Address,
	})

	return err
}

//DeleteStudent delete specific student
func DeleteStudent(id int) error {
	o := orm.NewOrm()
	student := Student{Id: id}
	_, err := o.Delete(&student)

	return err
}
