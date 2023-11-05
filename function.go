package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
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
		dt := &TextMessage{
			To:       "6285155476774",
			IsGroup:  false,
			Messages: "123",
		}
		response := atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://wa.my.id/send/message/text")
		resp = response.Response
	} else {
		resp = "Secret Salah"
	}
	fmt.Fprintf(w, resp)

}

type TextMessage struct {
	To       string `json:"to"`
	IsGroup  bool   `json:"isgroup,omitempty"`
	Messages string `json:"messages"`
}
