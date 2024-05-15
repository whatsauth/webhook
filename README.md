# Webhook Layanan WhatsAuth
WebHook Menggunakan method HTTP POST dengan Header bernama Secret  
![image](https://github.com/whatsauth/webhook/assets/11188109/7734295e-89bb-4b05-ab05-d2ee0bdb6019)  
Format JSON dalam body yang dikirim ke WebHook :  
![image](https://github.com/whatsauth/webhook/assets/11188109/c6454969-0700-4a33-a3b1-8d97e7ef0b8c)  
Lengkapnya :
```json
{
  "phone_number": "6281112000300",
  "reply_phone_number": "6281112000300",
  "chat_number": "6281112000300",
  "chat_server": "s.whatsapp.net",
  "group_name": "Example Group",
  "group_id": "G-123456",
  "group": "no",
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
Jika anda mengunakan Postman, bisa import [file json whatsauth](https://whatsauth.my.id/webhook/WhatsAuth.postman_collection.json)

## Contoh Source Code WebHook
### MongoDB
Untuk menjalankan salah satu contoh kode program yang ada di kami, kakak harus sudah punya MONGOSTRING atau akses ke database mongo untuk diinputkan ke kode program webhook.
Ikuti langkah berikut:
1. Sign up for mongodb.com and create one instance of Data Services of mongodb.
2. Go to Network Access menu > + ADD IP ADDRESS > ALLOW ACCESS FROM ANYWHERE  
   ![image](https://github.com/gocroot/gcp/assets/11188109/a16c5a73-ccdc-4425-8333-73c6fbf78e6d)  
3. Download [MongoDB Compass](https://www.mongodb.com/try/download/compass), connect with your mongo string URI from mongodb.com
4. Create database name iteung and collection reply  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/23ccddb7-bf42-42e2-baac-3d69f3a919f8)  
5. Import [this json](https://whatsauth.my.id/webhook/iteung.reply.json) into reply collection.  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/7a807d96-430f-4421-95fe-1c6a528ba428)  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/fd785700-7347-4f4b-b3b9-34816fc7bc53)  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/ef236b4d-f8f9-42c6-91ff-f6a7d83be4fc)  
6. Create a profile collection, and insert this JSON document with your 30-day token and WhatsApp number.  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/5b7144c3-3cdb-472b-8ab3-41fe86dad9cb)  
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/829ae88a-be59-46f2-bddc-93482d0a4999)  
   ```json
   {
     "token":"v4.public.asoiduasoijfiun98erjg98egjpoikr",
     "phonenumber":"6281111222333"
   }
   ```
   ![image](https://github.com/gocroot/alwaysdata/assets/11188109/06330754-9167-4bf4-a214-5d75dab7c60a)  

### Cloud Provider
Kami merekomendasikan penggunaan bahasa pemrograman golang guna mendukung komputasi hijau. 
Bagi kakak yang baru pertama kali belajar golang, bisa langsung saja ambil kursus di [W3School](https://www.w3schools.com/go/index.php), selesaikan semua excersises yang ada di sana.
Setelah itu, pilih salah satu contoh source code webhook untuk di deploy pilihan provider anda:
1. [Golang menggunakan free plan di AlwaysData.com](https://github.com/gocroot/alwaysdata/)
2. [Golang menggunakan AWS Lambda](https://github.com/gocroot/aws)
3. [Golang menggunakan Google Cloud Funtion](https://github.com/gocroot/gcp)
4. [Untuk PHP di Web Hosting](https://github.com/whatsauth/webhook/tree/main/hosting)


