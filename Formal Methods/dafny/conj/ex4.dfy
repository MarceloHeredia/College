method Main()
{
    var sequencia := [1,0,1];
    var multiconjunto_s := multiset(sequencia);

    assert multiconjunto_s == multiset{0,1,1};

    var conj := {1,0,1};
    var multi_c := multiset(conj);
    assert multi_c == multiset{0,1};
    
}