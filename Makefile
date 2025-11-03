task_vantara: main.go
	go build -o build/app
	cp .env build/

clean:
	rm -rf build/*

run: build/app
	./build/app