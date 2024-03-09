unocss:
ifeq ($(ENV),prod)
	pnpm postcss -m
else ifeq ($(ENV),dev)
	pnpm postcss -w
else
	@echo "Invalid environment specified. Use 'env=prod' or 'env=dev'."
	@exit 1
endif

typescript:
ifeq ($(ENV),prod)
	pnpm tsc
else ifeq ($(ENV),dev)
	pnpm tsc --watch
else
	@echo "Invalid environment specified. Use 'env=prod' or 'env=dev'."
	@exit 1
endif
