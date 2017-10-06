package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SomeTestSuite struct {
	suite.Suite
	SomeVar int
}

func (suite *SomeTestSuite) SetupTest() {
	suite.SomeVar = 7
}

func (suite *SomeTestSuite) TearDownTest() {
	// some cleanup
}

func (suite *SomeTestSuite) TestSome() {
	assert.Equal(suite.T(), 7, suite.SomeVar)
}
func (suite *SomeTestSuite) TestSome2() {
	assert.Equal(suite.T(), 7, suite.SomeVar)
}

func TestSomeTestSuite(t *testing.T) {
	suite.Run(t, new(SomeTestSuite))
}
