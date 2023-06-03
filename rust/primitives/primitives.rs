// Learning about RUST primitives
use std::any::type_name;

fn main() {
    // Variables can be type annotated.
    let success: bool = true;
    println!("success = {}", success);

    // Integer types can be expressed by suffix, or by default.
    // default integer type is i32
    let decimal = 1000;
    println!("decimal = {}", decimal);
    
    // There are various types of integers: i8, u8, i16, u16, i32, u32, i64, u64, isize, usize
    // isize and usize depend on the kind of computer your program is running on: 64 bits
    // if you're on a 64-bit architecture and 32 bits if you're on a 32-bit architecture.
    let decimal: i64 = 10000;

    // Rust allows a variable to be redeclared within the same scope
    // This is called shadowing.
    // Shadowing is the act of re-declaring a variable with the same
    // name as an existing variable in the same scope. In Rust, this
    // is allowed and can be used to change the type or value of a
    // variable while keeping the same variable name.
    println!("decimal = {}", decimal);

    let short_decimal = 128i16; // using suffix to declare type
    println!("short_decimal = {}", short_decimal);

    // In Rust, numbers can use underscores _ within numeric literals for
    // better readability by separating digits into groups. This is particularly
    // useful for large numbers, making it easier to parse and understand the
    // value at a glance. The underscores are ignored by the compiler and do not
    // affect the value itself.
    let a_million = 1_000_000;
    println!("a_million = {}", a_million);

    // Rust has two primitive compound types: tuples and arrays.
    // A tuple is a collection of values of different types.
    // Tuples are constructed using parentheses (), and each tuple itself is a value with type signature (T1, T2, ...), where T1, T2 are the types of its members.
    // Functions can use tuples to return multiple values, as tuples can hold any number of values.
    // Here is a function that returns a tuple.
    fn reverse(pair: (i32, bool)) -> (bool, i32) {
        // `let` can be used to bind the members of a tuple to variables
        let (integer, boolean) = pair;
    
        (boolean, integer)
    }
    // :? is used to print something debug-style
    // couldn't print tuple with default format, so used debug format
    // It tells Rust to use the std::fmt::Debug trait implementation for
    // the value being printed. The Debug trait provides a default way to
    // format values in a debug-oriented representation.
    println!("reverse((1, true)) = {:?}", reverse((1, true)));
    
    // To create one, we can use parenthesis to group together values.
    let long_tuple = (1u8, 2u16, 3u32, 4u64,
                    -1i8, -2i16, -3i32, -4i64,
                    0.1f32, 0.2f64,
                    'a', true);
    
    // Values can be extracted from the tuple using tuple indexing
    println!("long tuple first value: {}", long_tuple.0);
    println!("long tuple second value: {}", long_tuple.1);
    
    // Tuples can be tuple members
    let tuple_of_tuples = ((1u8, 2u16, 2u32), (4u64, -1i8), -2i16);
    
    // Tuples are printable
    println!("tuple of tuples: {:?}", tuple_of_tuples);
    
    // But long Tuples cannot be printed
    // let too_long_tuple = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
    //                     11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
    //                     21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
    //                     31, 32, 33, 34, 35, 36, 37, 38, 39, 40);
    // println!("too long tuple");

    let x = 1.0;
    // x = 2.0; // error: cannot assign twice to immutable variable `x`
    println!("x = {}, typeof(x)={}", x, type_name::<f64>());
    
    let mut mut_x = 1;
    println!("mut_x = {}, typeof(mut_x)={}", mut_x, type_name::<i32>());
    mut_x = 2; // this is fine
    println!("mut_x = {}, typeof(mut_x)={}", mut_x, type_name::<i32>());

    mut_x = 4294967294i64;
    println!("mut_x = {}, typeof(mut_x)={}", mut_x, type_name::<i64>());
}