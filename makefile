build:
	bash scripts/build.sh $(tag)

start:
	docker stack deploy -c docker-compose.yaml gcl