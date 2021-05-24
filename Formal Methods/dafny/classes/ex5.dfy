class {:autocontracts} FilaNatLimitada
{
    ghost const TamanhoMaximo: nat;
    ghost var Conteudo: seq<nat>;

    var a: array<nat>;
    var max: nat;
    var tail: nat;

    predicate Valid()
    {
        max > 0
        && a.Length == max
        && TamanhoMaximo == max
        && 0 <= tail <= max
        && Conteudo == a[0..tail]
    }

    constructor(n: nat)
    requires n > 0
    ensures TamanhoMaximo == n
    ensures Conteudo == []
    {
        max := n;
        a := new nat[n];
        tail := 0;

        //specification
        TamanhoMaximo := max;
        Conteudo := [];
    }

    method Enfileirar(e: nat)
    requires |Conteudo| < TamanhoMaximo
    ensures Conteudo == old(Conteudo) + [e]
    {
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

    method QuantidadeMaxima() returns (n: nat)
    ensures n == TamanhoMaximo
    ensures Conteudo == old(Conteudo)
    {
        return max;
    }

}

method Main()
{
    var fila := new FilaNatLimitada(5);
    fila.Enfileirar(1);
    fila.Enfileirar(2);
    assert fila.Conteudo == [1,2];
    var q := fila.Quantidade();
    assert q == 2;
    var e := fila.Desenfileirar();
    assert e == 1;
    assert fila.Conteudo == [2];


}