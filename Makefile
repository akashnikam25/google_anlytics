server:
	@cd cmd/gotracker && go build && ./gotracker -ip 103.249.241.32

dashboard:
	@cd cmd/dashboard && \
	go build -o localdash && \
	./localdash -site 1 -start 20231101 -end 20231130