package tests

import (
	"github.com/goravel/framework/testing"

	"CambodiaTaxPortal/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
