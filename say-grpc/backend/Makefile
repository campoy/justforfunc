build:
	GOOS=linux go build -o app
	docker build -t gcr.io/justforfunc-prep/say .
	rm -f app

push:
	gcloud docker -- push gcr.io/justforfunc-prep/say