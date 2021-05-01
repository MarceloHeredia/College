method Main()
{
    assert (set x | x in {0,1,2,3,4,5} && x < 3) == {0,1,2};

    assert (set x | x in {0,1,2} :: x + 0) == {0,1,2};
    assert (set x:nat, y:nat | x < 2 && y < 2 :: (x,y)) == {(0,0), (0,1), (1,0), (1,1)};
    

    assert {0*1, 1*1, 2*1} == {0,1,2}; //auxiliar dafny
    assert (set x | x in {0,1,2} :: x * 1) == {0,1,2};
}