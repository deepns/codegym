# Refreshing C++

Been a long time did any C++ stuff. Refreshing my rusted memory to unlearn my old C++ knowledge and learn the current standards.

## Topics to learn

Growing list of topics to learn

- Namespaces
- Classes
  - Constructors
  - Inheritance
  - Abstract base classes
- Templates
- Advancements from C++11 and beyond
- Data structures from the standard library
  - list
  - vector
  - set
  - map
  - stack
  - queue
- smart pointers

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