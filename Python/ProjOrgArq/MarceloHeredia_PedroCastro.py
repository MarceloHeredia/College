#Marcelo Heredia e Pedro Castro

import sys
import os
import lib.decodificador
import lib.codificador


def main():
    cdirectory = os.getcwd()#coletando a pasta atual no sistema
    #testando se operacao tem os parametros necessarios
    if len(sys.argv)-1 < 2:
        print('Instrucao de chamada: '
                '\n<python mipstranslator.py operacao arquivo> onde: \n'
                "operacao = -c (codificar) ou -d (decodificar)\n"
                "arquivo = nome do arquivo com extensao ex: file.asm")
        return

    if sys.argv[1] != '-d' and sys.argv[1] != '-c':
        print('Instrucao de operacao invalida:', sys.argv[1])
        print('\n\nInstrucao de chamada: '
            '\n<python mipstranslator.py operacao arquivo> onde: \n'
            'operacao = -c (codificar) ou -d (decodificar)\n'
            'arquivo = nome do arquivo com extensao ex: file.asm')
    path = cdirectory+"/in/"+sys.argv[2] #caminho completo ate o arquivo
    try:
        f = open(path,'r')

        if sys.argv[1] == '-c':
            text = f.readlines()
            lib.codificador.encode(text, sys.argv[2])#chama classe que executa codificacao
        
        elif sys.argv[1] == '-d':
            text = f.readlines() #todo conteudo do arquivo
            lib.decodificador.decode(text, sys.argv[2]) #chamada de classe que executa a decodificacao
    except Exception as e:
        print(e)


if __name__ == '__main__':
    main()



