# Marcelo Heredia
# Pedro Castro

        .text                   # Add what follows to the text segment of the program
        .globl  main            # Declare the label main to be a global one
main:		
		addiu 	$sp,$sp,-16	# liberadas 5 posicoes na pilha
		la 	$t0, A		# t0 contem o end de memoria da primeira posicao de A
		sw	$t0, 0($sp)	#salva o vetor na pilha
		addiu	$sp,$sp,4	#anda uma posicao na pilha
		la	$t0,Prim	# t0 contem o end memoria do Prim
		lw	$t0,0($t0)	#t0 contem o valor numerico de Prim
		sw	$t0,0($sp)	#salva o valor na pilha
		addiu	$sp,$sp,4	#anda uma posicao
		la	$t0,Ult		#t0 contem o end memo de Ult
		lw	$t0,0($t0)	#t0 contem o valor de Ult
		sw	$t0,0($sp)	#salva na pilha
		addiu	$sp,$sp,4	#anda uma posicao
		la	$t0,Valor	#t0 contem o end memo de Valor
		lw	$t0,0($t0)	#t0 contme o valor de Valor
		sw	$t0,0($sp)	#salva o valor na pilha
		addiu	$sp,$sp,4	#anda uma posicao na pilha
		jal	binSrch
		
		addu 	$a0,$zero,$v0 	#print do retorno da funcao
		addiu 	$v0,$zero,1
		syscall
		
		li	$v0,10		#saida
		syscall
		
binSrch:	
		addiu 	$sp,$sp,-16
		lw	$t1,0($sp)	#t1 contem endereço d vetor
		addiu	$sp,$sp,4	#anda uma posica
		lw	$t2,0($sp)	#t2 contem valor Prim
		addiu	$sp,$sp,4	
		lw	$t3,0($sp)	#t3 contem Ult
		addiu	$sp,$sp,4	
		lw	$t4,0($sp)	#t4 contem Valor
		addiu	$sp,$sp,4	#pilha na posicao inicial
		
		blt	$t3,$t2,invalido# IF  Ult<Prim retornar -1	
		#ELSE
		addu	$t8,$t2,$t3	#t8 contem a soma de Prim+Ult
		div	$t8,$t8,2	#t8 contem Prim+Ult/2 T8 <-> MEIO
		mulu	$t9,$t8,4	#t8 contem a posicao MEIO em formato 4x1
		addu	$t9,$t9,$t1	#t9 contem a posicao atual do vetor A[MEIO]
		lw	$s0,0($t9)	#s0 contem o valor da posicao atual do vetor
		
		beq	$t4,$s0,encontrado	#if Valor == A[MEIO] return meio
		#else 
		
		blt	$t4,$s0,menor	# IF valor < A[MEIO] 
		
		#else
		# VALOR > A[MEIO]
		addiu 	$sp,$sp,-16	#recolocar os atributos da função para chamada recursiva
		sw	$t1,0($sp)	#colocando A na pilha
		addiu	$sp,$sp,4	#anda uma posica
		
		addiu	$t8,$t8,1	# devemos colocar ao inves de Prim, Meio+1 (t8 continha meio)
		sw	$t8,0($sp)	#coloca MEIO+1 no lugar de PRIM
		addiu	$sp,$sp,4	
		sw	$t3,0($sp)	#coloca Ult no lugar de Ult
		addiu	$sp,$sp,4	
		sw	$t4,0($sp)	#coloca o valor procurado novamente no lugar de Valor
		addiu	$sp,$sp,4	#pilha na posicao inicial
		
		j	binSrch
		

menor:			
		addiu 	$sp,$sp,-16	#recolocar os atributos da função para chamada recursiva
		sw	$t1,0($sp)	#colocando A na pilha
		addiu	$sp,$sp,4	#anda uma posica
		
		sw	$t2,0($sp)	#coloca Pim no lugar de PRIM
		addiu	$sp,$sp,4	
		
		addiu	$t8,$t8,-1	# devemos colocar ao inves de Ult, Meio-1 (t8 continha meio)
		sw	$t8,0($sp)	#coloca MEIO-1 no lugar de ULT
		addiu	$sp,$sp,4	
		sw	$t4,0($sp)	#coloca o valor procurado novamente no lugar de Valor
		addiu	$sp,$sp,4	#pilha na posicao inicial
		
		j	binSrch	
		
						
invalido:
		addiu $v0,$zero,-1
		j esvaziaPilha
		
encontrado:	
		addu  $v0,$zero,$t8
		j esvaziaPilha

esvaziaPilha:
		addiu $sp,$sp,-16
		sw    $zero,0($sp)
		sw    $zero,4($sp)
		sw    $zero,8($sp)
		sw    $zero,12($sp)
		addiu $sp,$sp,16
		jr $ra

.data
OUTF: 		.asciiz "Encontrado na posição: "
A:     .word   -5, -1, 5, 9, 12, 15, 21, 29, 31, 58, 250, 325
Prim:  .word	0
Ult:   .word	11
Valor: .word	17