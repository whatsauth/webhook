package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/ws"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var WAKeyword = "wh4t5auth0"
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		if ws.IsLoginRequest(msg, WAKeyword) { //untuk whatsauth request login
			resp = HandlerQRLogin(msg, WAKeyword)
		} else { //untuk membalas pesan masuk
			resp = HandlerIncomingMessage(msg)
		}
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, resp.Response)
}
