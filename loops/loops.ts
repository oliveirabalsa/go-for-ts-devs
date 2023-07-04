function main() {
  // For loop
  for (let i = 0; i < 5; i++) {
    console.log(i);
  }

  // While loop
  let j = 0;
  while (j < 5) {
    console.log(j);
    j++;
  }

  // Infinite loop
  let k = 0;
  while (true) {
    console.log(k);
    k++;
    if (k === 5) {
      break;
    }
  }

  // For...of loop
  const numbers = [1, 2, 3, 4, 5];
  for (const value of numbers) {
    console.log("Value:", value);
  }

  // forEach loop
  numbers.forEach((value, index) => {
    console.log("Value:", value, "Index:", index);
  });
}

main();
