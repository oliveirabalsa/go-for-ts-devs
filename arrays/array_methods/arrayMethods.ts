function main() {
  // Array declaration and initialization
  let numbers: number[] = [1, 2, 3, 4, 5];
  console.log("Array:", numbers);

  // Push an element
  numbers.push(6);
  console.log("Pushed Array:", numbers);

  // Copy an array/slice
  const copiedNumbers: number[] = numbers.slice();
  console.log("Copied Array:", copiedNumbers);

  // Slice an array/slice
  const slicedNumbers: number[] = numbers.slice(1, 4); // [2, 3, 4]
  console.log("Sliced Array:", slicedNumbers);

  // Modify the slice
  slicedNumbers[0] = 99;
  console.log("Modified Slice:", slicedNumbers);
  console.log("Original Array:", numbers);

  // Flat and array
  const nested = [1, 2, [3, 4], [5, [6]]];
  const flat = nested.flat(2);
  console.log("Flattened:", flat);
}

main();
