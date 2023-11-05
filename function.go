package gcf

import (
	"fmt"
	"os"
	"net/http"
	"encoding/json"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/petapedia/peda"
)

func init() {
	functions.HTTP("PetaPedia", petaPediaPost)
}

func petaPediaPost(w http.ResponseWriter, r *http.Request) {
	var imsg model.IteungMessage
	var resp string
	token := os.Getenv("TOKEN")
	err := json.NewDecoder(r.Body).Decode(&imsg)
	if err != nil {
		resp = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Selamat Datang"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Salah"
		}
	}
	fmt.Fprintf(w, peda.GCFPostHandler("PASETOPRIVATEKEY", "MONGOULBI", "petapedia", "user", r))

}
