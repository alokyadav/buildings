package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBuildings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Buildings Suite")
}
