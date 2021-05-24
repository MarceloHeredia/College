method Main()
{
    var x:int, y:int;

    assume x > 2 && y > 3;

    assert y + x + 1 > 6;
    
    x := x+1;

    assert y+x > 6;

    y := y+x;

    assert y > 6;
}