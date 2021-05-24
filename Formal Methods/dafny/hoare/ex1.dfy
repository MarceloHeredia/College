method Main()
{
    var x:nat, y:nat;
    assume y<=4;
    x := y + 1;
    assert x <= 5;
}