# peek-in-pi
Something (late) for PI day, very simple program that allows you to "peek in pi"
just like [this website](http://pi.fathom.info/).

It allows you to find at what index (in the first million digits of PI) your provided sequence
appears.

## Quick start
Usage: ./main [SEQUENCE]

### Compile
```console
> go build main.go
```

### Run
```console
> main 69420
> Found 69420709223... at: 15773
```