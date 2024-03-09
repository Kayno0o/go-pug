docker:
ifeq ($(ENV),prod)
	pnpm tsc
	pnpm postcss -m
else ifeq ($(ENV),dev)
	pnpm postcss -w > ./logs/postcss.log 2>&1 &
	pnpm tsc --watch > ./logs/tsc.log 2>&1
else
	@echo "Invalid environment specified. Use 'env=prod' or 'env=dev'."
	@exit 1
endif
