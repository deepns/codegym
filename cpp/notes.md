# Refreshing C++

Been a long time did any C++ stuff. Refreshing my rusted memory to unlearn my old C++ knowledge and learn the current standards.

## Day 1

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