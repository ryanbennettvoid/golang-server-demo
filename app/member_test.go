package main

import (
  "testing"
  "time"

  "github.com/stretchr/testify/assert"
)

const MEMBER_NAME = "John Doe"
const MEMBER_ROLE = "Software Engineer"
const THREE_MONTHS = time.Hour * 24 * 90

func makeValidEmployee() Member {
  return NewEmployee(MEMBER_NAME, time.Now(), MEMBER_ROLE)
}

func makeValidContractor() Member {
  start := time.Now()
  end := start.Add(THREE_MONTHS)
  return NewContractor(MEMBER_NAME, start, end)
}

// --- Employee

func TestMemberNewEmployee(t *testing.T) {
  assert := assert.New(t)
  member := makeValidEmployee()
  assert.NoError(member.Validate())
}

func TestMemberEmployeeMissingName(t *testing.T) {
  assert := assert.New(t)
  member := makeValidEmployee()
  member.Name = ""
  assert.Error(member.Validate())
}

func TestMemberEmployeeMissingStartDate(t *testing.T) {
  assert := assert.New(t)
  member := makeValidEmployee()
  member.StartDate = time.Time{}
  assert.Error(member.Validate())
}

func TestMemberEmployeeMissingRole(t *testing.T) {
  assert := assert.New(t)
  member := makeValidEmployee()
  member.Role = ""
  assert.Error(member.Validate())
}

func TestMemberEmployeeInvalidType(t *testing.T) {
  assert := assert.New(t)
  member := makeValidEmployee()
  // --
  member.Type = MEMBER_TYPE_UNKNOWN_START
  assert.Error(member.Validate())
  member.Type = MEMBER_TYPE_UNKNOWN_END
  assert.Error(member.Validate())
  member.Type = -5
  assert.Error(member.Validate())
  member.Type = 5
  assert.Error(member.Validate())
}

// --- Contractor

func TestMemberNewContractor(t *testing.T) {
  assert := assert.New(t)
  member := makeValidContractor()
  assert.NoError(member.Validate())
}

func TestMemberContractorMissingName(t *testing.T) {
  assert := assert.New(t)
  member := makeValidContractor()
  member.Name = ""
  assert.Error(member.Validate())
}

func TestMemberContractorMissingStartDate(t *testing.T) {
  assert := assert.New(t)
  member := makeValidContractor()
  member.StartDate = time.Time{}
  assert.Error(member.Validate())
}

func TestMemberContractorMissingEndDate(t *testing.T) {
  assert := assert.New(t)
  member := makeValidContractor()
  member.EndDate = time.Time{}
  assert.Error(member.Validate())
}

func TestMemberContractorInvalidType(t *testing.T) {
  assert := assert.New(t)
  member := makeValidContractor()
  assert.NoError(member.Validate())
  // --
  member.Type = MEMBER_TYPE_UNKNOWN_START
  assert.Error(member.Validate())
  member.Type = MEMBER_TYPE_UNKNOWN_END
  assert.Error(member.Validate())
  member.Type = -5
  assert.Error(member.Validate())
  member.Type = 5
  assert.Error(member.Validate())
}
