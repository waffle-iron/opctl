// This file was generated by counterfeiter
package core

import (
  "sync"
  "github.com/dev-op-spec/engine/core/models"
)

type fakeAddDevOpUcExecuter struct {
  ExecuteStub        func(req models.AddDevOpReq)
  executeMutex       sync.RWMutex
  executeArgsForCall []struct {
    req models.AddDevOpReq
  }
}

func (fake *fakeAddDevOpUcExecuter) Execute(req models.AddDevOpReq) (err error) {
  fake.executeMutex.Lock()
  fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
    req models.AddDevOpReq
  }{req})
  fake.executeMutex.Unlock()
  if fake.ExecuteStub != nil {
    fake.ExecuteStub(req)
  }

  return
}

func (fake *fakeAddDevOpUcExecuter) ExecuteCallCount() int {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return len(fake.executeArgsForCall)
}

func (fake *fakeAddDevOpUcExecuter) ExecuteArgsForCall(i int) models.AddDevOpReq {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return fake.executeArgsForCall[i].req
}
