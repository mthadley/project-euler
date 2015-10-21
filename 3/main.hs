import Data.List

num = 600851475143

main = print $ find primeFactor [start, start - 1 .. 0]
	where
		start = sqrt' num
		primeFactor x = num `mod` x == 0 && isPrime x

isPrime :: Integer -> Bool
isPrime x = not $ any ((== 0) . (mod x)) $ [2..end]
	where end = sqrt' x

sqrt' :: Integer -> Integer
sqrt' = ceiling . sqrt . fromIntegral
