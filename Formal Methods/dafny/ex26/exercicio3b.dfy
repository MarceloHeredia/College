
method Main()
{
    var a := new int[3];
    a[0], a[1], a[2] := 1,2,3;
    var b := new int[3];
    b[0], b[1], b[2] := 3,1,2;

    assert a[..] == [1,2,3];
    assert b[..] == [3,1,2];
    assert permutacao(a[..], b[..]);
}

method swap(a: array<int>, i:int, j:int)
requires 0 <= i < j < a.Length
modifies a
ensures permutacao(a[..], old(a[..]))
ensures a[i] == old(a[j]) && a[j] == old(a[i])
ensures forall k :: 0 <= k < a.Length && k !in {i,j} ==> a[k] == old(a[k])
{
    a[i], a[j] := a[j], a[i];
}

method bubbleSort(a:array<int>)
ensures ordenado(a)
ensures permutacao(a[..], old(a[..]))
modifies a
{
    var i := 0;
    while i < a.Length
    decreases a.Length - i;
    invariant 0 <= i <= a.Length
    {
        var j := 0;
        bubbleLoop(a, i);
        i := i + 1;
    }
}

method bubbleLoop(a: array<int>, i:int)
requires 0 <= i < a.Length
ensures permutacao(a[..], old(a[..]))
modifies a
{
    var j := 0;
    while j < a.Length - i - 1
    decreases a.Length -i -1 -j
    {
        if a[j] > a[j+1]
        {
            swap(a, j, j+1);
        }
        j := j + 1;
    }
}

predicate permutacao(a:seq<int>, b:seq<int>)
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