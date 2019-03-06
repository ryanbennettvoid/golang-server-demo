package main

import (
  "testing"
  "time"

  "github.com/stretchr/testify/assert"
  "gitlab.com/codelittinc/golang-interview-project-ryan-bennett/repo"
  "gitlab.com/codelittinc/golang-interview-project-ryan-bennett/repo/dbclient"
)

func TestMain(t *testing.T) {

  // remove all members
  {
    c := dbclient.New()
    if err := c.Connect(); err != nil {
      panic(err)
    }
    defer c.Disconnect()
    c.Session.DB(c.DbName).C(repo.COLLECTION_MEMBERS).RemoveAll(nil)
  }

  assert := assert.New(t)

  assert.True(true)

  // create and start server
  server := NewServer()
  go func() {
    err := server.Listen()
    if err != nil {
      panic(err)
    }
  }()

  client := NewClient()

  {
    heartbeat, err := client.GetHeartbeat()
    assert.NoError(err)
    assert.Equal("hello", heartbeat.Message)
  }

  // get members (empty list)
  {
    members, err := client.GetMembers()
    assert.NoError(err)
    assert.Len(members, 0)
  }

  // create member
  newMember := repo.NewEmployee("John Doe", time.Now(), "Software Engineer")
  {
    err := client.CreateMember(newMember)
    assert.NoError(err)
  }

  // get member
  {
    _, err := client.GetMemberById(newMember.Id.Hex())
    assert.NoError(err)
    // a, err := newMember.ToJSON()
    // assert.NoError(err)
    // b, err := member.ToJSON()
    // assert.NoError(err)
    // assert.JSONEq(a, b)
  }

}
