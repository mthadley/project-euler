main = print $ ((^2) $ sum nums) - (sum $ map (^2) nums)
  where nums = [1..100]
