FROM rust:1.83.0-bullseye AS builder

WORKDIR /build

# trick to cache crates
RUN mkdir src && echo 'fn main() {}' > ./src/main.rs
COPY ./Cargo.* ./
RUN cargo build --release

COPY ./src ./src
RUN cargo build --release

FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder \
/build/target/release/main
./