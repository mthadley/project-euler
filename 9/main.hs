main = print $ pTriplet 1000

pTriplet :: Int -> Int
pTriplet x = head [
    a * b * c |
    a <- range, b <- range,
    let c = x - a - b, a + b + c == x, a^2 + b^2 == c^2
  ]
  where range = [1..x]
