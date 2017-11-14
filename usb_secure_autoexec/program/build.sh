mkdir ./build
cargo build --release
cp target/release/usb_secure_autoexec ./build
cp ./run_scripts.sh ./build
