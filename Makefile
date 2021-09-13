.PHONY: rlnlib

rlnlib:
	sh scripts/build.sh
	#cd lib/rln && cargo build --release --target=aarch64-apple-darwin
	#cbindgen --config ../cbindgen.toml --crate rln --output librln.h --lang c
	#mv lib/rln/target/release/librln.* lib/

# @TODO TARGET FOR BINDINGS.