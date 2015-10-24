main = print $ last $ take 10001 primes

primes :: [Integer]
primes = nums [2..]
  where nums (p:ps) = p : nums [x | x <- ps, x `mod` p > 0]
