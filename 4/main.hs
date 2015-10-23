main = print $ maximum palNums
  where
    nums = [100..999]
    palNums = [x * y | x <- nums, y <- nums, isPalNum $ x * y ]

isPalNum :: Int -> Bool
isPalNum x = y == reverse y
  where y = show x
