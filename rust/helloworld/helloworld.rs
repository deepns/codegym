// A hello world program in Rust

// To compile and run:
// $ rustc helloworld.rs
// $ ./helloworld

// A note about the syntax:
// Rust uses snake case for function and variable names.
// functions are declared with fn
// main is the entry point of the program
fn main() {
    // println! is a macro that prints text to the console
    println!("Hello, world!");
    println!("I'm a Rustacean too!");
    let x = 5 + 5;
    println!("Is `x` 10 or 100? x = {}", x);
}

