method Main()
{
    var ms0 : multiset<int> := multiset{};

    var ms1 := multiset{1,1,1,2};
    var ms2 := multiset{1,1};

    assert |ms1| == 4;
    assert ms1[1] == 3;

    assert 1 in ms1;
    assert 0 !in ms1;
    

    assert ms1 == multiset{1,2,1,1};

    assert ms1 != multiset{};

    assert multiset{1,2} <= ms1;

    assert ms1 + ms2 == multiset{1,1,1,1,1,2};
    assert ms1 * ms2 == multiset{1,1};

    assert ms1 - ms2 == multiset{1,2};

    assert ms1[2:=0] == multiset{1,1,1};
}