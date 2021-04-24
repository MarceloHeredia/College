fn main()
{/*
    let mm14 = calc_probs(1,4,5.,10.);
    print_probs(& mm14);

    let mm24 = calc_probs(2,4,5.,10.);
    print_probs(&mm24);
*/
    //let chegadas = 1250f64;
    simulate(3,3,156.25,62.5)
}

fn calc_probs(sv: i32, cp: i32, ld: f64, u:f64) -> Vec<f64>
{
    let mut x = 1;

    // calc of proportions
    let mut proportions = Vec::new();
    proportions.push(1.0); //P0 = 1

    for i in 1..cp+1
    {
        let val = proportions[(i-1) as usize] * (ld / (x as f64 * u));

        proportions.push(val);
        //print!("{}, ", proportions[i as usize]);

        if x<sv
        {
            x+=1;
        }
    }

    let s_proportion:f64 = proportions.iter().sum();
    //println!("Proportion: {}", s_proportion);

    // calc of probabilities
    let mut probabilities = Vec::new();

    for i in 0..cp+1
    {
        let val = proportions[i as usize] / s_proportion;
        probabilities.push(val);
        //print!("{:.4}, ", probabilities[i as usize]);
    }
    //let sprob:f64 = probabilities.iter().sum();

    probabilities

}

fn print_probs(p: &Vec<f64>)
{
    let space = 15;
    println!("| {0: ^spc$} | {1: ^spc$} | {2: ^spc$} |", "Clientes", "Indice", "Probabilidade", spc = space);
    let mut i = 0;
    for el in p
    {
        println!("| {0: ^spc$} | {1:^spc$.4} | {2:^spc$.2} |", i, el, el*100.0, spc=space);
        i+=1;
    }
    let s:f64 = p.iter().sum();
    println!("| {0: ^spc$} | {1:^spc$} | {2:^spc$.2} |\n", "", s, s*100.0, spc=space);
}

fn calc_losses(p: &Vec<f64>, ld: &f64) -> f64
{
    p.last().unwrap() * ld
}

fn calc_pop(p: &Vec<f64>) -> Vec<f64>
{
    let mut i = 0;
    let mut pop = Vec::new();
    for el in p
    {
        pop.push(el * i as f64);
        i += 1;
    }
    pop
}

fn calc_flow(p: &Vec<f64>, u:f64) -> Vec<f64>
{
    let mut v = Vec::new();
    let mut i = 0;
    for el in p
    {
        v.push(el * (u * i as f64));
        i += 1;
    }
    v
}

fn calc_utilization(p: &Vec<f64>, servers: i32) -> Vec<f64>
{
    let mut ut = Vec::new();
    ut.push(0f64);
    for i in 1..p.len()
    {
        let ci = servers.min(i as i32);
        ut.push(p[i] * (ci as f64 / servers as f64));
    }
    ut
}


fn print_table(p: &Vec<f64>, N: &Vec<f64>, D: &Vec<f64>, U: &Vec<f64>, W: f64)
{
    let space = 18;
    println!("| {0:^spc$} | {0:^spc$} | {1:^spc$} | {2:^spc$} | {3:^spc$} | {4:^spc$} |", "", 'N', 'D', 'U', 'W', spc=space);
    println!("| {0:^spc$} | {1:^spc$} | {2:^spc$} | {3:^spc$} | {4:^spc$} | {5:^spc$} |",
        "Clientes", "p%", "População", "Vazão", "Utilização", "Tempo de Resposta", spc=space);

    for i in 0..p.len()
    {
        println!("| {0:^spc$.4} | {1:^spc$.4} | {2:^spc$.4} | {3:^spc$.4} | {4:^spc$.4} | {5:^spc$.4} |",
           i , p[i], N[i], D[i], U[i], (if i==0 {0f64} else {W}), spc=space);
    }
    let pct:f64 = p.iter().sum();
    let pt:f64 = N.iter().sum();
    let ft:f64 = D.iter().sum();
    let ut:f64 = U.iter().sum();
    println!("| {0:^spc$.4} | {1:^spc$.4} | {2:^spc$.4} | {3:^spc$.4} | {4:^spc$.4} | {5:^spc$.4} |",
           "Total" , pct, pt, ft, ut, W, spc=space);
    println!("| {0:^spc$} | {1:^spc$} | {2:^spc$} | {3:^spc$} | {4:^spc$} | {5:^spc$} |",
        "escala", "-", "requisitions", "req/h", "%", "h", spc=space);

}

fn simulate(servers:i32, capacity:i32, lambda:f64, u:f64)
{
    let prob = calc_probs(servers, capacity, lambda, u);
    print_probs(&prob);

    let N = calc_pop(&prob);
    let D = calc_flow(&prob, u);
    let U = calc_utilization(&prob, servers);

    let ns:f64 = N.iter().sum();
    let ds:f64 = D.iter().sum();
    let W = ns/ds;

    print_table(&prob, &N, &D, &U, W)

}