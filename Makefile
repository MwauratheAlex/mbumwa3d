# load .env variables
#
# NOTE: Comment this out while running in dev.
# - only activate when doing migrations coz: issues
#
# ifneq (, $(wildcard ./.env))
# 	include .env
# 	export
# endif

MAIN_PATH="tmp/bin/main"

# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
templ:
	@templ generate --watch --proxy="http://localhost:3000" --open-browser=false

server:
	@go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "go build --tags dev -o ${MAIN_PATH} ./cmd/" --build.bin "${MAIN_PATH}" --build.delay "100" \
		--build.exclude_dir "node_modules" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true

watch-assets:
	npx tailwindcss -i internal/assets/app.css -o public/assets/styles.css --watch

# run esbuild to generate the index.js bundle in watch mode.
watch-esbuild:
	npx esbuild internal/assets/index.js --bundle --outdir=public/assets --watch=forever

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
sync_assets:
	go run github.com/cosmtrek/air@v1.51.0 \
		--build.cmd "templ generate --notify-proxy" \
		--build.bin "true" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_dir "public" \
		--build.include_ext "js,css"

# start app in developement
dev:
	@make -j5  templ server watch-assets sync_assets watch-esbuild

# build the application for production. This will compile your app
# to a single binary with all its assets embedded.
# build:
# 	@npx tailwindcss -i app/assets/app.css -o ./public/assets/styles.css
# 	@npx esbuild app/assets/index.js --bundle --outdir=public/assets
# 	@go build -o bin/app_prod cmd/app/main.go
## Check if node and npm are installed, if not install them
check-node:
	@if ! command -v node > /dev/null 2>&1; then \
		curl -fsSL https://deb.nodesource.com/setup_14.x | bash - && \
		apt-get install -y nodejs; \
	fi

tailwind-build:
	npx tailwindcss -i internal/assets/app.css -o public/assets/style.min.css --minify

templ-generate:
	templ generate

build: check-node
	make tailwind-build
	npx esbuild internal/assets/index.js --bundle --outdir=public/assets
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/mbumwa3d ./cmd/main.go

# create new migrations
db-create:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) create $(filter-out $@,$(MAKECMDGOALS)) sql

db-up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

# Pattern rule to catch all undefined targets
%:
	@echo "winding up..."
