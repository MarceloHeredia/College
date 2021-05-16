class Contador
{
    //ghost
    ghost var valor: int;

    var incs: int;
    var decs: int;
  
    predicate Valid()
    reads this
    {
        incs >= 0
        &&
        decs >= 0
        &&
        valor == incs - decs
    }

    constructor()
    ensures valor == 0
    ensures Valid()
    {
        incs := 0;
        decs := 0;
        valor := 0;
    }


    method Inc()
    requires Valid()
    modifies this
    ensures valor == old(valor) + 1
    ensures Valid()
    {
        incs := incs + 1;
        valor := valor + 1;//abs ghhost
    }

    method Decr()
    requires Valid()
    modifies this
    ensures valor == old(valor) - 1
    ensures Valid()
    {
        decs := decs + 1;
        valor := valor - 1;//abs
    }


    method GetValor() returns (v: int)
    requires Valid()
    ensures Valid()
    ensures v == valor
    ensures valor == old(valor)
    {
        return incs - decs;
    }

}

method Main()
{
    var c := new Contador();
    var v := c.GetValor();
    assert v == 0;

    c.Inc();
    c.Inc();

    v:= c.GetValor();
    assert v == 2;

    c.Decr();
    v := c.GetValor();
    assert v == 1;
    
}