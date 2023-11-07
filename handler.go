package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atdb"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/whatsauth/wa"
	"github.com/whatsauth/ws"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var WAKeyword = "wh4t5auth0"
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		if ws.IsLoginRequest(msg, WAKeyword) { //untuk whatsauth request login
			dt := &ws.WhatsauthRequest{
				Uuid:        ws.GetUUID(msg, WAKeyword),
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

func PostWithDB(w http.ResponseWriter, r *http.Request) {
	var WAKeyword = "wh4t5auth0"
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		if ws.IsLoginRequest(msg, WAKeyword) { //untuk whatsauth request login
			dt := &ws.WhatsauthRequest{
				Uuid:        ws.GetUUID(msg, WAKeyword),
				Phonenumber: msg.Phone_number,
				Delay:       msg.From_link_delay,
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/whatsauth/request")
		} else { //untuk membalas pesan masuk
			rply, _ := atdb.GetRandomDoc[Reply](Mongoconn, "reply", 1)
			replymsg := strings.ReplaceAll(rply[0].Message, "#BOTNAME#", msg.Alias_name)
			replymsg = strings.ReplaceAll(replymsg, "\\n", "\n")
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: replymsg,
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
		}
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, resp.Response)
}
