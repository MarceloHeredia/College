method Main()
{
    var s1: set<int> := {};
    var s2 := {1,2,3};

    assert {1,2} == {2,1};

    assert {1,1,2,2,2,3} == s2; //no repetitions

    assert s1 != s2;

    assert 1 in s2;

    assert 0 !in s2; // nao pertence

    assert |s1| == 0; //cardinalidade

    assert {} < {1,2}; //subconjunto

    assert {1,2} <= {1,2};

    assert {1,2} > {1};

    assert {1,2} >= {1,2};

    var s3 := {1,2};
    var s4 := {3,4};

    assert s3 + s4 == {1,2,3,4};
    assert s2 * s3 == {1,2};

    assert s2 - s3 == {3};
    assert s3 - s2 == {};

    assert s3 !! s4;
}