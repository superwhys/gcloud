# GCloud

```bash
.
├── Dockerfile.server     backend server dockerfile
├── Dockerfile.web        frontend server dockerfile
├── README.md
├── docker-compose.yml    docker compose config file
├── mysql                 mysql data storage folder
├── nginx                 mysql data storage folder
├── server                backend server
└── web                   frontend server
```

# Start
Modify the Configs in `.env` 
```txt
COMPOSE_PROJECT_NAME=gcloud
MAIL_ACCOUNT=
MAIL_PASSWORD=
REDIS_ADDR=gcloud-redis:6379
MYSQL_ADDR=gcloud-mysql:3306
MYSQL_USER=gcloud
MYSQL_PASSWORD=toD9GEdVpDRq
TENCENT_SECRETKEY=RgzYHyaHJLYlAjP1ZwDQyTtiheXOVhow
TENCENT_SECRETID=AKIDUG9EKg1ICEMUK2dbepEhvGi7EB4um2Ja
```

Add your email address and key


run this to build and start all services
```bash
docker compose up -d
```

run this to stop all services
```bash
docker compose down
```

run this to start specify service
```bash
docker compose up -d web/server/mysql/redis
```

run this to stop specify service
```bash
docker compose down web/server/mysql/redis
```
