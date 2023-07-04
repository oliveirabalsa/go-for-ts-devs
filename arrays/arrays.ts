function main() {
  // Array declaration and initialization
  const numbers: number[] = [];
  numbers[0] = 1;
  numbers[1] = 2;
  numbers[2] = 3;
  numbers[3] = 4;
  numbers[4] = 5;
  console.log(numbers);

  // Array literal
  const names: string[] = ["Alice", "Bob", "Charlie"];
  console.log(names);

  // Array length
  console.log("Length:", numbers.length);

  // Iterating over an array
  for (let i = 0; i < numbers.length; i++) {
    console.log(numbers[i]);
  }
}

main();
