CC      	= go
TARGET		= diceware
PREFIX		= /usr

.PHONY: build clean install

build: $(TARGET)

$(TARGET):
	$(CC) build

clean:
	rm -f $(TARGET)

# https://www.gnu.org/software/make/manual/html_node/DESTDIR.html
install:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -m 0755 diceware $(DESTDIR)$(PREFIX)/bin/

