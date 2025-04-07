function main() {
  const num: number = -8.7;
  const base = 2;
  const exp = 3;

  // Absolute value
  const abs = Math.abs(num);
  console.log("Abs:", abs);

  // Rounding
  const round = Math.round(num);
  console.log("Round:", round);

  // Truncate
  const trunc = Math.trunc(num);
  console.log("Trunc:", trunc);

  // Ceiling
  const ceil = Math.ceil(num);
  console.log("Ceil:", ceil);

  // Floor
  const floor = Math.floor(num);
  console.log("Floor:", floor);

  // Square root
  const sqrt = Math.sqrt(64);
  console.log("Sqrt(64):", sqrt);

  // Power
  const power = Math.pow(base, exp);
  console.log("Pow(2, 3):", power);

  // Random number
  const random = Math.random();
  console.log("Random (0~1):", random);
}

main();
