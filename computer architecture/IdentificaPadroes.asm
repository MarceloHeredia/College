# Marcelo Heredia
# Pedro Castro

#este codigo identifica padroes numericos em uma string e diz quantas vezes esse padrao ocorre


        .text                   # Add what follows to the text segment of the program
        .globl  main            # Declare the label main to be a global one
main:		
		la	$a0,vetorDados 	# $a0 contem o endereço de memoria do vetorDados (parametro da funcao)
		addu	$s3,$zero,$a0	# $s3 salva o endereço inicial de vetorDados no MAIN
		jal	carregaVetor	# chama o metodo carregaVetor 
		xor	$s0,$s0,$s0	#garante que $s0 está zerado
		addu	$s0,$zero,$v0	#$s0 contem o tamanho de vetorDados ($v0 = retorno da funcao)
		
		la	$a0,vetorPadrao	# $a0 contem o endereco de memodia de vetorPadrao (parametro da funcao)
		 addu	$s4,$zero,$a0	# $s4 salva o endereço inicial de vetorPadrao no Main
		jal	carregaVetor	#chama o metodo carregaVetor
		xor	$s1,$s1,$s1	#garante q $s1 esta zerado
		addu	$s1,$zero,$v0	#$s1 contem o tamanho de vetorPadrao ($v0 = retorno da funcao)
		
		li	$t8,0		#$t8 contem a contabilizacao de padrão (começa em 0)
		li	$t9,0		#$t9 contem a posicao atual dos dados
		
loopMain:	addu	$t5,$t9,$s1	#$t5 contem a soma da posicao atual com o tamanho do vetor de padroes
		bgt	$t5,$s0,imprimeFim #se a soma for maior q o tamanho do vetor de dados, não tem como ter um padrão, fim
		
		addiu	$sp,$sp,-24	#pilha recebe tamanho 24 (5x4bits) -> para mandar os parametros para o método + 4 do $ra
		sw	$s3,0($sp)	#salva a posicao de memoria do vetorDados na 1 posicao da pilha
		addiu	$sp,$sp,4	#muda a posicao atual da pilha
		sw	$t9,0($sp)	#salva 	POSICAO de vetor dados dentro da pilha
		addiu	$sp,$sp,4	 #muda posicao atual da pilha
		sw	$s4,0($sp)	#salva a posicao de memoria de vetorPadrao
		addiu	$sp,$sp,4	#muda posicao atual da pilha
		xor	$t5,$t5,$t5	#zera $t5 para usa-lo como 0 para o chamamento da funcao
		sw	$t5,0($sp)	#salva a posicao de padrao (atualmente é zero)
		addiu	$sp,$sp,4	#muda a pilha para a ultima posicao reservada de PARAMETROS
		sw	$s1,0($sp)	#salva o tamanho do padrao
		addiu	$sp,$sp,4	#a pilha esta na posicao onde sera guardado retorno
		
		jal	encontraPadrao
		
		addu	$t8,$t8,$v0	#$t8 incrementa o valor do retorno em v0
		addiu	$t9,$t9,1	#incrementa a posicao atual dos dados
		
		j	loopMain	#volta para o loop
		
		
imprimeFim:		#impressao final
		la    	$a0,OUTF        
		li    	$v0,4        	# Imprime "Digite o o. numero: "
  		syscall   
		
		add	$a0,$t8,$zero
		li	$v0,1		#imprimir inteiro
		syscall
		
		li	$v0,10		#saida
		syscall
		
		
carregaVetor:	
		move 	$t0,$a0		#move o parametro da funcao para $t0
		#define tamanho do vetor
		la    	$a0,STAM        
		li    	$v0,4        	# Imprime "Digite o o. numero: "
  		syscall   
  		
		li    	$v0,5          
		syscall            	# Leitura de um inteiro. Lembre-se que o valor digitado no teclado será armazenado no registrador $v0 (olhar tabela de códigos syscall) 
  		move  	$t1, $v0     	# Move o valor armazenado em $v0 para $t1
  		
  		li	$t2,0		#$t2 contem a posicao atual do vetor(começa em 0)
  		
leitura:	bge	$t2,$t1,returnCV #se a posicao>= tamanho para o laço (while pos < tam)
		
		la	$a0,RVAL	#carrega o endereço da palavra a ser printada
		li	$v0,4		#Imprime
		syscall
		
		li	$v0,5
		syscall			#leitura de inteiro
		move	$t3,$v0		#move o numero digitado pelo usuario para $t3
		
		mulu	$t4,$t2,4	#multiplica a posicao por 4 (inteiros tem 4 bits) para acesso (na primeira exec vai da 0)
		addu	$t7,$t0,$t4	#adiciona a posicao à referencia atual do vetor
  		sw	$t3,0($t7)	#salva o inteiro na posicao atual do vetor
  		addiu	$t2,$t2,1	#incrementa 1 em posicao
  		j	leitura		#pula para leitura (rotaciona o laço)
		
		#no momento que a posicao ficar igual ou maior que o tamanho, ele pula para leitura e dps pula para returnCV
				
returnCV:	#retorno da funcao carregaVetor
		move	$v0,$t1		#coloca o tamanho do vetor em $v0 para retorno
		jr	$ra		#retorna
		
		
encontraPadrao:	
		sw	$ra,0($sp)
		addiu	$sp,$sp,4	
		
aux:		addiu	$sp,$sp,-24	#volta a pilha para a posicao necessaria
		lw	$t0,0($sp)	#t0 contem o endereço do vetorDados
		addiu	$sp,$sp,4	#muda a posicao da pilha para proxima
		lw	$t1,0($sp)	#t1 contem a posicao
		mulu	$t2,$t1,4	#t2 contem a posicao * 4
		addu	$t0,$t0,$t2	#t0 esta na posicao correta do vetor
		
		addiu	$sp,$sp,4	#muda posicao da pilha	
		lw	$t3,0($sp)	#t3 contem endereço do vetorPadrao
		addiu	$sp,$sp,4	#muda posicao
		lw	$t4,0($sp)	#t4 contem posicao de vetorPadrao
		mulu	$t2,$t4,4	#t2 contem posicao * 4
		addu	$t3,$t3,$t2
		
		lw	$t0,0($t0)	#t0 é agora o valor armazenado na posicao do vetor
		lw	$t3,0($t3)	#t3 é agora o valor...
		
		bne	$t0,$t3,notequal	#se n for igual vai retornar 0
		#senao
		addiu	$sp,$sp,4	#pilha agora esta em tamanhoPadrao
		lw	$t0,0($sp)	#to contem agora o tamanho do padrao
		addiu	$t0,$t0,-1	#t0 contem agora o tamanho -1
		beq	$t4,$t0,equal	#se posicao vetpadrao == tamanhoPadrao-1 é igual
		#senao
		
		addiu	$sp,$sp,-12	#sp volta para a posicao 20 (ela estava em 8 ate entao) e vetorDados n sofrera modificacao
		addiu	$t1,$t1,1	#_posDads + 1
		sw	$t1,0($sp)	#armazena de volta na pilha
		addiu	$sp,$sp,8	#sp pula para posPadrao
		addiu	$t4,$t4,1	#t4 posPadrao+1
		sw	$t4,0($sp)
		addiu	$sp,$sp,12	#sp na posicao inicial
		j	aux

equal:
		addiu	$sp,$sp,4	#sp agora esta no endereço de retorno
		lw	$ra,0($sp)	#ra esta no endereco de retorno
		addiu	$v0,$zero,1	#v0 é 1
		addiu	$sp,$sp,4	
		jr	$ra		#retorna

notequal:	xor	$v0,$v0,$v0	#v0 é zero
		addiu	$sp,$sp,8	#sp contem o endereco de retorno
		lw	$ra,0($sp)	#ra contem o endereco de retorno
		addiu	$sp,$sp,4	
		jr	$ra		#retorna
		
.data
STAM: 		.asciiz "Informe o número de dados a serem inseridos no vetor: "
RVAL: 		.asciiz "Informe um dados a ser inserido no vetorDados: "
OUTF: 		.asciiz "Quantidade de padrões contabilizados: "
vetorDados:     .word   0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0
vetorPadrao:	.word	0x0 0x0 0x0 0x0 0x0 0x0
