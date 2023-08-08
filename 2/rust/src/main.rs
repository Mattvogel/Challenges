fn main() {
    let mut a: i64 = 1;
    let mut b: i64 = 1;
    let mut i: i64 = 0;
    let mut evens: Vec<i64> = Vec::new();
    while i < 4000000 {
      i = a + b;
      a = b;
      b = i;
      if i % 2 == 0 {
        evens.push(i);
      }
    }
    let sum: i64 = evens.iter().sum();
    
    println!("Sum: {}", sum);
}
