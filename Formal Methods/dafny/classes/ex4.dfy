class {:autocontracts} FilaNat
{
    ghost var Conteudo: seq<nat>;

    var a: array<nat>;
    var tail: nat;

    predicate Valid()
    {
        a.Length > 0
        && 0 <= tail < a.Length
        && Conteudo == a[0..tail]
        
    }

    constructor()
    ensures Conteudo == []
    {
        a := new nat[5];
        tail := 0;

        //specification
        Conteudo := [];
    }

    method Enfileirar(e: nat)
    ensures Conteudo == old(Conteudo) + [e]
    {
        if tail == a.Length
        {
            var novo := new nat[2*a.Length];
            forall i | 0 <= i < a.Length
            {
                novo[i] := a[i];
            }
            a := novo;
        }

        a[tail] := e;
        tail := tail + 1;

        Conteudo := Conteudo + [e];
    }

    method Desenfileirar() returns (e: nat)
    requires |Conteudo| > 0
    ensures e == old(Conteudo)[0]
    ensures Conteudo == old(Conteudo)[1..]
    {
        e := a[0];
        tail := tail -1;
        forall i | 0 <= i < tail
        {
            a[i] := a[i+1];
        }

        Conteudo := a[0..tail];

    }

    method Quantidade() returns (n: nat)
    ensures n == |Conteudo|
    ensures Conteudo == old(Conteudo)
    {
        n := tail;
    }

}

method Main()
{
    var fila := new FilaNat();
    fila.Enfileirar(1);
    fila.Enfileirar(2);
    assert fila.Conteudo == [1,2];
    var q := fila.Quantidade();
    assert q == 2;
    var e := fila.Desenfileirar();
    assert e == 1;
    assert fila.Conteudo == [2];


}