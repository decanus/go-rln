DIRECTORY=./libs
if [[ -d "$DIRECTORY" ]]
then
    echo "$DIRECTORY exists on your filesystem. Delete it and run the script again."
    exit 0
fi

pushd lib/rln

export RUSTFLAGS="-Ccodegen-units=1"
rustup default 1.52.1
cargo install cargo-lipo -f
cargo install cargo-strip -f

rustup target add x86_64-unknown-linux-gnu
rustup target add aarch64-unknown-linux-gnu
rustup target add x86_64-apple-darwin
rustup target add aarch64-apple-darwin
rustup target add x86_64-pc-windows-gnu
rustup target add aarch64-linux-android
rustup target add armv7-linux-androideabi
rustup target add i686-linux-android
rustup target add x86_64-linux-android
rustup target add x86_64-unknown-linux-musl
rustup target add aarch64-apple-ios x86_64-apple-ios

cargo build --release --target=aarch64-linux-android --lib
cargo strip --target aarch64-linux-android
cargo build --release --target=armv7-linux-androideabi --lib
cargo strip --target armv7-linux-androideabi
cargo build --release --target=i686-linux-android --lib
cargo strip --target i686-linux-android
cargo build --release --target=x86_64-linux-android --lib
cargo strip --target x86_64-linux-android
cargo build --target=x86_64-unknown-linux-gnu --release
cargo strip --target x86_64-unknown-linux-gnu
cargo build --target=aarch64-unknown-linux-gnu --release
cargo strip --target aarch64-unknown-linux-gnu
cargo build --target=x86_64-apple-darwin --release
cargo strip --target x86_64-apple-darwin
cargo build --target=aarch64-apple-darwin --release
cargo strip --target aarch64-apple-darwin
cargo build --target=x86_64-pc-windows-gnu --release
cargo strip --target x86_64-pc-windows-gnu
cargo build --target=x86_64-unknown-linux-musl --release
cargo strip --target x86_64-unknown-linux-musl
cargo lipo --release --targets=aarch64-apple-ios,x86_64-apple-ios

popd

TOOLS_DIR=`dirname $0`
COMPILE_DIR=${TOOLS_DIR}/../lib/rln/target
rm -rf $COMPILE_DIR/x86_64-apple-ios $COMPILE_DIR/aarch64-apple-ios
for platform in `ls ${COMPILE_DIR} | grep -v release | grep -v debug`
do
  PLATFORM_DIR=${DIRECTORY}/$platform
  mkdir -p ${PLATFORM_DIR}
  cp ${COMPILE_DIR}/$platform/release/librln.* ${PLATFORM_DIR}
done
