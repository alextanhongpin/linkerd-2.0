build-server:
	cd go-server && \
	docker build -t alextanhongpin/go-server .

inject-linkerd:
	cat go-server/deployment.yaml | linkerd inject - | kubectl apply -f -

dashboard:
	@linkerd dashboard &
