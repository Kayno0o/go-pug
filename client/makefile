unocss:
ifeq ($(ENV),prod)
	pnpm unocss -m
else ifeq ($(ENV),dev)
	pnpm unocss -w
else
	@echo "Invalid environment specified. Use 'env=prod' or 'env=dev'."
	@exit 1
endif
