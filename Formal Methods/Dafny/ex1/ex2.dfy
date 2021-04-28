function Pot (x: nat, y: nat): nat
{
    if y == 0
    then 1
    else x * Pot(x, y-1)
}

method Potencia(x:nat, y:nat) returns (a: nat)
ensures a == Pot(x,y)
{
    var b := x;
    var p := y;
    a := 1;
    while p > 0
    invariant Pot(b, p)*a == Pot(x,y)
    {
        a := a * b;
        p := p - 1;
    }
}