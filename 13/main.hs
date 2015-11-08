main = do
  nums <- readFile "numbers.txt"
  print $ (take 10 . show . sum . map readInteger . lines) nums

readInteger :: String -> Integer
readInteger = read
