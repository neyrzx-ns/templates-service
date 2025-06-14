generate:
	@./scripts/generate.sh

.PHONY: generate


# infra

infra-up:
	@docker compose -f deployments/docker-compose.yaml up -d
