package main

import (
  "net/http"
  "net/http/httputil"
  "encoding/json"
  "log"
  "fmt"
  "time"
  "math/rand"
)

// type ReturnJson struct {
//   SampleStr string `json:"sample1"`
//   TestStr string `json:"sample2"`
//   TestInt int `json:"sample3"`
//   SampleArr []int `json:"sample4"`
//   SampleHash map[string]string `json:"sample5"`
// }

type ReturnJson struct {
  Status int  `json:"sample1"`
  Res1 string `json:"msg"`
  Res2 string `json:"price"`
  Body1 string `json:"bodymsg"`
  Body2 string `json:"bodyprice"`
}

func jsonHandler(writer http.ResponseWriter, req *http.Request) {
  defer req.Body.Close()

  // POSTリクエスト（HTTPBodyにパラメーターを設定）の場合
//   var returnJson ReturnJson
//   err1 := json.NewDecoder(req.Body).Decode(&returnJson)
//   if err1 != nil {
//     log.Fatal(err1)
//   }
  // POSTリクエスト（HTTPBodyにパラメーターを設定）の場合

  // GETリクエスト（クエリパラメーターを設定）の場合
  empty := ""
  returnJson := ReturnJson{http.StatusOK, req.FormValue("msg"), req.FormValue("price"), empty, empty}
  // GETリクエスト（クエリパラメーターを設定）の場合


  res, err1 := json.Marshal(returnJson)

  // ステータスコード表 https://golang.org/src/net/http/status.go?h=Request
  // ランダムに429と500を返すサンプルAPI
  err_msg := "error!!!"
  rand.Seed(time.Now().UnixNano())
  ran := rand.Intn(100)
  if ran % 20 == 0 {
    http.Error(writer, err_msg, http.StatusTooManyRequests)
    return
  } else if ran % 15 == 0 {
    http.Error(writer, err_msg, http.StatusInternalServerError)
    return
  } else {
    if err1 != nil {
      http.Error(writer, err1.Error(), http.StatusInternalServerError)
      return
    }

    dump, err2 := httputil.DumpRequest(req, true)
    fmt.Println(string(dump))

    if err2 != nil {
      http.Error(writer, err2.Error(), http.StatusInternalServerError)
      return
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.Write(res)
  }
}

func httpSample() {
  http.HandleFunc("/sample", jsonHandler)
  log.Println("start http listening :3000")
  http.ListenAndServe(":3000", nil)
}
