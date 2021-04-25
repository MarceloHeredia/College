--Marcelo Heredia
import Data.Char (toLower, isAlpha)
import System.IO

removeNonAlpha :: String -> String
removeNonAlpha = map(toLower) . filter (isAlpha)

isPalindrome :: String -> Bool
isPalindrome xs = xs == reverse xs

palindromo :: IO ()
palindromo = do putStrLn "Digite uma cadeia de caracteres:"
                word <- getLine
                if (isPalindrome (removeNonAlpha word)) then
                    do putStrLn "É palíndromo."
                else
                    do putStrLn "Não é palíndromo."

--- QUESTÃO 2
type CIN = String

getDigit :: Char -> Int
getDigit c = read [c]

fixLen :: String -> String
fixLen xs
  | length xs == 1 = "0" ++ xs
  | otherwise = xs

getSum :: CIN -> CIN
getSum = fixLen . show . sum . map (getDigit)

addSum :: CIN -> CIN
addSum xs = xs ++ getSum xs

validar :: CIN -> Bool
validar xs = getSum (take 8 xs) == drop 8 xs


--- QUESTÃO 3

getInt :: IO Int
getInt = do xs <- getLine
            return (read xs :: Int)

somaAux :: Int -> IO Int
somaAux num = do n <- getInt
                 if num == 1 then
                   return n
                 else do res <- somaAux (num -1)
                         return (n + res)

somador :: IO ()
somador = do putStr "Quantos números? "
             n <- getInt
             res <- somaAux n
             putStr "O total é "
             putStrLn (show res)
