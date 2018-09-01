package kiev

import(

)

const (
  CONN_TYPE = "tcp"
  CONN_PORT = 8745
)

func New(host string) {
}

func Get(key string) {
  // build protocol string "GET #{key}"
  // conn.Write()
  // return response
}

func Set(key string, json string) {
  // build protocol string "SET #{key} #{data}"
  // conn.Write()
  // return response
}

func Del(key string) {
  // build protocol string "DET #{key}"
  // conn.Write()
  // return response
}