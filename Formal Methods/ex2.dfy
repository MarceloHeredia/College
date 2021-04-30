method Main()
{
    var a := new int [5];

    var i := 0;
    while i < a.Length
    decreases a.Length - i
    {
        a[i] := 0;
        i := i + 1;
    }
}