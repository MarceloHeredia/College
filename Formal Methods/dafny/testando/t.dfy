function Pot(x:nat, y:nat):nat
{
    if y==0
    then 1
    else x * Pot(x, y-1)
}