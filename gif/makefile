all: moving.gif

GIFS = d01.gif  d04.gif  d07.gif  d10.gif  d13.gif  d16.gif  d19.gif \
	d02.gif  d05.gif  d08.gif  d11.gif  d14.gif  d17.gif  d20.gif \
	d03.gif  d06.gif  d09.gif  d12.gif  d15.gif  d18.gif 

$(GIFS): d2b mkgifs
	./mkgifs

moving.gif: $(GIFS)
	convert -delay 35 -loop 0 d??.gif moving.gif

d2b: d2b.go
	go build d2b.go

clean:
	-rm -rf d2b *.gif *.load *.pts
