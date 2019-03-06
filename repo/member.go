package repo

import (
  "encoding/json"
  "errors"
  "time"

  "gopkg.in/mgo.v2/bson"
)

// dates are rounded a bit since mongo
// doesn't store with full precision
// (this helps with testing)
const DATE_PRECISION = time.Millisecond

const (
  MEMBER_ERROR_MISSING_NAME      = "member name is missing"
  MEMBER_ERROR_INVALID_TYPE      = "member type is invalid"
  MEMBER_ERROR_MISSING_STARTDATE = "member start date is missing"
  MEMBER_ERROR_MISSING_ENDDATE   = "member end date is missing"
  MEMBER_ERROR_MISSING_ROLE      = "member role is missing"
  MEMBER_ERROR_DATES_MISMATCH    = "member end date is before start date"
)

const (
  MEMBER_TYPE_UNKNOWN_START = iota
  MEMBER_TYPE_EMPLOYEE      = iota
  MEMBER_TYPE_CONTRACTOR    = iota
  MEMBER_TYPE_UNKNOWN_END   = iota
)

type Member struct {
  Id        bson.ObjectId `json:"_id" bson:"_id"`
  Name      string        `json:"name" bson:"name"`
  Type      int           `json:"type" bson:"type"`
  StartDate time.Time     `json:"start_date" bson:"start_date"`
  EndDate   time.Time     `json:"end_date" bson:"end_date"`
  Role      string        `json:"role" bson:"role"`
  Tags      []string      `json:"tags" bson:"tags"`
  Hidden    bool          `json:"hidden" bson:"hidden"`
}

func NewEmployee(name string, startDate time.Time, role string) Member {
  return Member{
    Id:        bson.NewObjectId(),
    Name:      name,
    Type:      MEMBER_TYPE_EMPLOYEE,
    StartDate: startDate.Round(DATE_PRECISION),
    Role:      role,
    Tags:      []string{},
    Hidden:    false,
  }
}

func NewContractor(name string, startDate time.Time, endDate time.Time) Member {
  return Member{
    Id:        bson.NewObjectId(),
    Name:      name,
    Type:      MEMBER_TYPE_CONTRACTOR,
    StartDate: startDate.Round(DATE_PRECISION),
    EndDate:   endDate.Round(DATE_PRECISION),
    Tags:      []string{},
    Hidden:    false,
  }
}

func (m *Member) Validate() error {

  if m.Type <= MEMBER_TYPE_UNKNOWN_START || m.Type >= MEMBER_TYPE_UNKNOWN_END {
    return errors.New(MEMBER_ERROR_INVALID_TYPE)
  }

  if m.Name == "" {
    return errors.New(MEMBER_ERROR_MISSING_NAME)
  } else if m.StartDate.IsZero() {
    return errors.New(MEMBER_ERROR_MISSING_STARTDATE)
  }

  switch m.Type {
  case MEMBER_TYPE_CONTRACTOR:
    // require an end date to calculate duration
    if m.EndDate.IsZero() {
      return errors.New(MEMBER_ERROR_MISSING_ENDDATE)
    }
  case MEMBER_TYPE_EMPLOYEE:
    // require a role
    if m.Role == "" {
      return errors.New(MEMBER_ERROR_MISSING_ROLE)
    }
  }
  return nil
}

func (m *Member) TermDuration() (time.Duration, error) {
  if m.StartDate.IsZero() {
    return time.Duration(0), errors.New(MEMBER_ERROR_MISSING_STARTDATE)
  } else if m.EndDate.IsZero() {
    return time.Duration(0), errors.New(MEMBER_ERROR_MISSING_ENDDATE)
  } else if m.EndDate.Before(m.StartDate) {
    return time.Duration(0), errors.New(MEMBER_ERROR_DATES_MISMATCH)
  }
  return m.EndDate.Sub(m.StartDate), nil
}

func (m *Member) ToJSON() (string, error) {
  data, err := json.Marshal(m)
  if err != nil {
    return "", err
  }
  return string(data), nil
}
