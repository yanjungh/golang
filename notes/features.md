# Features

Go passes everything by value, meaning that go could be copying structs ro values. 
Therefore we want to use pointer in many cases, which is cheap, since only the pointer
gets copied.

# The `init` function

The `init()` function is one of the two predefined functions: `init()` and `main()`.  It's typically used to
initialize specific state of a program.

   * It can be used within a `package` block.
   * No matter how many times the package is imported, the `init` is called once.
   * It must have no arguments and no return values.