package webhook

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

func TestGetENVToken(t *testing.T) {
	fmt.Println(os.Getenv("MONGOSTRINGAWANGGA"))
}

func TestUpdateGetData(t *testing.T) {
	rply, err := atdb.GetRandomDoc[Reply](Mongoconn, "reply", 1)
	fmt.Println(err)
	fmt.Println(rply)
}
