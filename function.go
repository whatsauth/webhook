package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/aiteung/module/model"
)

func init() {
	functions.HTTP("WebHook", WebHookPost)
}

func WebHookPost(w http.ResponseWriter, r *http.Request) {
	var imsg model.IteungMessage
	var resp string
	err := json.NewDecoder(r.Body).Decode(&imsg)
	if err != nil {
		resp = "error parsing application/json: " + err.Error()
	} else if r.Header.Get("Secret") == os.Getenv("SECRET") {
		resp = "berhasil"
	} else {
		resp = "Secret Salah"
	}
	fmt.Fprintf(w, resp)

}
