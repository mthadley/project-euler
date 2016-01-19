main = readFile "numbers.txt" >>= (print . take 10 . show . sum . map readInteger . lines)

readInteger :: String -> Integer
readInteger = read
