package utils

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestUtils(t *testing.T) {
	suite.Run(t, new(utilsTestSuite))
}

type utilsTestSuite struct {
	suite.Suite
}
