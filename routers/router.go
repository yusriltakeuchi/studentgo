// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"studentgo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// ns := beego.NewNamespace("/v1",
	// 	// beego.NSNamespace("/object",
	// 	// 	beego.NSInclude(
	// 	// 		&controllers.ObjectController{},
	// 	// 	),
	// 	// ),
	// 	beego.NSNamespace("/student",
	// 		beego.NSInclude(
	// 			&controllers.ObjectController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(ns)
	beego.Router("/api/student", &controllers.StudentController{}, "get:GetStudents")
	beego.Router("/api/student", &controllers.StudentController{}, "post:InsertStudent")

	beego.Router("/api/student/:id:int", &controllers.StudentController{}, "put:UpdateStudent")
	beego.Router("/api/student/:id", &controllers.StudentController{}, "get:GetStudent")
	beego.Router("/api/student/email/:email", &controllers.StudentController{}, "get:GetStudentByEmail")
	beego.Router("/api/student/:id", &controllers.StudentController{}, "delete:DeleteStudent")
}
