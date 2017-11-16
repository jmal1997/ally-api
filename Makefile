.PHONY: clean

clean:
	rm -f cli
	rm -f rehash
	rm -f *.xml.xml
	rm -f examples/*.xml.xml

cli:
	gofmt -w ./cmd/cli
	go build github.com/jackmanlabs/ally-api/cmd/cli

rehash:
	gofmt -w ./cmd/rehash
	go build github.com/jackmanlabs/ally-api/cmd/rehash