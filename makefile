build:
	git pull origin master
	buildah bud --rm --no-cache -t onn-sh:latest .

run: stop
	docker run -d -p 8080:80 --name=onn-sh --restart=always onn-sh:latest

stop:
	docker rm -f onn-sh || true

dev-run: stop
	docker run -p 8080:8080 --name=onn-sh --restart=always onn-sh:latest