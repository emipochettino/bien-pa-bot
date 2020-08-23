.PHONY: deploy
deploy:
	@docker image rm -f bien-pa-bot;\
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .;\
docker build -t bien-pa-bot .;\
rm main;\
docker run --env-file config.env bien-pa-bot &>>logs.txt;\

.PHONY: start
start:
	@docker run -d --env-file config.env bien-pa-bot

.PHONY: stop
stop:
	@docker stop bien-pa-bot

.PHONY: logs
logs:
	@tail -f -n 2000 logs.txt

