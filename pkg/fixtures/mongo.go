package fixtures

import (
	"github.com/globalsign/mgo"
	"github.com/stretchr/testify/suite"
	"github.com/technosophos/moniker"
)

var names = moniker.NewAlliterator()

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
	url    string
	dbName string
	DB     *mgo.Database
}

func (m *Mongo) SetupTest() {
	sess, err := mgo.Dial(m.url)
	m.Require().NoError(err)
	m.DB = sess.DB(m.dbName)
}

func (m *Mongo) TearDownTest() {
	m.Require().NoError(m.DB.DropDatabase())
}

func New(url string) *Mongo {
	dbName := names..NameSep("-")
	return &Mongo{
		Suite: suite.New(),
		url: url,
		dbName: names.NameSep("-") + "-athens-testing",
	}
}
