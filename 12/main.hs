main = print $ head [x | x <- triangleNums, divisors x > 500]
  where
    divisors x = 1 + 2 * length [y | y <- [2 .. ceiling $ sqrt $ fromIntegral x], x `mod` y == 0]
    triangleNums = scanl1 (+) [1..]
