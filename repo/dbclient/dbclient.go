package dbclient

import (
  "os"

  "gopkg.in/mgo.v2"
)

const DEFAULT_DBNAME = "rbdb"

type DbClient struct {
  Url     string
  Session *mgo.Session
  DbName  string
}

func New() DbClient {
  url := os.Getenv("MONGO_HOST")
  if len(url) == 0 {
    url = "localhost"
  }
  c := DbClient{
    Url:    url,
    DbName: DEFAULT_DBNAME,
  }
  return c
}
