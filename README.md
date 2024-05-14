# Webhook Layanan WhatsAuth
WebHook Menggunakan method HTTP POST dengan Header bernama Secret  
![image](https://github.com/whatsauth/webhook/assets/11188109/7734295e-89bb-4b05-ab05-d2ee0bdb6019)  
Format JSON dalam body yang dikirim ke WebHook :  
![image](https://github.com/whatsauth/webhook/assets/11188109/c6454969-0700-4a33-a3b1-8d97e7ef0b8c)  
Lengkapnya :
```json
{
  "phone_number": "1234567890",
  "reply_phone_number": "0987654321",
  "chat_number": "1122334455",
  "chat_server": "server1.chat.example",
  "group_name": "Example Group",
  "group_id": "G-123456",
  "group": "yes",
  "alias_name": "JohnDoe",
  "messages": "This is a test message",
  "from_link": true,
  "from_link_delay": 30,
  "is_group": true,
  "filename": "document.pdf",
  "filedata": "base64EncodedString",
  "latitude": 37.7749,
  "longitude": -122.4194,
  "liveloc": true
}
```

## Contoh Source Code WebHook
Pilih salah satu contoh source code webhook untuk di deploy:
1. [Golang menggunakan free plan di AlwaysData.com](https://github.com/gocroot/alwaysdata/)
2. [Golang menggunakan AWS Lambda](https://github.com/gocroot/aws)
3. [Untuk Go di Google Cloud Funtion](https://github.com/gocroot/gcp)
4. [Untuk PHP di Web Hosting](https://github.com/whatsauth/webhook/tree/main/hosting)


