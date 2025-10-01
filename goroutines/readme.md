time sleep itu memblokir goroutine, bukan program. 

Jika ada 3 goroutine yang jalan maka time sleep akan memblokir 1 goroutine hingga ia bangun dan melanjutkan
pekerjaannya, semenatara goroutine lain sudah seleisan terlebih dahulu.