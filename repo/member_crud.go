package repo

import (
  "encoding/json"
  "time"

  "gitlab.com/codelittinc/golang-interview-project-ryan-bennett/repo/dbclient"
  "gopkg.in/mgo.v2/bson"
)

const COLLECTION_MEMBERS = "members"
const DATE_LAYOUT = "2006-01-02T15:04:05.000-07:00"
const ZERO_DATE = "0001-01-01T00:00:00Z"

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
    "_id":    bson.ObjectIdHex(id),
    "hidden": false,
  }

  var member Member
  err := collection.Find(query).One(&member)
  if err != nil {
    return Member{}, err
  }

  return member, nil
}

func UpdateMemberById(id string, inMember Member) error {

  c := dbclient.New()
  if err := c.Connect(); err != nil {
    panic(err)
  }
  defer c.Disconnect()

  var props map[string]interface{}
  data, err := json.Marshal(inMember)
  if err != nil {
    return err
  }
  err = json.Unmarshal(data, &props)
  if err != nil {
    return err
  }
  if _, ok := props["_id"]; ok {
    delete(props, "_id")
  }

  // change date strings to data objects
  fields := []string{"start_date", "end_date"}
  for _, field := range fields {
    if _, ok := props[field]; ok {
      val := props[field].(string)
      t := time.Time{}
      if val != ZERO_DATE {
        timeObj, err := time.Parse(DATE_LAYOUT, val)
        if err != nil {
          return err
        }
        t = timeObj
      }
      props[field] = t
    }
  }

  collection := c.Session.DB(c.DbName).C(COLLECTION_MEMBERS)
  query := bson.M{
    "_id":    bson.ObjectIdHex(id),
    "hidden": false,
  }
  action := bson.M{
    "$set": props,
  }

  err = collection.Update(query, action)
  if err != nil {
    return err
  }

  return nil
}

func DeleteMemberById(id string) error {

  c := dbclient.New()
  if err := c.Connect(); err != nil {
    panic(err)
  }
  defer c.Disconnect()

  collection := c.Session.DB(c.DbName).C(COLLECTION_MEMBERS)
  query := bson.M{
    "_id":    bson.ObjectIdHex(id),
    "hidden": false,
  }
  action := bson.M{
    "$set": bson.M{
      "hidden": true,
    },
  }

  err := collection.Update(query, action)
  if err != nil {
    return err
  }

  return nil
}
