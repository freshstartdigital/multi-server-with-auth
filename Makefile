development:
	docker compose -f infra/development/docker-compose.yml up -d && cd app && npm run dev

development-down:
	docker compose -f infra/development/docker-compose.yml down

development-build:
	docker compose -f infra/development/docker-compose.yml build --no-cache
	docker compose -f infra/development/docker-compose.yml up -d

development-rebuild:
	docker compose -f infra/development/docker-compose.yml build --no-cache
	docker compose -f infra/development/docker-compose.yml up -d
	cd app && npm run dev

development-logs:
	docker compose -f infra/development/docker-compose.yml logs -f

testing:
	docker compose -f infra/testing/docker-compose.yml up -d

testing-down:
	docker compose -f infra/testing/docker-compose.yml down

testing-build:
	docker compose -f infra/testing/docker-compose.yml up -d --build

docs:
	cd docs && npm run docs:dev