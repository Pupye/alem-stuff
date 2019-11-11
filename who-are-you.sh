token=$(curl https://01.alem.school/assets/superhero/all.json)
echo $token | jq '.[] | select(.id == 70) | .name'