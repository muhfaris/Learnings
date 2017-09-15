#!/bin/bash
# convert string to ascii
# convert ascii to string
# by Faris .Afif
  
print('::::::::::::::::::::::')
print('::1. Convert to ASCII')
print('::2. Convert to STRING')
print(':--------------------:')
print(':  author by Faris   :')
print('::::::::::::::::::::::')
pil = input("Masukkan Pilihan 1/2 = ")
if pil =='1':
	mes = input("Masukkan Pesan : ")
	sem = []
	print("Encode String to ASCII : ")
	for c in mes:
		sem.append(ord(c))
		print(sem)
if pil =='2':
	pesan = input("Masukkan kode ASCII : ")
	print("Decode ASCII to String : ")
	for c in pesan:
		print(decode(c))
