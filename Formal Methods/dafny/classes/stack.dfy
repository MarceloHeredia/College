class {:autocontracts} Stack
{
    ghost const maxLen: int;
    ghost var data: seq<int>;

    var arr: array<int>;
    var max: int;
    var head: int; //position

    predicate Valid()
    {
        max > 0 &&
        arr.Length == max &&
        maxLen == max &&
        0 <= head <= max &&
        data == arr[0..head]
    }

    constructor(n: int)
    requires n > 0
    ensures maxLen == n
    ensures data == []
    {
        max := n;
        arr := new int[n];
        head := 0;

        //spec
        maxLen := n;
        data := [];
    }

    method Push(e: int) returns (b: bool)
    ensures old(|data|) < maxLen ==> data == old(data) + [e]
    ensures old(|data|) >= maxLen ==> data == old(data)
    ensures b == (old(|data|) < maxLen)
    ensures b == (|data| != old(|data|))
    {
        if head < max
        {
            arr[head] := e;
            head := head + 1;

            //spec
            data := data + [e];

            return true;
        }
        return false;
    }

    method Pop() returns (e: int)
    requires |data|>0
    ensures data == old(data)[0..old(|data|)-1]
    ensures e == old(data)[|data|]
    {
        head := head - 1;
        e := arr[head];

        //spec
        data := arr[0..head];
    }

    method Drop()
    ensures old(|data|)>0 ==> data == old(data)[0..old(|data|)-1]
    {
        if head > 0
        {
            head := head -1;

            //spec
            data := arr[0..head];
        }
    }

    method Top() returns (e: int)
    requires |data| > 0
    ensures data == old(data)
    ensures e == data[|data|-1]
    {
        return arr[head-1];
    }

    method Full() returns (b: bool)
    ensures b == (|data| == maxLen)
    ensures data == old(data)
    {
        return head == max;
    }

    method Empty() returns (b: bool)
    ensures b == (|data| == 0)
    ensures data == old(data)
    {
        return head==0;
    }

    method Count() returns (n: int)
    ensures n == |data|
    ensures data == old(data)
    {
        n := head;
    }

    method Capacity() returns (s: int)
    ensures s == maxLen
    ensures data == old(data)
    {
        s := max;
    }

    method Reverse()
    requires |data| > 0
    ensures |data| == old(|data|)
    ensures reversed(data, old(data))
	ensures forall k :: 0 <= k < head ==> arr[k] == old(arr[head-1 - k]);
    {

        var i := 0;
        var j := head-1;
        while i < (j+1)/2
        invariant 0 <= i <= (j+1)/2
        invariant 0 <= j < head <= arr.Length;
		invariant forall k :: 0 <= k < i || j-i < k <= j ==> arr[k] == old(arr[j-k]);
		invariant forall k :: i <= k <= j-i ==> arr[k] == old(arr[k]);
        {
            arr[i], arr[j-i] := arr[j-i], arr[i];
            i := i + 1;
        }

        data := arr[0..head];
        
    }
}
predicate permutacao (a:seq<int>, b:seq<int>)
{
    multiset(a) == multiset(b)
}
predicate reversed(sq1: seq<int>, sq2: seq<int>)
requires |sq1| == |sq2|
{
    forall i :: 0 <= i < |sq1| ==> sq1[i] == sq2[|sq2|-1-i]
}

method Main()
{
    var s := []; 
    var t := [];
    assert reversed(s,t); //validate predicate for empty 
    s := [1,2];
    t := [2,1];
    assert reversed(s,t); //validate predicate for not empty

    var stack := new Stack(4);
    var cap := stack.Capacity();
    assert cap == 4; 

    var empt := stack.Empty();
    assert empt; 

    var full := stack.Full();
    assert !full;

    var push := stack.Push(1);
    assert push == true;
    push := stack.Push(2);
    assert push == true;
    push := stack.Push(3);
    assert push == true;
    push := stack.Push(4);
    assert push == true; //max = 4

    push := stack.Push(5);
    assert push == false; //should not push anymore

    empt := stack.Empty();
    assert !empt;

    full := stack.Full();
    assert full;

    var count := stack.Count();
    assert count == 4;

    var top := stack.Top();
    assert top == 4;
    assert stack.data == [1,2,3,4];
    
    stack.Reverse();
    assert stack.data == [4,3,2,1];

    stack.Drop();
    assert stack.data == [4,3,2];

    var pop := stack.Pop();
    assert pop == 2;
    
}