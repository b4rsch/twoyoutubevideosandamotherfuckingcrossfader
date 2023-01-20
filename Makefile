build:
	docker build . -t tyvamfc

start:
	docker run -p "8089:8089" --rm tyvamfc
