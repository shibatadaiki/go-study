package main

import (
  "net/http"
  "net/http/httputil"
  "encoding/json"
  "log"
  "fmt"
)

// type ReturnJson struct {
//   SampleStr string `json:"sample1"`
//   TestStr string `json:"sample2"`
//   TestInt int `json:"sample3"`
//   SampleArr []int `json:"sample4"`
//   SampleHash map[string]string `json:"sample5"`
// }

type ReturnJson struct {
   Status int    `json:"sample1"`
   Result string `json:"sample2"`
   }

func jsonHandler(writer http.ResponseWriter, r *http.Request) {
  returnJson := ReturnJson{http.StatusOK, "ok"}

  res, err1 := json.Marshal(returnJson)

  if err1 != nil {
    http.Error(writer, err1.Error(), http.StatusInternalServerError)
    return
  }

  dump, err2 := httputil.DumpRequest(r, true)
  fmt.Println(string(dump))

  if err2 != nil {
    http.Error(writer, err2.Error(), http.StatusInternalServerError)
    return
  }

  writer.Header().Set("Content-Type", "application/json")
  writer.Write(res)
}

func httpSample() {
  http.HandleFunc("/sample", jsonHandler)
  log.Println("start http listening :3000")
  http.ListenAndServe(":3000", nil)
}
