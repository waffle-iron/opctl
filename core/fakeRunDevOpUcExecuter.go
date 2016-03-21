// This file was generated by counterfeiter
package core

import (
  "sync"
  "github.com/dev-op-spec/engine/core/models"
)

type fakeRunDevOpUcExecuter struct {
  ExecuteStub        func(devOpName string) (devOpRun models.DevOpRunView, err error)
  executeMutex       sync.RWMutex
  executeArgsForCall []struct {
    devOpName string
  }
  executeReturns     struct {
                       result1 models.DevOpRunView
                       result2 error
                     }
}

func (fake *fakeRunDevOpUcExecuter) Execute(devOpName string) (devOpRun models.DevOpRunView, err error) {
  fake.executeMutex.Lock()
  fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
    devOpName string
  }{devOpName})
  fake.executeMutex.Unlock()
  if fake.ExecuteStub != nil {
    return fake.ExecuteStub(devOpName)
  } else {
    return fake.executeReturns.result1, fake.executeReturns.result2
  }
}

func (fake *fakeRunDevOpUcExecuter) ExecuteCallCount() int {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return len(fake.executeArgsForCall)
}

func (fake *fakeRunDevOpUcExecuter) ExecuteArgsForCall(i int) string {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return fake.executeArgsForCall[i].devOpName
}

func (fake *fakeRunDevOpUcExecuter) ExecuteReturns(result1 models.DevOpRunView, result2 error) {
  fake.ExecuteStub = nil
  fake.executeReturns = struct {
    result1 models.DevOpRunView
    result2 error
  }{result1, result2}
}
