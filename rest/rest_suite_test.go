package rest

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "testing"
)

func TestReSt(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "ReST Suite")
}
