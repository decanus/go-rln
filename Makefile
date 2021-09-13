.PHONY: rlnlib

rlnlib:
	cd lib/rln && cargo build --release
	cbindgen --config ../cbindgen.toml --crate rln --output librln.h --lang c
	mv lib/rln/target/release/librln.* lib/

# @TODO TARGET FOR BINDINGS.