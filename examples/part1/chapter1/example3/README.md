# Hello World app

## Build it

```bash
go build -o app.exe *.go
ls -alh .

```

## Run it

```bash
./app.exe
```

### Output

```bash
$ go build -o app.exe *.go && ./app.exe
golang
myInt1+myInt2 = 2
myNb1/myNb2 = 2.3333333
myNb3/myNb4 = 2.054233389411802
myNb3*myNb4 = 24.71607073735472
myNb5/myNb6 = -0.6240236642384905
myNb5*myNb6 = -1179.1092802918636
myBool1 && myBool2 = false
myBool1 || myBool2 = true
!myBool1 =  false
Type: bool Value: true
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
0 0 false ""
42 Golang rocks Everybody knows "knows"
```
