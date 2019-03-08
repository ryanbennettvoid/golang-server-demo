package main

import (
  "testing"
  "time"

  "github.com/ryanbennettvoid/golang-server-demo/repo"
  "github.com/ryanbennettvoid/golang-server-demo/repo/dbclient"
  "github.com/stretchr/testify/assert"
)

/*

  API Integration Test

*/

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
  newMember.Tags = []string{"Javascript", "Golang"}
  {
    err := client.CreateMember(newMember)
    assert.NoError(err)
  }

  // get member
  {
    member, err := client.GetMemberById(newMember.Id.Hex())
    assert.NoError(err)
    a, err := newMember.ToJSON()
    assert.NoError(err)
    b, err := member.ToJSON()
    assert.NoError(err)
    assert.JSONEq(a, b)
  }

  newRole := "Management"
  newMember.Role = newRole
  newMember.Tags = append(newMember.Tags, "C++")
  // update member
  {
    err := client.UpdateMemberById(newMember.Id.Hex(), newMember)
    assert.NoError(err)
  }

  // get member again
  {
    member, err := client.GetMemberById(newMember.Id.Hex())
    assert.NoError(err)
    a, err := newMember.ToJSON()
    assert.NoError(err)
    b, err := member.ToJSON()
    assert.NoError(err)
    assert.JSONEq(a, b)
  }

  // get members (expect 1)
  {
    members, err := client.GetMembers()
    assert.NoError(err)
    assert.Len(members, 1)
  }

  // create member
  newContractor := repo.NewContractor("John Doe", time.Now(), time.Now().Add(time.Hour*24*90))
  {
    err := client.CreateMember(newContractor)
    assert.NoError(err)
  }

  // get members (expect 2)
  {
    members, err := client.GetMembers()
    assert.NoError(err)
    assert.Len(members, 2)
  }

  // delete member
  {
    err := client.DeleteMemberById(newContractor.Id.Hex())
    assert.NoError(err)
  }

  // get members (expect 1)
  {
    members, err := client.GetMembers()
    assert.NoError(err)
    assert.Len(members, 1)
  }

}
