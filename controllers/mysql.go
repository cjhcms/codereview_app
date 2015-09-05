package controllers

import (
	"codereview_app/models"
	//"encoding/json"
  //"strconv"
  "fmt"
	"github.com/astaxie/beego"
	"reflect"
	"errors"
)

// Operations about object
type MysqlController struct {
	beego.Controller
}
var routerFunc = map[string]interface{} {
	//"querybyid":models.QueryByID,
	"querybytime":models.QueryDataByTime,
	"savearticle":models.SaveArticle,
}
//反射方法的应用：如何解决？？？？？
//如果返回值类型为（result []reflect.value,err error）,无法获取函数真正应该返回的值
// func Call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value, err error) {
func Call(m map[string]interface{}, name string, params ... map[string]interface{}) (result []reflect.Value, err error) {
    f := reflect.ValueOf(m[name])
    if len(params) != f.Type().NumIn() {
        err = errors.New("The number of params is not adapted.")
        return
    }
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
		fmt.Println(in)
		// result := make([]reflect.Value,10)
    result = f.Call(in)
		fmt.Println("result:")
		fmt.Println(result[0])
		return
    // return result[0].Interface(),nil
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (m *MysqlController) Get() {
		/**
		* 获取参数的方法：
		1.直接通过input.Params获取；
	  2.通过input.Request.Form获取
		*/
    objectId := m.Ctx.Input.Params[":objectId"]
		if m.Ctx.Input.Request.Form == nil {
			  m.Ctx.Input.Request.ParseForm()
		}
		stime := m.Ctx.Input.Request.Form.Get("stime")
		etime := m.Ctx.Input.Request.Form.Get("etime")
		var req_params = make(map[string]interface{},1)
    for k,v := range m.Ctx.Input.Request.Form {
			fmt.Printf("key:%s:value:%s \n",k,v[0])
			req_params[k] = v[0]
		}
		fmt.Println(stime)
		fmt.Println(etime)
    // if objectId != "" {
    //   oid, err:= strconv.Atoi(objectId)
    //   fmt.Println(oid)
    //   if err != nil {
    //     //m.Data["json"] = err
    //   }
			/**
			*通过反射方法调用函数获取方法
			*/
			// ob1,err := Call(routerFunc,"querybytime",oid,stime,etime)
      if objectId == "" {
				m.Data["json"] = "u are request with wrong url."
			}else{
				fmt.Printf("objectId:    %s\n",objectId)
				ob1,err := Call(routerFunc,objectId,req_params)
				/**
				*原始方法调用函数获取结果
				*/
	      //ob, err := models.QueryDataByTime(oid,stime,etime)
	      if err != nil {
	        fmt.Println("err:::")
	        m.Data["json"] = err
	      }else{
	        fmt.Println("success!")
					ob := ob1[0].Interface()
	        m.Data["json"] = ob
	      }
	    }
	    //fmt.Println(m.Data["json"])
	    //m.Ctx.Output.Json(m.Data["json"], true, true)
    m.ServeJson()
}
// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [Post]
func (m * MysqlController) Post(){
  objectId := m.Ctx.Input.Params[":objectId"]
	if m.Ctx.Input.Request.Form == nil {
			m.Ctx.Input.Request.ParseForm()
	}
	var req_params = make(map[string]interface{},1)
	for k,v := range m.Ctx.Input.Request.Form {
		fmt.Printf("key:%s:value:%s \n",k,v[0])
		req_params[k] = v[0]
	}
	ob1,err := Call(routerFunc,objectId,req_params)
	if err != nil {
		m.Data["json"] = err
	}else{
		ob := ob1[0].Interface()
		m.Data["json"] = ob
	}
	m.ServeJson()
}
