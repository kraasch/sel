
run:
	go run .

hub_update:
	@hub_ctrl ${HUB_MODE} ln "$(realpath ./sel)" "${HOME}/.local/bin/sel"
