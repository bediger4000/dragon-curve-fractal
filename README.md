# Dragon Curve Fractal

I wanted to see if I could draw the famous
[Dragon Curve fractal(https://en.wikipedia.org/wiki/Dragon_curve) using Go.

## Iterated Function System

Initially I thought I would go through creating a line segment,
folding it, and then folding the newly created line segments.

It turns out you can draw the Dragon Curve using an "Iterated
Function System".


   f1(z) = (1 + i)z/2
   f2(z) = 1 - (1 - i)z/2

In this, z is a complex number, a+bi.

I took this to mean that for any complex number you start with,
and you're supposed to start with 0+0i and 1+0i,
you get 2 additional complex numbers.
You plug those 2 additional numbers into f1() and f2(),
and get 4 complex numbers,
and so on and so forth.

I did the calculations recursively, except I skipped recursive
calls on the result of f2(z) if it equaled f1(z).
There's a lot of duplicated points in the output,
and since I used [gnuplot](http://www.gnuplot.info/) to do output, I sorted and uniqued
the outputs.

## The fractal

![dragon curve](dragon.png?raw=true)

I used [gnuplot](http://www.gnuplot.info/) to do the plotting,
which is about as easy as it gets.
I did match the pixel width and height of the image to the
x and y widths of the Dragon Curve to avoid gnuplot distorting it.

The plaid effect in the image is a result of mapping floating
point numbers to integer-indexed pixels. I thought it quite fetching.

## Recreating the fractal

You need a Go compiler.
I did this under Linux, you should too.
To recreate the image:

    $ make
