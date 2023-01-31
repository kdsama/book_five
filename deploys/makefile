install-docker:
	curl -fsSL https://get.docker.com -o get-docker.sh
	sh get-docker.sh
install-redis:
	docker pull bitnami/redis:latest
run-redis:
	docker run -it -d -p 6379:6379 -e ALLOW_EMPTY_PASSWORD=yes --name redis --restart always bitnami/redis:latest 
install-postgres:
	docker pull bitnami/postgresql:latest
run-postgres:
	docker run -it -d -p 5432:5432 -e POSTGRESQL_USERNAME=admin -e POSTGRESQL_PASSWORD=admin123 --restart always -e POSTGRESQL_DATABASE=book_five -v /home/kshitij/postgres:/bitnami/postgresql  --name postgres bitnami/postgresql:latest 
install-mongodb:
	docker pull bitnami/mongodb:latest
run-mongodb:
	docker run -it -d -p 27017:27017 --name mongodb -v /home/kshitij/postgres:/bitnami/mongodb --restart always bitnami/mongodb:latest

