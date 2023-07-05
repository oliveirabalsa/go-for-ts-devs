function main() {
  // Declare a variable and assign a value
  let value: number = 42;

  // Declare a pointer variable and assign the address of the variable
  let pointer: number = value;

  // Access the value using the pointer
  console.log("Value:", pointer);

  // Modify the value using the pointer
  pointer = 24;
  console.log("Modified value:", value);
}

main();
