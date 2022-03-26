# Websites Checker Instruction

## Prerequisite
- docker
- docker compose
- csv file

## Local Deployment
build & deploy
```sh
cd exam-webchecker
docker-compose up -d --build
```
down local servers
```sh
docker-compose down
```

## Manaul Test
1. after deploying these services, go to the url `http://localhost:3001/websites-checker`
2. Upload the examples of csv files that we already have in this project. (`test.csv`, `top500Domains.csv`)
3. See the result.


Thank you so much for this opportunity :)
