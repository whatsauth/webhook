# PHP Hosting

Untuk yang menggunakan PHP di web hosting bisa memakai contoh kode [berikut](index.php).
Hal yang pertama dilakukan adalah, pastikan hosting anda tidak memblokir domain atau ip dari alamat api.wa.my.id
Cara pengecekan bisa lewat SSH dari hosting, contoh jika hosting memblokir alamat api wa maka akan seperti ini:  
![image](https://github.com/whatsauth/webhook/assets/11188109/6c58afed-d8aa-4fa6-a1d4-a35a27ed7e6c)  
atau bisa juga ketika melakukan test menggunakan postman ke webhook di hosting, akan terasa lama dan muncul keterangan seperti ini:  
![image](https://github.com/whatsauth/webhook/assets/11188109/17676d2d-b1b3-4b54-ad99-4c9637f8b6fa)

## Langkah langkah
Buka Cpanel dan masuk ke File Manager buat folder baru bernama webhook di dalam public_html  
![image](https://github.com/whatsauth/webhook/assets/11188109/1a39bd75-1f86-4b38-a068-8becc87f087e)  
Kemudian buat file index.php di dalam folder webhook yang tadi dibuat  
![image](https://github.com/whatsauth/webhook/assets/11188109/a90824d4-f75e-4948-97c9-c2f1d6e19780)  
Kemudian klik edit file index.php yang tadi kita buat  
![image](https://github.com/whatsauth/webhook/assets/11188109/d5d348ab-17af-4c31-abbe-9e53ac54d919)  
Paste kan kode yang dicontohkan di atas, edit bagian SECRET_TOKEN dan TOKEN  
![image](https://github.com/whatsauth/webhook/assets/11188109/70022ace-aa63-48fb-9cca-14c42578a402)  
Klik simpan perubahan, kita cek langsung buka URL dari file tersebut tampak seperti ini  
![image](https://github.com/whatsauth/webhook/assets/11188109/3f244557-70db-4a01-8fe7-26644f38b970)  
Mari kita test dengan PostMan atau Thunder Client

