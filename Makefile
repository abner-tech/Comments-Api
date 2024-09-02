.PHONY: run
run:
	@echo 'Running Application...'
	@go run ./cmd/api -port=4000 -env=production