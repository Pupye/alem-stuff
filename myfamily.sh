token=$(curl -s https://01.alem.school/assets/superhero/all.json)

echo $token | jq '.[0].connections.relatives' | sed 's/"//g'