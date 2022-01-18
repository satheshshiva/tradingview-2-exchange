build:
	go build .

run:
	go run .

deploy:
	gcloud app deploy

prerequisite: #sudo recommended. After install run chown -R username: <install-dir>
	curl https://sdk.cloud.google.com | bash
	gcloud init

tail:
	gcloud app logs tail -s default

versions:
	gcloud app versions list