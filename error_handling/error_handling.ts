function divide(x: number, y: number): [number, string] {
  if (y === 0) {
    return [0, "division by zero"];
    // or
    // throw new Error("")
  }
  return [x / y, "division by zero"];
}

function main() {
  let [result, err] = divide(10, 2);
  if (err) {
    console.log("Error:", err);
  } else {
    console.log("Result:", result);
  }

  [result, err] = divide(10, 0);
  if (err) {
    console.log("Error:", err);
  } else {
    console.log("Result:", result);
  }
}

// !INFO: There are other ways to implement the error in ts like "throw new Error()".

main();
