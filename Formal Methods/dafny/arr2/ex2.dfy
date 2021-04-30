method Main()
{
    var a := new int[5];
    var i := 0;
    while i < a.Length
    decreases a.Length - i;
    invariant 0 <= i <= a.Length
    invariant forall j :: 0 <= j < i ==> a[j] == 0
    {
        a[i] := 0;
        i := i + 1;
    }
    var v := a[3];
    assert v == 0;
}