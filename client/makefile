docker:
ifeq ($(ENV),prod)
	pnpm postcss -m
else ifeq ($(ENV),dev)
	pnpm postcss -w
else
	@echo "Invalid environment specified. Use 'env=prod' or 'env=dev'."
	@exit 1
endif
