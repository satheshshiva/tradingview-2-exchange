build:
	go build .

run:
	go run .

deploy:
	gcloud -q app deploy

prerequisite: #sudo recommended. After install run chown -R username: <install-dir>
	curl https://sdk.cloud.google.com | zsh
	gcloud init

tail:
	gcloud app logs tail -s default

versions:
	gcloud app versions list

browse:
	gcloud app browse