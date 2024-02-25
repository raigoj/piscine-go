curl -s https://01.kood.tech/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Raigo\"}}){id}}"}' | cut -d ":" -f4 | cut -d "}" -f1 
