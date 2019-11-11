token=$(curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"Pupye\"}}){id}}"}')
echo $token | jq '.data.user[0].id'
