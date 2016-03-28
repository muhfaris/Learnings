#!/usr/bin/env python

##
# Learning - 5
# I'm learning with my self, so i'm sorry if my language not good.
# 
# More on Strings
#  
# function array in string
# function find
# function split
##

print("Learning-5") #print text "Learning -5"
print("String on Array")
strings = ['i','love','python','easy','learn','and','learn','again']
for a in strings:
	print(a)
print("")

print("Function Find")
string = "I love Python" #this sentence
find = string.find('love') #finding word of sentence
print(find) # if value 2 its true, i love python is 3 word.
print("")


print("Function Split")
print("From = ",string)
print("To = ", string.split( )) #split of setence with single quote
