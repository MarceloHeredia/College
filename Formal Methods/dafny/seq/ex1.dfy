method Main()
{
    var s := [1,2,3,4,5];
    var vazia: seq<int> := [];
    
    assert |s| == 5; //len
    assert |vazia| == 0;

    assert s[0] == 1; //index

    assert s[1..3] == [2,3]; //slice
    assert s[1..] == [2,3,4,5];
    assert s[..4] == [1,2,3,4];

    assert [1,2,3] + [2,3] == [1,2,3,2,3]; //concat
    assert s[2 := 6] == [1,2,6,4,5]; //alter

    assert 1 in s; //contains
    assert 0 !in s;

    assert vazia < s; //prefix
    assert [1] < s;
    assert [1,2,3,4,5] <= s;

    assert !([0,1] < s);

    assert forall i :: i in s ==> 1 <= i < 6;

    var t := [3.14, 2.7, 1.41, 1985.44, 100.0, 37.2][1:0:3];

    assert |t| == 3;
    assert t[0] == [3.14];
    assert t[1] == [];
    assert t[2] == [2.7, 1.41, 1985.44];

    assert t[0][0] == 3.14;
}