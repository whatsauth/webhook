package webhook

import (
	"os"

	"github.com/aiteung/atdb"
)

var DBmongoinfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRINGAWANGGA"),
	DBName:   "waapi",
}

var Mongoconn = atdb.MongoConnect(DBmongoinfo)
