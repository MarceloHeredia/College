#Marcelo Heredia e Pedro Castro

import os
import array
import sys

class Registradores:
    __regMap = {
        '$0': 0,
        '$1': 1,
        '$2': 2,
        '$3': 3,
        '$4': 4,
        '$5': 5,
        '$6': 6,
        '$7': 7,
        '$8': 8,
        '$9': 9,
        '$10': 10,
        '$11': 11,
        '$12': 12,
        '$13': 13,
        '$14': 14,
        '$15': 15,
        '$16': 16,
        '$17': 17,
        '$18': 18,
        '$19': 19,
        '$20': 20,
        '$21': 21,
        '$22': 22,
        '$23': 23,
        '$24': 24,
        '$25': 25,
        '$26': 26,
        '$27': 27,
        '$28': 28,
        '$29': 29,
        '$30': 30,
        '$31': 31,
        '$zero': 0,
        '$at': 1,
        '$v0': 2,
        '$v1': 3,
        '$a0': 4,
        '$a1': 5,
        '$a2': 6,
        '$a3': 7,
        '$t0': 8,
        '$t1': 9,
        '$t2': 10,
        '$t3': 11,
        '$t4': 12,
        '$t5': 13,
        '$t6': 14,
        '$t7': 15,
        '$s0': 16,
        '$s1': 17,
        '$s2': 18,
        '$s3': 19,
        '$s4': 20,
        '$s5': 21,
        '$s6': 22,
        '$s7': 23,
        '$t8': 24,
        '$t9': 25,
        '$k0': 26,
        '$k1': 27,
        '$gp': 28,
        '$sp': 29,
        '$fp': 30,
        '$ra': 31,
    }

    def regToNumber(self, register):
        return format(self.__regMap[register], '05b')

# Classe que mapeia as operaçoes com tuplas


class Operacoes:

    # (opCode, format, funct)
    __opCodeToTuple = {
        'xor': (0,'RN',38),#tipo R normal
        'addu': (0,'RN',33),
        'addiu': (9,'I'),#tipo I
        'and': (0,'RN',36),
        'andi': (12,'I'),
        'slt': (0,'RN',42),
        'sll': (0,'RS',0),#tipo R shift
        'srl': (0,'RS',2),
        'ori': (13,'I'),
        'sw': (43,'ILS'),#tipo I load/store
        'lw': (35,'ILS'),
        'lui': (15,'IL'),#tipo I lui
        'beq': (4,'BI'),
        'bne': (5,'BI'),
        'j': (2,'J'),#tipo J
        'jr': (0,'JR',8)#JR
    }

    def getOpcode(self, opCode):
        return format(self.__opCodeToTuple[opCode][0], '06b')

    def getType(self,opCode):
        return self.__opCodeToTuple[opCode][1]

    def getFunct(self,opCode):
        return format(self.__opCodeToTuple[opCode][2], '06b')


cdirectory = os.getcwd()#coletando a pasta atual no sistema
INI_ADRESS = 0x00400000 #endereco inicial do programa
#inicializando classes
R = Registradores()
O = Operacoes()

def s32(value): #traduz um hexa em complemento de 2 para um decimal com sinal
    return -(value & 0x80000000) | (value & 0x7fffffff)

def to_twoscomplement(value,bits): #converte decimal em bin complemento de 2
    if value < 0:
        value = ( 1<<bits ) + value
    formatstring = '{:0%ib}' % bits
    return formatstring.format(value)

def encode(text,filename):
    fileout = cdirectory+"/out/out_"+filename
    f = open(fileout,"w+")
    
    global INI_ADRESS
    labels = {} #dicionario para colocar as labels encontradas
    ctext = []  #inicializando lista para colocar o texto pronto para ser codificado
    lst_idx = -1 #variavel para controle de linhas da lista ctext, sera utilizada para auxilio com labels

    for i in range(len(text)):
        if not(text[i].startswith('.text')) and not(text[i].startswith('.globl')): #linhas ignoradas pelo codificador
            lst_idx += 1 #incrementa a posicao de lst_idx sempre que existir uma linha valida de instrucao
            if text[i].startswith('main:'): #se encontrar label main
                labels['main'] = lst_idx #salva no dicionario o label e a respectiva linha
                ctext.append((text[i])[5:].lstrip().rstrip()) #adiciona ao txt corrigido sem o lbl e cortando espacos no comeco e fim da linha
            elif text[i].startswith('L_'): #se encontrar label comum
                idx = text[i].find(':') #coleta o indice do final do label
                labels[(text[i])[:idx]]=lst_idx #salva no dict o label e a respectiva linha
                ctext.append((text[i])[idx+1:].lstrip().rstrip())
            else:
                ctext.append(text[i].lstrip().rstrip())

    for i in range(len(ctext)):
        f.write(codificaInstr(ctext[i],labels,i)+'\n')


def codificaInstr(linha,labels,index): #linha do comando, labels existentes no codigo, indice da linha
    idxf_instr = linha.find(' ')#procura o indice do primeiro espaco na linha (sera o fim da instrucao)
    instr = linha[:idxf_instr]
    linha = linha[idxf_instr:].lstrip() #remove instrucao e espacos da linha
    elems = linha.split(',') #divide a linha em funcao das virgulas
    
    for i in range(len(elems)):
        elems[i] = elems[i].lstrip().rstrip() #remove espacos no comeco e fim de cada elemento
    tipo = O.getType(instr)
    opcode = O.getOpcode(instr)


    if tipo=='RN': 
        func = O.getFunct(instr)
        rd = R.regToNumber(elems[0])
        rs = R.regToNumber(elems[1])
        rt = R.regToNumber(elems[2])
                                #opcode(6) rs(5) rt (5) rd(5) shift(5) func(6)
        ret = "{0:#0{1}x}".format(int(opcode.zfill(6)+rs+rt+rd+'0'.zfill(5)+func,2),10)
    
    elif tipo=='RS':
        func = O.getFunct(instr)
        rd = R.regToNumber(elems[0])
        #nao ha RS em shifts
        rt = R.regToNumber(elems[1])
        shamt = bin(int(elems[2],16))[2:].zfill(5)
                        #opcode(6) rs(5) rt(5) rd(5) shmt(5) func(6)
        ret = "{0:#0{1}x}".format(int(opcode.zfill(6)+'0'.zfill(5)+rt+rd+shamt+func,2),10)
    
    elif tipo=='I':
        rt = R.regToNumber(elems[0])
        rs = R.regToNumber(elems[1])
        imm = bin(int('0x'+(elems[2])[-4:],16))[2:].zfill(16) #ajustando o hexa para 16 bits e transformando em binario
                    #opcode(6) rs(5) rt (5) imm(16)
        ret = "{0:#0{1}x}".format(int(opcode.zfill(6)+rs+rt+imm,2),10)
    
    elif tipo=='IL':
        rt = R.regToNumber(elems[0])
        imm = bin(int('0x'+(elems[1])[-4:],16))[2:].zfill(16) #ajustando o hexa para 16 bits e transformando em binario
                    #opcode(6) rs(5) rt (5) imm(16)
        ret = "{0:#0{1}x}".format(int(opcode.zfill(6)+'0'.zfill(5)+rt+imm,2),10)

    elif tipo=='BI':
        rs = R.regToNumber(elems[0])
        rt = R.regToNumber(elems[1])

        posLbl = labels[elems[2]]

        offset = labels[elems[2]]-index-1
        offset = to_twoscomplement(offset,16)
                    #opcode(6) rs(5) rt (5) offset(16)
        ret = "{0:#0{1}x}".format(int(opcode.zfill(6)+rs+rt+offset,2),10)
    elif tipo=='ILS':
        rt = R.regToNumber(elems[0])
        separate = elems[1].split('(')
        imm = bin(int((separate[0])[-4:],16))[2:].zfill(16)
        rs = R.regToNumber((separate[1])[:-1])
                    #opcode(6) rs(5) rt (5) immediate(16)
        ret = "{0:#0{1}x}".format(int(opcode.zfill(6)+rs+rt+imm,2),10)
    
    elif tipo=='J':
        lbl = labels[elems[0]]
        steps = lbl*4 #cada linha eh 1 instrucao simples
        target = INI_ADRESS+steps #soma o endereço inicial do programa com as instrucoes andadas
        target = bin(target)[2:].zfill(32)#coleta o valor binario 32bits do target do jump
        target = target[4:-2]#remove os 4 bits mais significativos e os 2 menos significativos
        ret = "{0:#0{1}x}".format(int(opcode.zfill(6)+target,2),10)

    else: #tipo=='JR':
        func = O.getFunct(instr)[1:]#por padrao, funct retorna 6 bits, a funcao de JR tem 5 bits
        rs = R.regToNumber(elems[0])#a lista elems so tera um elemento nesse caso o RS
        #ind hex + opcode(6) + rs(5) + 0(16) + func(5) 
        ret = "{0:#0{1}x}".format(int(opcode.zfill(6)+rs+'0'.zfill(16)+func,2),10)

    return ret