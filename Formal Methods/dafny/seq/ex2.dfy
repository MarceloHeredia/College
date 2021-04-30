method Main()
{
    var a := new int[5];
    a[0] := 0;
    a[1] := 1;
    a[2] := 2;
    a[3] := 3;
    a[4] := 4;

    var s := a[..];

    assert s == [0,1,2,3,4];
    assert |s| == a.Length;

    assert a[1..3] == [1,2];

    assert 10 !in a[..];

    //assert 3 in a[..];
    assert a[3] == 3;
    assert exists i :: 0 <= i < a.Length && a[i] == 3;
}