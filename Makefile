.PHONY: bench, test

bench:
	cd bench; go test -bench=.

test:
	cd test; go test .