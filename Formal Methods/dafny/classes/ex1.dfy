class Contador
{
    var valor: int;

    method Inc()
    modifies this
    ensures valor == old(valor) + 1
    {
        valor := valor + 1;
    }

    method Decr()
    modifies `valor
    ensures valor == old(valor) - 1
    {
        valor := valor - 1;
    }

    method GetValor() returns (v: int)
    ensures v == valor
    ensures valor == old(valor)
    {
        return valor;
    }

    constructor()
    ensures valor == 0
    {
        valor := 0;
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