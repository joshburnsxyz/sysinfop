GOX := $(shell which go)
BIN := sysinfop
PREFIX := /usr

sysinfop:
	$(GOX) build -o ./$(BIN)

clean:
	rm -f ./$(BIN)
	
install:
	mv ./$(BIN) $(PREFIX)/bin/$(BIN)
