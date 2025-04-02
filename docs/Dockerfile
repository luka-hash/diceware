from golang:1.23 as builder
env CGO_ENABLED=0
copy go.mod ./
copy main.go ./
copy wordlist.go ./
run go build -o /bin/diceware ./

from scratch 
copy --from=builder /bin/diceware /bin/diceware
cmd ["/bin/diceware"]
