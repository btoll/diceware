PREFIX=/usr

build: diceware

diceware:
	go build

clean:
	rm -f diceware

# https://www.gnu.org/software/make/manual/html_node/DESTDIR.html
install:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -m 0755 diceware $(DESTDIR)$(PREFIX)/bin/

