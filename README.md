# chatapp
# chat
Aplikasi ini dibangun dengan
1. Bahasa Pemrograman : Golang
2. ORM : GORM
3. Framework : Echo 
4. Database : MySQL

Setting Bawaan untuk Database MySQL
1. Nama database : chat
2. User database : root
3. Password database : 1234Abcd

Cara Menjalankan
1. Jalankan aplikasi database anda, seperti XAMPP
2. Jalankan main.exe
3. Maka server Go akan berjalan dengan port 8000

Jawaban Test dengan End Point nya 

1. User can send a message to another user
  > end point : localhost:8000/sendchat
  * akses link dengan menggunakan post man
  * pilih body kemudian pilih form data
  * buat 3 key : sender dengan value id sender, receiver dengan value id receiver, message dengan value isi pesan
  * maka sistem akan mengecek apakah sender dan receiver pernah berkirim pesan atau belum
  * jika sudah maka akan diambil room id nya
  * jika belum maka akan dibuatkan room dan diambil room id nya
  
2. User can list all messages in a conversation between them and another user
  > end point : localhost:8000/listchats/1
  * satu pada end point diatas adalah id dari room yang digunakan untuk percakapan
  * jika diakses, maka akan mengubah status dari room id dari 0 menjadi 1 yang ada didalam tabel chat
  
3. User can reply to a conversation they are involved with
  > end point : localhost:8000/sendchat
  * cara menggunakan sama dengan send chat tinggal diubah sender menjadi id pengirim dan receiver menjadi id penerima
  * maka secara otomatis akan mengambil room id yang sebelumnya telah terbuat dan menambahkan record pada tabel chat
  
4. User can list all their conversations (if user A has been chatting with user C & D, the list for A will shows A-C & A-D)
  > end point : localhost:8000/listrooms/2
  * dua diatas adalah id dari user silahkan diganti untuk melihat list percakapan user tersebut dengan user yang lain
  * a. each conversation is accompanied by its last message
    * pada saat mengakses end point diatas maka secara otomatis last message yang akan muncul diatas untuk pertama kali karena sudah di sorting by system
  * b. each conversation is accompanied by unread count
