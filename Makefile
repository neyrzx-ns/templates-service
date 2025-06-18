generate:
	@./scripts/generate.sh

.PHONY: generate


# infra

infra-up:
	@docker compose -f deployments/docker-compose.yaml up -d


# migrations
migration-create:
	 @read -p "name: " migration_name; ./scripts/migrations/create.sh $$migration_name


.PHONY: migration-create
