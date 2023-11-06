package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/wa"
	"github.com/whatsauth/ws"
)

func WebHookPost(w http.ResponseWriter, r *http.Request) {
	var WAKeyword = "wh4t5auth0"
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		if strings.Contains(msg.Message, WAKeyword) && msg.From_link { //untuk whatsauth request login
			dt := &ws.WhatsauthRequest{
				Uuid:        strings.Replace(msg.Message, WAKeyword, "", 1),
				Phonenumber: msg.Phone_number,
				Delay:       msg.From_link_delay,
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/whatsauth/request")
		} else { //untuk membalas pesan masuk
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: "Hai hai hai kak " + msg.Alias_name,
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
		}
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, resp.Response)
}
