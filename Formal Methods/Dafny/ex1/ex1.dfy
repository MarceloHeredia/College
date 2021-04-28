predicate Par(x:nat)
{
  x % 2 == 0
}

function Fib(n: nat) : nat
{
  if n < 2
  then n
  else Fib(n-2) + Fib(n-1)
}

method Triplo(x: int) returns (r: int)
ensures r == 3 * x
{
  r := 3 * x;
}