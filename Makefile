# Up on the daemon develop
develop:
	docker-compose -f docker-compose.yml build
	docker-compose up -d
final:
	docker-compose -f docker-compose.prod.yml build
	docker-compose up -d
# stop
stop:
	docker-compose stop
# down
down:
	docker-compose down

# prune
container-prune:
	docker container prune
network-prune:
	docker network prune
builder-prune:
	docker builder prune
images-prune:
	docker image prune