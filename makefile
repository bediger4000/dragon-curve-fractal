
dragon.png: d3.pts dragon.load
	gnuplot < dragon.load

d3.pts: d3
	./d3 17 | sort | uniq > d3.pts

d3: d3.go
	go build d3.go

clean:
	-go clean
	-rm -rf d3.pts
	-rm -rf *.png
