predicate NaoPertence(x:int, a:array<int>)
reads a
{
    forall i :: 0 <= i < a.Length ==> a[i] != x
}

predicate NaoPertence2(x:int, a:array<int>)
reads a
{
    x !in a[..]
}

function Sum(xs: seq<int>): int
{
    if |xs| == 0
    then 0
    else
    xs[0] + Sum(xs[1..])
}

method Somatorio(a: array<int>) returns (s : int)
ensures s == Sum(a[..])
{
    s := 0;
    var i := 0;
    while i < a.Length
    decreases a.Length - i
    invariant 0 <= i <= a.Length
    invariant s == Sum(a[a.Length-i..])
    {
        s := s + a[a.Length - i - 1];
        i := i + 1;
    }
}