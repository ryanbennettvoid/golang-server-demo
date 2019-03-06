package repo

import (
  "gitlab.com/codelittinc/golang-interview-project-ryan-bennett/repo/dbclient"
  "gopkg.in/mgo.v2/bson"
)

const COLLECTION_MEMBERS = "members"

func GetMembers() ([]Member, error) {
  c := dbclient.New()
  if err := c.Connect(); err != nil {
    panic(err)
  }
  defer c.Disconnect()

  collection := c.Session.DB(c.DbName).C(COLLECTION_MEMBERS)
  query := bson.M{"hidden": false}

  var members []Member
  err := collection.Find(query).All(&members)
  if err != nil {
    return []Member{}, err
  }

  return members, nil
}

func InsertMember(inMember Member) error {
  c := dbclient.New()
  if err := c.Connect(); err != nil {
    panic(err)
  }
  defer c.Disconnect()

  collection := c.Session.DB(c.DbName).C(COLLECTION_MEMBERS)
  err := collection.Insert(&inMember)
  if err != nil {
    return err
  }

  return nil
}

func GetMemberById(id string) (Member, error) {
  c := dbclient.New()
  if err := c.Connect(); err != nil {
    panic(err)
  }
  defer c.Disconnect()

  collection := c.Session.DB(c.DbName).C(COLLECTION_MEMBERS)
  query := bson.M{
    "id":     bson.ObjectIdHex(id),
    "hidden": false,
  }

  var member Member
  err := collection.Find(query).One(&member)
  if err != nil {
    return Member{}, err
  }

  return member, nil
}
