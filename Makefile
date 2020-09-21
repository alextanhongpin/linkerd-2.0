# Run build-server or build-client.
build-%:
	cd go-$* && \
	docker build -t alextanhongpin/go-$* .


inject-%:
	cat go-$*/deployment.yaml | linkerd inject - | kubectl apply -f -

dashboard:
	@linkerd dashboard &
