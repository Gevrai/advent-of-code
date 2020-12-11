#!/bin/env bash
echo "$((2#$(sed -E 's/B|R/1/g;s/F|L/0/g' input.txt|sort|tail -1)))"
echo $[2#$(tr 'BR' '1'<input.txt|tr 'FL' '0'|sort|tail -1)]

for i in $(sed -E 's/B|R/1/g;s/F|L/0/g' input.txt|sort);{ [[ $[$p+2] == $[2#$i] ]] && echo $[$p+1];p=$[2#$i];}
