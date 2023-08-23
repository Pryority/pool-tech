install:
	cd cmd/web && npm install

chain:
	anvil

deploy:
	forge script cmd/server/contracts/script/Counter.s.sol:CounterScript --rpc-url http://127.0.0.1:8545 --broadcast

dev:
	air

start:
	docker build -t gowind .
	docker run -d --name gowind-app -p 8080:8080 gowind /app/goserver
	cd web && npm run dev

stop:
	docker stop gowind-app
	docker rm gowind-app
