.PHONY: deploy
deploy:
	@docker image rm -f bien-pa-bot;\
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .;\
docker build -t bien-pa-bot .;\
rm main;\
docker rm -f bien-pa-bot;\
docker run -d --restart unless-stopped --name bien-pa-bot --env-file config.env bien-pa-bot;\

.PHONY: start
start:
	@docker run -d --restart unless-stopped --name bien-pa-bot --env-file config.env bien-pa-bot

.PHONY: stop
stop:
	@docker stop bien-pa-bot

.PHONY: logs
logs:
	@tail -f -n 2000 logs.txt

