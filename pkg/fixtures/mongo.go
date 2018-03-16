package fixtures

import (
	"github.com/globalsign/mgo"
	"github.com/stretchr/testify/suite"
)

// DefaultURL is the default Mongo URL for testing
const DefaultURL = "127.0.0.1:27017"

// Mongo is a text fixture for use with github.com/stretchr/testify/suite tests
//
// use it like this:
//
//	type MyTests struct {
//		*fixtures.Mongo
//	}
//
//	func RunMyTests(t *testing.T) {
//		suite.Run(t, &MyTests{Mongo: New(DefaultURL)})
//	}
type Mongo struct {
	suite.Suite
	URL     string
	Session *mgo.Session
}

func (m *Mongo) SetupTest() {
	sess, err := mgo.Dial(m.URL)
	m.Require().NoError(err)
	m.Session = sess
}

func (m *Mongo) TearDownTest() {

}
