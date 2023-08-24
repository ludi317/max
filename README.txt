See https://medium.com/@ludirehak/printing-lol-doubled-the-speed-of-my-go-code-e32e02fc3f92 for a walkthrough and discussion.




To run benchmarks:
1) Adjust input (values slice) to desired size and distribution (random, increasing, decreasing) in max_test.go

2) Run the following commands:
$ go test -c -o max.test && ./max.test -test.bench . -test.v -test.count 6  | tee bench/raw/increasing.txt && benchstat -col /name bench/raw/increasing.txt | tee bench/stats/increasing.txt
$ benchstat -col /name bench/raw/increasing.txt | tee bench/stats/increasing.txt


To use pprof:
1) Add `-test.cpuprofile cpu.out` to the go test command

2) $ go tool pprof -http=":" max.test cpu.out

3) In the browser, click on the Source
