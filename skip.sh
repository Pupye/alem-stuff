declare -i y=1

for i in $(ls -1);

do rem=$(( $y % 2 ))

if [ $rem -eq 0 ]
then
  echo "$i"  
fi
((y++))
done; 


 
