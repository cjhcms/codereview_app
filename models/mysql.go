package models
//package main
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
    "errors"
    "time"
)

var (
  db *sql.DB
  db2 *sql.DB
)
func init() {
    dba, err := sql.Open("mysql", "root:@/shima_blog?charset=utf8")
    dbb, err2 := sql.Open("mysql", "root:@/shima_blog?charset=utf8")
    if err != nil || err2 != nil{
        log.Fatalf("Open database error: %s,%s\n", err,err2)
    }
    //defer db.Close()
    err = dba.Ping()
    err2 = dbb.Ping()
    if err != nil || err2 != nil {
        log.Fatal("Ping error: %s,%s\n",err,err2)
    }
    //全局变量与局部变量
    db = dba
    db2 = dbb
    fmt.Println("model init!")
}

type ResultObject struct{
  ClusterName string
  KpiName string
  KpiValue string
}
type MyModel struct{}
func (mm * MyModel) test(a string) (res string, err error) {
  return a + " hello", nil
}
func SaveArticle(params map[string]interface{}) (resultcode int,err error){
  fmt.Println("go in func:SaveArticle!")
  var title,create_time,mender,keywords,content interface{}
  stmt, err := db.Prepare("INSERT INTO code_review_artical(title,create_time,mender,keywords,content) VALUES (?,?,?,?,?)")
  defer stmt.Close()
  //var keys := [5]string{"title","create_time","mender","keywords","content"}
    if val, ok := params["title"]; ok {
      title = val
    }else{
      title = ""
    }
    if val, ok := params["create_time"]; ok {
      create_time = val
    }else{
      //go语言格式化处理时间
      timestamp := time.Now().Unix()
      tm := time.Unix(timestamp,0)
      create_time = tm.Format("2006-01-02 15:04:05")
      // if s, ok := create_time.(string); ok {
      //   fmt.Printf("create_time: %s \n" + s);
      // }
    }
    if val, ok := params["mender"]; ok {
      mender = val
    }else{
      mender = ""
    }
    if val, ok := params["keywords"]; ok {
      keywords = val
    }else{
      keywords = ""
    }
    if val, ok := params["content"]; ok {
      content = val
    }else{
      content = ""
    }
    res, err := stmt.Exec(title,create_time,mender,keywords,content)
    if err != nil {
      fmt.Printf("error ocurred : %s",res)
      fmt.Println("go out func:SaveArticle!")
      return 5011, err
    }
    fmt.Println("go out func:SaveArticle!")
    return 200,nil
}
func QueryDataByTime(params map[string]interface{}) (obj []ResultObject,err error){
  var stime interface{}
  var etime interface{}
  fmt.Println("models:")
  if value, ok := params["stime"]; ok {
    stime = value
    fmt.Println(stime)
  }
  if value, ok := params["etime"]; ok {
    etime = value
    fmt.Println(etime)
  }
  // stime := params["stime"]
  // etime := params["etime"]
  rows, err := db2.Query(`select cluster_name, kpi_name, kpi_value from m_machine_odps_withflag_kpi
    where collect_time  between ? and ?`,stime,etime)
  if err != nil {
      log.Println(err)
  }
  defer rows.Close()
  var ResultObjects = make([]ResultObject,1)
  var cluster_name string
  var kpi_name string
  var kpi_value string
  for rows.Next() {
      err := rows.Scan(&cluster_name, &kpi_name, &kpi_value)
      if err != nil {
          log.Fatal(err)
      }
      restmp := ResultObject{cluster_name,kpi_name,kpi_value}
      ResultObjects = append(ResultObjects,restmp)
  }
  err = rows.Err()
  if err != nil {
      log.Fatal(err)
      return nil ,errors.New("error");
  }
  fmt.Println(len(ResultObjects))
  return ResultObjects,nil
}

func insert(db *sql.DB) {
    stmt, err := db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
    defer stmt.Close()

    if err != nil {
        log.Println(err)
        return
    }
    stmt.Exec("guotie", "guotie")
    stmt.Exec("testuser", "123123")

}


func main() {
    db, err := sql.Open("mysql", "root:@/tesla_odps_view?charset=utf8")
    if err != nil {
        log.Fatalf("Open database error: %s\n", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    //insert(db)

    rows, err := db.Query("select cluster_name, kpi_name, kpi_value from m_machine_odps_withflag_kpi where id = ?", 20)
    if err != nil {
        log.Println(err)
    }

    defer rows.Close()
    var cluster_name string
    var kpi_name string
    var kpi_value string
    for rows.Next() {
        err := rows.Scan(&cluster_name, &kpi_name, &kpi_value)
        if err != nil {
            log.Fatal(err)
        }
        log.Println("*******************************")
        log.Println(cluster_name, kpi_name, kpi_value)
        fmt.Println(cluster_name, kpi_name, kpi_value)
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}
