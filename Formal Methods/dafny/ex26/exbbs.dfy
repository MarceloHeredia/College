predicate permutacao (a:seq<int>, b:seq<int>)
{
    multiset(a) == multiset(b)
}
predicate ordenadoEntre(a:array<int>, e:int, d:int)
requires 0 <= e <= d <= a.Length
reads a
{
    forall i,j ::e <= i <= j < d ==> a[i] <= a[j]
}
predicate ordenado(a:array<int>)
reads a
{
    ordenadoEntre(a,0,a.Length)
}

method bubbleSort(a:array<int>)
ensures ordenado(a)
ensures permutacao(a[..], old(a[..]))
modifies a
{
    if a.Length > 1
    {
        var i := 1;
        while i < a.Length
        invariant 1 <= i <= a.Length
        invariant ordenadoEntre(a,0,i)
        invariant permutacao(a[..], old(a[..]))
        {
            bubbleStep(a,i);
            i := i + 1;
        }
    }
}
method bubbleStep(a: array<int>, i: int)
requires 0 <= i < a.Length
requires ordenadoEntre(a,0,i)
ensures ordenadoEntre(a,0,i+1)
ensures permutacao(a[..], old(a[..]))
modifies a
{
    var j := i;
    while j>0 && a[j-1] > a[j]
    invariant 0 <= j <= i
    invariant ordenadoEntre(a,0,j) && ordenadoEntre(a,j,i+1)
    invariant 1 < j + 1 <= i ==> a[j-1] <= a[j+1]
    invariant permutacao(a[..], old(a[..]))
    {
        a[j-1], a[j] := a[j], a[j-1];
        j := j - 1;
    }
}