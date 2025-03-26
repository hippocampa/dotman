src := "main.go"
out_name := "dotshadow"
run:
	@go run $(src)

build:
	@go build -o $(out_name)

clean:
	@rm dotshadow
