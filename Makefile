NAME=$(shell basename ${PWD})

clean:
	rm ${NAME}

build:
	go build -o $(NAME)

run:
	go run main.go < in

check:
	$(shell diff <(./${NAME} < in) out)