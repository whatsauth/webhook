package webhook

import (
	"os"

	"github.com/aiteung/atdb"
)

var DBmongoinfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "waapi",
}

var Mongoconn = atdb.MongoConnect(DBmongoinfo)
