serve:
	go run app/*.go -serve

export:
	go run app/*.go -export

.PHONY: serve export
	
