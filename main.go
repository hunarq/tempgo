package main

import (
  "fmt"
  "strconv"
  "log"
  "net/http"
)

type pounds float32
type degrees float32

func (p pounds) String() string {
  return fmt.Sprintf("£%.2f", p)
}

type database map[string]pounds

func (d database) foo(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "foo: %s\n", d["foo"])
}

func (d database) f2c(w http.ResponseWriter, r *http.Request) {
  param1 := r.URL.Query().Get("p1")
  a, err := strconv.ParseInt( param1, 10, 64);
  if err != nil {
    // handle the error in some way
  }

  p1 := float32(a - 32)*5/9 
  fmt.Fprintf(w, "%f C\n", p1)
}

func (d database) c2f(w http.ResponseWriter, r *http.Request) {
  param1 := r.URL.Query().Get("p1")
  a, err := strconv.ParseInt( param1, 10, 64);
  if err != nil {
    // handle the error in some way
  }
  p1 := 32+ float32(a)*9/5 
  fmt.Fprintf(w, "%f F\n", p1)
}

func main() {
  db := database{
    "foo": 1,
    "bar": 2,
    "baz": 3,
  }

  http.HandleFunc("/foo", db.foo)
  http.HandleFunc("/f2c", db.f2c)
  http.HandleFunc("/c2f", db.c2f)
  

  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

