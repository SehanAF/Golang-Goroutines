
/**
== Pengenalan Parallel Programming ==
~ Saat ini kita hidup di era multicore, dimana jarang sekali kita menggunakan prosseror yang single core.
~ Semakin canggih perangkat keras, maka software pun akan mengikuti, dimana sekarang kita bisa dengan mudah membuat       proses parallel di aplikasi.
~ Parallel programming sederhananya adalah memecahkan suatu masalah dengan cara membaginya menjadi yang lebih keci,dan dijalankan secara bersamaan pada waktu yang bersamaan pula.

== Contoh Parallel ==
~ Menjalankan beberapa aplikasi sekaligus di sistem operasi kita ( office,editor,browser, dan lain-lain )
~ Beberapa koki menyiapkan makanan di restoran, dimana tiap koki membuat makanan masing-masing.
~ Antrian di Bank, dimana tiap teller melayani nasabahnya masing-masing.

== Process VS Thread ==
= Process =
~ Process adakah sebuah eksekusi program.
~ Process mengkomsumsi memory besar.
~ Process saling terisolasi dengan process lain.
~ Process lama untuk dijalankan dihentikan.

= Thread =
~ Thread adalah sebuah eksekusi program.
~ Thread menggunakan memory kecil.
~ Thread bisa saling berhubungan jika dalam process yang sama.
~ Thread cepat untuk dijalankan dan dihentikan.

== Parallel VS Concurrency ==
~ Berbeda dengan paralel (menjalankan beberapa perkerjaan secara bersamaan), concurrency adalah menjalankan beberapa pekerjaan secara bergantian.
~ Dalam parallel kita biasanya membutuhkan banyak Thread, sedangkan dalam concurrench, kita hanya membutuhkan sedikit thread.

== Contoh Concurrency ==
~ Saat kita makan dicafe, kita bisa makan, lalu ngobrol, lalu minum, makan lagi, ngobrol lagi, minum lagi dan seterusnya. Tetapi kita tidak bisa pada saat yang bersamaan munum,makan dan ngobrol, hanya bisa melakuan satu hal pada satu waktu, namun bisa berganti kapanpun kita mau.

== CPU-bound ==
~ Banyak algoritma dibuat yang hanya membutuhkan CPU untuk menjalankannya, Algoritma jenis ini biasanya sangat tergantuungn dengan kecepatan CPU.
~ Contoh yang paling populer adakah Machine Learning, oleh karena itu sekarang banyak sekali teknologi Machine Learning yang banyak menggunakan CPU karena memiliki core yang lebih banyak dibanding CPU biasanya.
~Jenis algoritma seperti ini tidak ada benefitnya menggunakan Concurrency Programming, namun bisa dibantu dengan implementasi Parallel Programming.

== I/O-bound
~ I/0-bound adalah kebalikan dari sebelumnya, dimana biasanya algoritma atau aplikasinya sangat tergantung dengan kecepatan input output device yang digunakan.
~ Comtohnya aplikasi seperti membaca data dari file, database, dan lain-lain.
~ Kebanyakan saat ini, biasanya kita akan membuat aplikasi jenis seperti ini.
~ Aplikasi jenis I/O-bound, walaupun bisa terbantu dengan implementasi Parallel Programming, tapi benefitnya akan lebih baik jika menggunakan Concurrency Programming.
~ Bayangkan kita membaca data dari database dan Thread harus menunggu 1 detik untuk mendapat balasan dari database, padahal waktu 1 detik itu jika menggunakan Concurrency Programming, bisa digunakan untuk melakukan hal lain lagi.

== Pengenalan Goroutines ==
~ Goroutines adalah sebuah thread ringan yang dikelola oleh Go Runtime.
~ Ukuran Goroutine sangat kecil, sekitar 2kb, jauh lebih kecil dibandingkan Threadd yang bisa sampai 1mb atau 1000kb.
~ Namun tidak seperti thread yang berjalan parallel, goroutine berjalan secara concurrrent.

== Cara Kerja Goroutine ==
~ Sebenarnya, Goroutine dijalankan oleh Go Schedduler dalam thread, dimana jumlah thread nya sebanyak GOMAXPROCS (biasanya sejumlah core CPU).
~ Jadi sebenarnya tidak bisa dibilang Goroutine itu pengganti Thread, karena Goroutine sendiri berjalan diatas Thread.
~ Namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua sudah diatur oleh Go Scheduler.

== Cara Kerja Go Scheduler ==
~ Dalam Go-Scheduler, kita akan mengenal beberapa terminologi:
> G : Goroutine
> M : Thread (Machine)
> P : Processor

== Membuat Goroutine ==
~ Untuk membuat goroutine di Golang sangatlah sederhana.
~ Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine.
~ Saat sebuat function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai.
~ Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai.

== Goroutine Sangat Ringan
~ Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan.
~ Kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory.
~ Tidak seperti thread yang ukurannya berat, goroutine sanagatlah ringan.

== Pengenalan Channel ==
~ Channel adalah tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine.
~ Di Channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda.
~ Saat melakukan pengiriman data ke Channel, goroutine akan ter-block, sampai ada yang menerima data tersebut.
~ Maka dari itu, channel disebut sebagai alat komunikasi sychronous(blocking).
~ Channel cocok sekali sebagai alternatif seperti mekanisme async await yang terdapat di beberapa bahasa pemrograman lain.

== Karakteristik Channel ==
~ Secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi, harus menunnggu data yang ada di channel diambil.
~ Channel hanya bisa menerima satu jenis data.
~ Channel bisa diambil dari lebih dari satu goroutine.
~ Channel harus di close jika tidak digunakan, atau bisa menyababkan memory leak.

== Membuat Channel ==
~ Channel di Go-Lang direpresentasikan dengan tipe data chan.
~ Untuk membuat channel sangat mudah, kita bisa menggunakan make(), mirip ketika membuat map.
~ Namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukan kedalam channel tersebut.

== Membuat dan Menerima Data dari Channel
~ Seperti yang sudah dibahas sebelumnya, channel bisa dinakan untuk mengirim dan menerima data.
~ Untuk mengirim data,kita bisa gunakan kode: channel<-data
~ Sedangkan untuk menerima data, bisa gunakan kode: data <- channel
~ Jika selesai, jangan lupa untuk menutup channel menggunakan function close()

== Channel Sebagai Parameter ==
~ Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter.
~ Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah pass by value,artinya value akan diduplikasi lalu dikirim ke function paramter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer(agar pass by reference).
~Berbeda dengan Channel,Kita tidak perlu melakukan hal tersebut.

== Channel In dan Out ==
~ Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut.
~ Kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk mengirim data,atau hanya dapat digunakan untuk menerima data.
~ Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk(mengirim data) atau out(menerima data) 

== Cara mengeksekusi Data ==
~ go test -v -run=TestVariableChannel 

== Range Channel ==
~ Kadang-kadang ada kasus sebauh channel ikirm data secara terus menerus oleh pengirim.
~ Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data.
~ Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel.
~ Ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti.
~ Ini lebih sederhana pada kita melakukan pengecekan channel secara manual.

== Select Channel == 
~ Kadan ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine. 
~ Lalu kita ingin mendapatkan data dari channel tersebut.
~ Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Go-Lang.
~ Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersmaan di beberapa channe;, maka akan dipilih secara random.

== Default Select ==
~ Apa yang terjadi jika kita melakukan select terhadap channel yang ternyata tidak ada datanya?
~ Maka kita akan menunggu sampai data ada.
~ Kadang mungkin kita ingin melakukan sesuatu jika misal semua channel tidak ada datanya ketika kita melakuka select channel.
~ Dalam select, kita bisa menambahkan default, dimana ini akan dieksekusi jika memang di semua channel yang kita select tidak ada datanya.

== Race Condition ==
= Masalah Dengan Goroutine =
~ Saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel.
~ Hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan.
~ Hal ini bisa menyebabkan masalah yang namanya Race Condition.

== Mutex (Mutual Exclusion) ==
~ Untuk mengatasi masalah race condition tersebut, di Go-Lang terdapat sebuah struct bernama sync.Mutex
~ Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking terhadap mutex, maka tidak ada yang bisa melukan locking lagi sampai kita melakukan unlock.
~ Dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang diperbolehkan, setelah goroutine tersebut melakukan unlock, baru goroutine selanjutnya diperbolehkan melakukan lock lagi.
~ Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi.








*/