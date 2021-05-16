class Celula
{
    var dado:int;

    constructor()
    ensures dado == 0
    {
        dado := 0;
    }
}

class Contador
{
    //ghost
    ghost var valor: int;

    ghost var Repr: set<object>;

    var incs: Celula;
    var decs: Celula;

  
    predicate Valid()
    reads this, Repr
    {
        this in Repr && incs in Repr && decs in Repr
        &&
        incs != decs
        &&
        incs.dado >= 0
        &&
        decs.dado >= 0
        &&
        valor == incs.dado - decs.dado
    }

    constructor()
    ensures valor == 0
    ensures Valid()
    ensures fresh(Repr - {this})
    {
        incs := new Celula();
        decs := new Celula();
        valor := 0;
        Repr := {this, incs, decs};
    }

    method Inc()
    requires Valid()
    modifies this, incs
    ensures valor == old(valor) + 1
    ensures incs.dado == old(incs.dado) + 1
    ensures Valid()
    ensures fresh(Repr - old(Repr))
    {
        incs.dado := incs.dado + 1;
        valor := valor + 1;
    }

    method Decr()
    requires Valid()
    modifies this, decs
    ensures valor == old(valor) - 1
    ensures decs.dado == old(decs.dado) + 1
    ensures Valid()
    ensures fresh(Repr - old(Repr))
    {
        decs.dado := decs.dado + 1;
        valor := valor - 1;
    }


    method GetValor() returns (v: int)
    requires Valid()
    ensures Valid()
    ensures v == valor
    ensures valor == old(valor)
    ensures fresh(Repr - old(Repr))
    {
        v := incs.dado - decs.dado;
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