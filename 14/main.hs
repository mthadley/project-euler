import Data.Array
import Data.List
import Data.Function

main = print $ fst $ maximumBy (compare `on` snd) $ zip [1 .. max_num] $ map (length . c_seq) [1 .. max_num]

max_num :: Int
max_num = 1000000

c_seq :: Int -> [Int]
c_seq x = c_list ! x

c_list :: Array Int [Int]
c_list = listArray (1, max_num) $ map collatz [1 .. max_num]

collatz :: Int -> [Int]
collatz 1 = [1]
collatz x = x : collatz val
  where val = if even x then x `div` 2 else 3 * x + 1
