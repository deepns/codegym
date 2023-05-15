# Learning about RUST

## Getting started

- Installed RUST using [rustup](https://www.rust-lang.org/tools/install). **rustup** is the tool to install and manage RUST tool chains
- RUST metadata and toolchains installed at `$HOME/.rustup/`
- **cargo** is the official package manager for rust, installed at `$HOME/.cargo/`. All commands and binaries are added to `$HOME/.cargo/bin`
- List of commands after the rust installation. `rustc` to compile and generate the executable

```console
✗ ls $HOME/.cargo/bin 
cargo         cargo-fmt     clippy-driver rust-analyzer rust-gdbgui   rustc         rustfmt
cargo-clippy  cargo-miri    rls           rust-gdb      rust-lldb     rustdoc       rustup
```

## Rust by Example

Starting from [rust-by-example](https://doc.rust-lang.org/stable/rust-by-example/)

> Rust is a modern systems programming language focusing on safety, speed, and concurrency. It accomplishes these goals by being memory safe without using garbage collection.

- Interesting to see that memory safety ensure without using garbage collection? - don't know how it is done.

### Hello, World

Added a [hello world program](helloworld/helloworld.rs). Some observations.

- program begins with `main()`
- functions declared with **fn** keyword
- statements end with `;`
- variables and functions are named in snake case
- `println!` is a macro from [std::fmt](https://doc.rust-lang.org/std/fmt/) module
- supports line comments (`//` ) and block level comments (`/*...*/`)
- No inclusion of module or header for the hello world program. How macros from std::fmt are imported?

### Primitives

- Integer types, default i32. types can be expressed by default or using suffix
- interesting to see some rust specific nuances that I haven't seen before in the languages I used
  - variables are **immutable** by default. To declare a mutable variable, use `mut` keyword
  - variables can be re-declared in the same scope with same or different type (called **shadowing** in rust context)
  - numbers can be expressed with `_` for better readability. Compiler ignores `_`.
- `{:?}` - enables debug formatting of a type
