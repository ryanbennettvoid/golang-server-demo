package dbclient

import (
  "gopkg.in/mgo.v2"
)

func (c *DbClient) Connect() error {
  session, err := mgo.Dial(c.Url)
  if err != nil {
    return err
  }
  c.Session = session
  c.Session.SetMode(mgo.Monotonic, true)
  return nil
}
