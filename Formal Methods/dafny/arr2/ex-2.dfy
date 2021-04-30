method FindMaxIndex(a: array<int>) returns (index:nat)
requires a.Length > 0
ensures 0 <= index < a.Length
ensures forall k :: 0 <= k < a.Length ==> a[k] <= a[index]
{
    index := 0;
    var i := 0;

    while i < a.Length
    decreases a.Length - i
    invariant 0 <= i <= a.Length
    invariant 0 <= index < a.Length
    invariant forall j :: 0 <= j < i ==> a[j] <= a[index ]
    {
        if a[i] > a[index]
        {
            index := i;
        }
        i := i + 1;

    }
}


method FindMaxValue(a :array<int>) returns (val: int)
requires a.Length > 0
ensures forall i :: 0 <= i < a.Length ==> a[i] <= val
ensures exists i :: 0 <= i < a.Length && a[i] == val
{
    val := a[0];
    var i := 0;
    while i < a.Length
    decreases a.Length - i
    invariant 0 <= i <= a.Length
    invariant forall j :: 0 <= j < i ==> a[j] <= val
    invariant exists j :: 0 <= j < a.Length && a[j] == val
    {
        if a[i] > val
        {
            val := a[i];
        }
        i := i + 1;
    }
}