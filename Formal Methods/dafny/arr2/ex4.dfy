predicate Sorted(a: array<int>)
reads a
{
    forall j, k :: 0 <= j < k < a.Length ==> a[j] < a[k]
}

method DoSomething(a: array<int>)
requires a.Length > 0
modifies a
{
    a[0] := 1;
}
