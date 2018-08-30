package main

import (
    "fmt"
    "net"
    "os"
    "regexp"
    "strings"
    "encoding/json"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "8745"
    CONN_TYPE = "tcp"
)

var validCmdRegexp = regexp.MustCompile(`(GET|SET|DEL) `)
var validKeyRegexp = regexp.MustCompile(` [a-zA-Z-_0-9:]+`)
var jsonRegexp = regexp.MustCompile(` [{|\[].+`)
var db = map[string]string{}

func main() {

    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer l.Close()
    fmt.Println("Kiev started on " + CONN_HOST + ":" + CONN_PORT)
    
    // Listen for an incoming connection.
    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }

        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
  buf := make([]byte, 1024)

  _, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }

  // parse buffer into cmd, key, json document
  cmd := strings.TrimSpace(validCmdRegexp.FindString(string(buf)))
  key := strings.TrimSpace(validKeyRegexp.FindString(string(buf)))

  if(cmd == "") {
    conn.Write([]byte("Invalid Operation: Must be GET, SET, or DEL"))
    conn.Close()
    return
  }

  if(key == "") {
    conn.Write([]byte("Invalid Key: Must be [a-zA-Z-_0-9:]+"))
    conn.Close()
    return
  }

  if(cmd == "SET") {
    var doc interface{}
    stream := strings.TrimSpace(jsonRegexp.FindString(string(strings.Trim(string(buf), "\x00"))))
    err = json.Unmarshal([]byte(stream), &doc)

    if(err != nil) {
      fmt.Println(err)
      conn.Write([]byte("Invalid JSON Document"))
      conn.Close()
      return
    }

    db[key] = stream
    conn.Write([]byte("OK"))
  }

  if cmd == "GET" {
    conn.Write([]byte(db[key]))
  }

  if cmd == "DEL" {
    delete(db, key)
    conn.Write([]byte("OK"))
  }

  logOutput := []string{cmd, key, "..."}
  fmt.Println(strings.Join(logOutput, " "))

  conn.Close()
}