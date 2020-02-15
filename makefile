all: dragon.png dragon2.png

dragon.png: d3.pts dragon.load
	gnuplot < dragon.load

dragon2.png: a.pts b.pts dragon2.load
	gnuplot < dragon2.load

d3.pts: d3
	./d3 17 | sort | uniq > d3.pts

a.pts b.pts: d2b
	./d2b 20 0 0  > c.pts
	./d2b 20 1 0  > d.pts
	cat c.pts d.pts > x.pts
	grep '1$$' x.pts > a.pts
	grep '2$$' x.pts > b.pts

d3: d3.go
	go build d3.go

d2b: d2b.go
	go build d2b.go

clean:
	-go clean
	-rm -rf d3.pts x.pts a.pts b.pts c.pts d.pts
	-rm -rf *.png
