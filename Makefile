.PHONY: rlnlib
rlnlib:
	cd lib/rln && cargo build --release
	cp lib/rln/target/release/librln.* lib/