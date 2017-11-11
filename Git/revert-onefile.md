## Mengembalikan file lama tertentu
Contoh disini akan mengambalikan file lama dengan nama `config.json`

### Lihat `ID log` terlebih dahulu
perintah yang digunakan `git log`

### Cek file
perintah yang digunakan `git checkout id_log`
cek apakah file sudah benar dan ingat alamat path filenya

### Kembalikan file lama dengan branch yang sudah ada
Jadi mengembalikan fila yang sudah di cek sebelumnya, dengan branch yang sudah
ada. Jadi perubahan terbaru tidak hilang

perintah : `git checkout id_log -- path_file`

### Lakukan git commit
Pesan yang akan digunakan untuk mengambalikan file lama

### Lakukan push
