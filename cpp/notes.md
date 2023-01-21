# Refreshing C++

Been a long time did any C++ stuff. Refreshing my rusted memory to unlearn my old C++ knowledge and learn the current standards.

## Topics to learn

Growing list of topics to learn

- Namespaces
- Classes
  - Constructors
  - Inheritance
  - Abstract base classes
- Virtual stuff
  - virtual classes
  - virtual functions
- Templates
- Advancements from C++11 and beyond
- Fom the standard library
  - [ ] array
  - [x] list
  - [ ] vector
  - [ ] set
  - [ ] map
  - [ ] stack
  - [ ] queue
  - [ ] function
  - [ ] tie
  - [ ] tuple
- [ ] smart pointers
- [ ] logging framework
- [ ] date and time manipulations
- [ ] constexpr
- [ ] lambdas
- [ ] initializer_list
- [ ] std::allocator
- [x] nullptr
- [ ] Design patterns
  - [ ] Factory
  - [ ] Observer
  - [ ] Command

## Day 24

- exploring list and maps
- list - remove, remove_if and erase (using ranges)
- map - just getting started. map iteration, insertion and capacity methods

## Day 23

- Reading [Command design pattern](https://refactoring.guru/design-patterns/command)

## Day 22

- List remove and erase
- Getting started with design patterns. exploring observer pattern

## Day 21

- Didn't work

## Day 20

- little bit of std::reference_wrapper, std::shared_ptr

## Day 19

- explored more on std::tuple, std::tie, std::make_tuple, [structured bindings](https://riptutorial.com/cplusplus/example/3384/structured-bindings)

## Day 18

- added a concrete example to learn getopt/getopt_long
- explored std::tuple, std::tie, nullptr, std::stod along the way

## Day 17

- understanding getopt and getopt_long further..finally cleared the doubts on the optional argument

## Day 16

- exploring std::array

## Day 14-15

- didn't have much time to focus on this
- did little reading on lambda, and command line parsing
- getopt and getopt_long

## Day 13

- More on lambdas. Read some sections of this [page](https://www.learncpp.com/cpp-tutorial/introduction-to-lambdas-anonymous-functions/)

## Day 12

- vector operations and access
- Factory and AbstractFactory patterns
- Review of Google style guide
- brush up on vector operations

## Day 11

- More on simple classes
- Inheritance vs composition
- playing with std::vector

## Day 10

- typid, const_cast and simple classes
- getting started to play with std::list

## Day 9

- Hands on with explicit constructors, typeid and const_cast operators.
- found out the diff between using stdio and cstdio headers

## Day 8

- Some more hands on with constructor initialization and default constructors

## Day 7

- Hands on with some lambda and class constructor basics

## Day 6

- Reading about constructors from https://learn.microsoft.com/en-us/cpp/cpp/constructors-cpp

## Day 5

- Review of modern c++. Range based loops, auto keyword, lambda expressions, constexpr in place of macros, std::variant in place of unions.
- Things to try out more
  - Lambdas
  - Callable objects
  - Exception handling
  - Move constructor and assignment

## Day 4 - namespaces - continued

- Wrote a basic namespaces within .cpp files and using header files
- what next? - nested namespace, std namespace, using directive vs declaration, namespace alias

## Day 3 - Reading namespaces, watching OOP concepts on YT CodeBeauty

- Links to read
  - [ ] [Modern C++](https://learn.microsoft.com/en-us/cpp/cpp/welcome-back-to-cpp-modern-cpp)
  - [ ] [Namespaces](https://learn.microsoft.com/en-us/cpp/cpp/namespaces-cpp)

## Day 2 - Anonymous namespaces

- using anonymous namespaces to restrict global names to file scope (just like static)

## Day 1 - Hello World

- Set up vscode on mac
- have installed xcode before, came with clang compiler. `/usr/bin` shows multiple compilers, but all seem to point to only Apple's clang.

```console
➜  hello-world git:(dev) ✗ ls -l /usr/bin/clang* /usr/bin/g++ /usr/bin/c++
-rwxr-xr-x  76 root  wheel  167136 Oct 28 04:43 /usr/bin/c++
-rwxr-xr-x  76 root  wheel  167136 Oct 28 04:43 /usr/bin/clang
-rwxr-xr-x  76 root  wheel  167136 Oct 28 04:43 /usr/bin/clang++
-rwxr-xr-x  76 root  wheel  167136 Oct 28 04:43 /usr/bin/clangd
-rwxr-xr-x  76 root  wheel  167136 Oct 28 04:43 /usr/bin/g++

➜  hello-world git:(dev) ✗ c++ --version
Apple clang version 14.0.0 (clang-1400.0.29.202)
Target: arm64-apple-darwin22.1.0
Thread model: posix
InstalledDir: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin
➜  hello-world git:(dev) ✗ clang --version
Apple clang version 14.0.0 (clang-1400.0.29.202)
Target: arm64-apple-darwin22.1.0
Thread model: posix
InstalledDir: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin
➜  hello-world git:(dev) ✗ clang++ --version
Apple clang version 14.0.0 (clang-1400.0.29.202)
Target: arm64-apple-darwin22.1.0
Thread model: posix
InstalledDir: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin
```

- Configured C/C++ options, set up build tasks (`Cmd+Shift+B` to trigger build), .gitignore etc.