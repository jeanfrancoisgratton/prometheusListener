all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./prometheusSDlistener.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.23.1 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/sbin || mkdir -p $(DESTDIR)$(PREFIX)/sbin
	install -pm 755 $(BUILDDIR)/prometheusSDlistener.exe $(DESTDIR)$(PREFIX)/sbin/prometheusSDlistener
	strip $(DESTDIR)$(PREFIX)/sbin/prometheusSDlistener
