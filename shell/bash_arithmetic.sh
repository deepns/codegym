#! /usr/bin/env bash

## Learning bash arithmetic

# Using let

# default usage, with no spaces
let a=1+9 # assigns a=10
echo "a=$a"

# If spacing needed for better readability, enclose the expression between quotes
let "b = 2 + 8" # assigns b=10
echo "b=$b"

# Take values from variables
let c=$a+$b # assigns sum of a and b to c
echo "c=$c"

# No escaping needed for *
let d=$a*$b # assigns product of a and b to d
echo "d=$d"

# incr, decr
let d++ # increment d by 1 and assign to d
echo $d

# Using expr

# evaluate the expression and writes to standard output
expr 2 + 7

# multiplication operator * need to be escaped
expr 7 * 3 # syntax error

expr 7 \* 3 # evaluates to 21

# evaluate and assign
e=$(expr 2 + 7)
echo "e=$e"

# invalid expressions. printed as is without evaluating
expr 2+7 # doesn't evaluate the expression

expr "2 + 7" # doesn't evaluate the expression

# Using double parantheses
echo $((3+5)) # prints 8

# spaces are totally fine
# evaluate the expression and assign the value to a variable
f=$((4 + 4))
echo "f=$f"

# variables do not need $ when used calculating with (( ))
g=$((f * 2))
echo "g=$g"

# increment
((g--))
echo "g=$g"

((g += 5))
echo "g=$g"
