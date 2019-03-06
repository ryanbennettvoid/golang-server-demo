package dbclient

import (
  "errors"
)

func (c *DbClient) Disconnect() error {
  if c.Session == nil {
    return errors.New("attempted to disconnect but no session found")
  }
  c.Session.Close()
  return nil
}
