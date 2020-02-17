all:
	@echo "building image..."
	@sleep 3
	@docker build -t ascii-art . >> log.txt
	@echo "OK!"
	@echo "Finding image...\n"
	@echo "- - - - - - - - -"
	@docker images | grep ascii-art
	@echo "- - - - - - - - -"
	@echo "Running container in daemon mode..."
	@docker run -d -p 8080:5500 ascii-art | tee -a log.txt
	@echo "Done."