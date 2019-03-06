package main

import (
  "errors"
  "time"
)

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
  Name      string
  Type      int
  StartDate time.Time
  EndDate   time.Time
  Role      string
  Tags      []string
}

func NewEmployee(name string, startDate time.Time, role string) Member {
  return Member{
    Name:      name,
    Type:      MEMBER_TYPE_EMPLOYEE,
    StartDate: startDate,
    Role:      role,
  }
}

func NewContractor(name string, startDate time.Time, endDate time.Time) Member {
  return Member{
    Name:      name,
    Type:      MEMBER_TYPE_CONTRACTOR,
    StartDate: startDate,
    EndDate:   endDate,
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
