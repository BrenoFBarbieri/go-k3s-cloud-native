.DEFAULT_GOAL := help

.PHONY: help infra build push deploy restart logs pods dev redeploy

APP_NAME := go-api
VERSION := v1
REGISTRY := localhost:5000

help:
	@echo "Usage: make [command]\n"
	@echo "Available commands:"
	@echo "  dev       - Run the full development workflow (infra, build, push, deploy)"
	@echo "  redeploy  - Rebuild, push, and restart the application deployment"
	@echo "  logs      - View logs of the application pods"
	@echo "  pods      - List all pods in the cluster"
	@echo "  down      - Stop the k3s cluster"
	@echo "  clean     - Stop and remove the k3s cluster"
	@echo "  purge     - Stop, remove, and delete all images and volumes associated with the k3s cluster"
	@echo ""

infra:
	docker compose up -d
	@echo "Waiting for k3s to start..."
	@until docker exec k3s kubectl get pods > /dev/null 2>&1; do \
		sleep 2; \
	done

build:
	docker build -t ${REGISTRY}/${APP_NAME}:${VERSION} api/

push:
	docker push ${REGISTRY}/${APP_NAME}:${VERSION}

deploy: build push
	cat k8s/manifests/*.yaml | docker exec -i k3s kubectl apply -f -

restart:
	docker exec k3s kubectl rollout restart deployment/${APP_NAME}

logs:
	docker exec k3s kubectl logs -l app=${APP_NAME} -f

pods:
	docker exec k3s kubectl get pods -o wide

dev: infra build push deploy

redeploy: build push restart

down:
	docker compose down

clean:
	docker compose down -v

purge:
	@echo "WARNING: This will remove all containers, volumes, and images associated with the k3s cluster."
	@read -p "Are you sure you want to proceed? (y/N) " confirm && [ "$$confirm" = "y" ]
	docker compose down -v --rmi all
