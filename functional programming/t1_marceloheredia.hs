-- Marcelo Heredia
--1
insere :: Int -> [Int] -> [Int]
insere n [] = [n]
insere n (x:xs)
  | n > x = x:(insere n xs)
  | otherwise = n:x:xs
-- para um numero N e uma lista
-- se n for maior que um elemento x da lista
--    concatena x e chama recursiva N e o restante da lista
-- se n for menor ou igual a x
--    concatena N, x e o restante da lista
-- insere 3 [1,2,4,5]
-- [1] ++ insere 3 [2,4,5]
-- [1] ++ [2] + insere 3 [4,5]
-- [1] ++ [2] + [3] ++ [4,5]

--2
ordenaInsere :: [Int] -> [Int]
ordenaInsere [] = []
ordenaInsere (x:xs) = insere x (ordenaInsere  xs)

--se a lista estiver vazia, retorna a lista
--senao, chama insere x (ordena insere xs) ou seja, fará uma recursao chamando ordenaInsere
-- quando chegar na lista vazia, a recursão irá para fora e começara a ser chamado o insere x []
-- após isso, cada retorno da recursao tera um elemento a mais ordenado na lista a direita para chamada do insere
--ordenaInsere [3,2,1]
-- insere 3 (ordenaInsere [2,1])
-- insere 2 (ordenaInsere [1])
-- insere 1 (ordenaInsere []) -- retornara []
-- saida da recursao >  insere 1 [] = [1]
-- insere 2 [1] = [1,2]
-- insere 3 [1,2] = [1,2,3]

--3
uneOrdenado :: [Int] -> [Int] -> [Int]
uneOrdenado [] [] = []
uneOrdenado [] (y:ys) = y : uneOrdenado [] ys
uneOrdenado (x:xs) [] =  x : uneOrdenado xs []
uneOrdenado (x:xs) (y:ys)
  | x < y = x : uneOrdenado xs (y:ys)
  | otherwise = y : uneOrdenado (x:xs) ys
-- se ambas listas estiverem vazias, retorna lista vazias
-- se uma apenas estiver vazia, retorna o primeiro elemento da outra concatenado a chamada recursiva com o resto da listas
--caso ambas tenham elementos, compara os primeiros elementos de cada uma e retorna o menor concatenado a chamada recursiva
--uneOrdenado [1,2] [4]
-- 1 : uneOrdenado [2] [4]
-- 1 : 2 : uneOrdenado [] [4]
-- 1:2:4 : uneOrdenado [][]
-- 1:2:4:[]

--4
ordenaUne :: [Int] -> [Int]
ordenaUne [] = []
ordenaUne [x] = [x]
ordenaUne xs = uneOrdenado (ordenaUne esq) (ordenaUne dir)
                where (esq,dir) = halve xs
--criterios de parada sao caso a chamada recursiva contenha lista vazia ou apenas um elemento, como solicitado no enunciado
-- listas vazias ou um elemento ja sao considerados ordenados
--chamada recursiva fará divisoes sucessivas na metade do vetor, ate restarem apenas vetores de 0 ou 1 posicoes
--apos isso os vetores serao juntados de forma ordenada pelo metodo uneOrdenado
--ordenaUne [3,2,1]
--uneOrdenado (ordenaUne [3]) (ordenaUne[2,1])
--uneOrdenado (uneOrdenado [] [3]) (uneOrdenado ( ordenaUne [2]) (ordenaUne [1]))
--uneOrdenado ( [3] ) (uneOrdenado ([][2])([][1]))
--uneOrdenado ([3])([1,2])
-- [1,2,3]
halve :: [a] -> ([a],[a]) --funcao auxiliar vista em aula para fazer as divisoes sucessivas
halve xs = (take len xs, drop len xs)
            where len = length xs `div` 2

--5
--a funcao zipWith recebe uma funcao/operacao e duas listas de conteudo generico
--retorna uma lista com o resultado da operacao executada entre as duas listas
--exemplo
--zipWith (+) [2,2,2] [1,2,3]
-- (+) (x:xs) (y:ys) = (+) 2 1 : zipWith (+) [2,2] [2,3]
-- 3 : 4 : zipWith (+) [2] [3]
-- 3 : 4 : 5 : zipWith (+) [] []
-- 3 : 4 : 5 : []
-- [3,4,5]

cresc :: (Ord a) => [a] -> Bool
cresc xs = and ( zipWith (<=) xs (tail xs))

disjuntas :: (Ord a) => [a] -> [a] -> Bool
disjuntas [] _ = True
disjuntas (x:xs) ys
  | elem x ys = False
  | otherwise = disjuntas xs ys

-- para cada elemento da primeira lista, verifica se ele esta contido na segunda
-- se estiver, retorna Falso
-- se nao, chama recursivo testando os proximos da primeira lista, ate que a primeira lista esvazie sem achar correspondencias
-- [1,2,4] [2,3,5]
-- elem 1 [2,3,5]  -- false entra no otherwise
-- disjuntas [2,4] [2,3,5]
-- elem 2 [2,3,5] -- true Retorna False, nao eh disjunta
