package controllers

import (
	"encoding/json"
	"strconv"
	"studentgo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

//StudentController struct controller student
type StudentController struct {
	beego.Controller
}

//GetStudents Get all studentts data
func (stucon *StudentController) GetStudents() {
	students := models.GetAllStudents()

	response := map[string]interface{}{}
	response["status"] = 200
	response["message"] = "Successfully get students"
	response["data"] = students

	stucon.Data["json"] = response
	stucon.ServeJSON()
	return
}

//GetStudent Get student by specific id
func (stucon *StudentController) GetStudent() {
	//Convert string to int
	id, _ := strconv.Atoi(stucon.Ctx.Input.Param(":id"))

	student := models.GetStudentByID(id)
	response := map[string]interface{}{}

	//If student not found
	if student == nil {
		response["status"] = 404
		response["message"] = "Student not found"
		stucon.Data["json"] = response
		stucon.ServeJSON()
		return
	}

	response["status"] = 200
	response["message"] = "Successfully get student"
	response["data"] = student
	stucon.Data["json"] = response
	stucon.ServeJSON()
	return
}

//GetStudentByEmail Get students by specific email
func (stucon *StudentController) GetStudentByEmail() {
	email := stucon.Ctx.Input.Param(":email")
	response := map[string]interface{}{}

	student := models.GetStudentByEmail(email)

	if student == nil {
		response["status"] = 404
		response["message"] = "Student not found"
		stucon.Data["json"] = response
		stucon.ServeJSON()
		return
	}

	response["status"] = 200
	response["message"] = "Successfully get student"
	response["data"] = student

	stucon.Data["json"] = response
	stucon.ServeJSON()
	return
}

//InsertStudent insert new student
func (stucon *StudentController) InsertStudent() {
	var student models.Student
	json.Unmarshal(stucon.Ctx.Input.RequestBody, &student)

	//Make validation
	validator := validation.Validation{}
	validator.Required(student.Name, "name")
	validator.Required(student.Email, "email")
	validator.Required(student.Address, "address")

	response := map[string]interface{}{}
	if validator.HasErrors() {
		response["status"] = 500
		response["message"] = "ERROR"
		response["error"] = validator.Errors
		stucon.Data["json"] = response
		stucon.ServeJSON()
		return
	}

	//Insert data
	err := models.InsertStudent(student)
	if err == nil {
		response["status"] = 201
		response["message"] = "Successfully insert students"
		stucon.Data["json"] = response
		stucon.ServeJSON()
		return
	}
}

//UpdateStudent Update some student data
func (stucon *StudentController) UpdateStudent() {
	var student models.Student
	json.Unmarshal(stucon.Ctx.Input.RequestBody, &student)

	id, _ := strconv.Atoi(stucon.Ctx.Input.Param(":id"))

	//Make validation
	validator := validation.Validation{}
	validator.Required(student.Name, "name")
	validator.Required(student.Email, "email")
	validator.Required(student.Address, "address")

	response := map[string]interface{}{}
	if validator.HasErrors() {
		response["status"] = 500
		response["message"] = "ERROR"
		response["error"] = validator.Errors
		stucon.Data["json"] = response
		stucon.ServeJSON()
		return
	}

	exists := models.GetStudentByID(id)
	if exists == nil {
		response["status"] = 404
		response["message"] = "Student not found"
		stucon.Data["json"] = response
		stucon.ServeJSON()
		return
	}
	err := models.UpdateStudent(student, id)
	if err == nil {
		response["status"] = 200
		response["message"] = "Successfully update students"
		stucon.Data["json"] = response
		stucon.ServeJSON()
		return
	}
}

//DeleteStudent Delete student data by specific id
func (stucon *StudentController) DeleteStudent() {
	//Convert string to int
	id, _ := strconv.Atoi(stucon.Ctx.Input.Param(":id"))
	response := map[string]interface{}{}

	err := models.DeleteStudent(id)
	if err == nil {
		response["status"] = 200
		response["message"] = "Successfully delete student"
		stucon.Data["json"] = response
		stucon.ServeJSON()
		return
	}
}
