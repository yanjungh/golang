# Pointers

## Difference between []*User and *[]User

The rule is to read the type from left to right.  Thus,

   1. `[]*User` means a slice of pointers to the `User` struct.
   2. `*[]User` means a pointer to a slice of the `User` struct.

Sample code: pointer2.go
   
Reference: https://stackoverflow.com/questions/50659408/difference-between-users-and-users-in-golang