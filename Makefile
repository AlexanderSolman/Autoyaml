CC := go build
BIN := autoyaml
TEST_YAML := *.yaml

.PHONY: build
build:
	$(CC) -o $(BIN) -a .

.PHONY: clean
clean:
	@rm -f $(BIN)

.PHONY: cyaml
cyaml:
	@rm -f $(TEST_YAML)

.PHONY: run
run:
	@./$(BIN)
