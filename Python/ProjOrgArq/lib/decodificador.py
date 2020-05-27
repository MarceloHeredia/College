#Marcelo Heredia e Pedro Castro

import os
import array
import sys

cdirectory = os.getcwd()#coletando a pasta atual no sistema


#classe que mapeia registradores
class Registradores:
    __regMap = {
        0: '$0',    #$zero
        1: '$1',    #$at
        2: '$2',    #$v0
        3: '$3',    #$v1
        4: '$4',    #$a0
        5: '$5',    #$a1
        6: '$6',    #$a2
        7: '$7',    #$a3
        8: '$8',    #$t0
        9: '$9',    #$t1
        10: '$10',  #$t2
        11: '$11',  #$t3
        12: '$12',  #$t4
        13: '$13',  #$t5
        14: '$14',  #$t6
        15: '$15',  #$t7
        16: '$16',  #$s0
        17: '$17',  #$s1
        18: '$18',  #$s2
        19: '$19',  #$s3
        20: '$20',  #$s4
        21: '$21',  #$s5
        22: '$22',  #$s6
        23: '$23',  #$s7
        24: '$24',  #$t8
        25: '$25',  #$t9
        26: '$26',  #$k0
        27: '$27',  #$k1
        28: '$28',  #$gp
        29: '$29',  #$sp
        30: '$30',  #$fp
        31: '$31',  #$ra
    }
    def numberToReg(self,register):
        return self.__regMap.get(int(register,2))

#classe que mapeia instrucoes do tipo R
class InstrR:
    __instr = {
        0x26: 'xor',
        0x24: 'and',
        0x21: 'addu',
        0x0:  'sll',
        0x2:  'srl',
        0x2a: 'slt',
        0x8:  'jr'  #licenca poetica...
    }
    def instr(self,nmbr):
        return self.__instr.get(int(nmbr,2))

#classe que mapeia instrucoes do tipo I
class InstrI:
    __instr = {
        0x9: 'addiu',
        0xc: 'andi',
        0xd: 'ori',
        0xf: 'lui',
        0x4: 'beq',
        0x5: 'bne',
        0x23:'lw',
        0x2b:'sw'
    }

    def instr(self, nmbr):
        return self.__instr.get(int(nmbr,2))

#instanciacao das classes
R = Registradores()
IR = InstrR()
II = InstrI()
count = 1

#coleta valor do negativo em complemento de 2
def twos_comp(val, bits):#coleta negativo complemento de 2
    """compute the 2's complement of int value val"""
    if (val & (1 << (bits - 1))) != 0: # if sign bit is set e.g., 8bit: 128-255
        val = val - (1 << bits)        # compute negative value
    return val   

#transforma numero negativo em hexa complemento de 2
def bindigits(n, bits): #preenche complemento de 2 com '1's
    s = bin(n & int("1"*bits, 2))[2:]
    return ("{0:0>%s}" % (bits)).format(s)

#transforma hexa comp 2 em decimal com sinal
def s16(value):
    return -(int(value,16) & 0x80000000) | (int(value,16) & 0x7fffffff)

#processo principal de decodificacao
def decode(text, filename):
    fileout = cdirectory+"/out/out_"+filename
    f = open(fileout,"w+")
    f.write('.text\n'
            '.globl main\n'
            )
    text2 = ["" for x in range (len(text))]
    text2[0]='main:'

    for i in range(len(text)):
        if (text[i].startswith('0x')):
            text[i] = (text[i])[2:]

    for i in range(len(text)):
        text[i] = str(bin(int('0x'+text[i],16))[2:].zfill(32))
        text2[i] += (decInstr((text[i])[0:6], text[i], text2,i)+'\n')

    for i in range(len(text2)):
        f.write(text2[i])

def decInstr(opcode, line, text, i):
    if line.startswith('L_'):
        idx = linefind(':')
        cleanline = linha_texto[idx:]
    else:
        cleanline = line

    if int(opcode) == 0: #tipo R (6,5,5,5,5,6)(op,rs,rt,rd,smt,func)
        rs = cleanline[6:11]  #RS
        rt = cleanline[11:16] #RT
        rd = cleanline[16:21] #RD
        smt = cleanline[21:26]#SHAMT
        fnc = cleanline[26:32]#FUNCAO
        return decInstrR(rs,rt,rd,smt,fnc)

    elif int(opcode,2) == 2:#verificando se e J (jump)
        return decInstrJ(line, text,i) 

    else: #tipo I (6,5,5,16)(op,rs,rt(dest),immediate)
        op = cleanline[0:6]
        rs = cleanline[6:11]
        rt = cleanline[11:16]
        imm = cleanline[16:32]
        return decInstrI(op,rs,rt,imm,text,i)



def decInstrR(rs,rt,rd,smt,fnc):
    instr = IR.instr(fnc)
    rs = R.numberToReg(rs)
    rt = R.numberToReg(rt)
    rd = R.numberToReg(rd)
    smt = "{0:#0{1}x}".format((int(smt,2)),10)

    if instr == 'jr':
        return '{0} {1}'.format(instr,rs)
        

    if instr == 'sll' or instr == 'srl':
        dados = (rd,rt,smt)
    else:
        dados=(rd,rs,rt)
    
    x = ('{0} {1},{2},{3}'.format(instr,dados[0],dados[1],dados[2]))
    return x


def decInstrI(op,rs,rt,imm,text,pos): #tipo I op $rs $rt immediate/offset
    global count
    instr = II.instr(op)
    rs = R.numberToReg(rs)
    rt = R.numberToReg(rt)
    imm = hex(int(bindigits(twos_comp(int(imm,2),16),32),2))
    if len(imm) < 10:
        imm = "{0:#0{1}x}".format((int(imm,16)),10)

    if instr == 'addiu' or instr == 'andi' or instr == 'ori':
        dados=(rt,rs,imm)
        x = ('{0} {1},{2},{3}'.format(instr,dados[0],dados[1],dados[2]))

    elif instr == 'lui':
        dados=(rt,imm)
        x = ('{0} {1},{2}'.format(instr,dados[0],dados[1]))
    elif instr == 'beq' or instr == 'bne':
        linha_destino = pos+s16(imm)+1
        linha_texto = text[linha_destino]

        if linha_texto.startswith('L_'):
            idx = linha_texto.find(':')
            label = linha_texto[0:idx]
        elif linha_texto.startswith('main:'):
            label = 'main'
        elif linha_destino == pos:
            label = 'L_'+str(count)
            linha_texto = label+':'+linha_texto
            count+=1
            instr = linha_texto+instr #adicionando label antes da instrucao
        else:
            label = 'L_'+str(count)
            linha_texto = label+':'+linha_texto
            text[linha_destino] = linha_texto
            count+=1
        dados = (rs,rt,label)
        x = ('{0} {1},{2},{3}'.format(instr,dados[0],dados[1],dados[2]))
    else: #lw ou sw
        dados = (rt,imm,rs) #(lw-sw) rt,imm(rs) ou exemplo [lw $14, 28($29)]
        x = ('{0} {1},{2}({3})'.format(instr,dados[0],dados[1],dados[2]))
        
    return x


def decInstrJ(line, text,pos): #tipo J (2,26)(op,target) (0000-target-00)!
    global count
    target = line[6:32]
    target = '0000'+target+'00' #endereco completo 
    adress = int(target,2) - 0x00400000   #endereco completo em inteiros
    linha = int(adress/4) #cada linha tem 1 instrucao e cada instrucao ocupa 4 posicoes do endereco
    linha_texto = text[linha] #coleta a linha onde esta armazenado o 

    if linha_texto.startswith('L_'):
        idx = linha_texto.find(':')
        label = linha_texto[0:idx]
    elif linha_texto.startswith('main:'):
        label = 'main'
    elif linha == pos:
        label = 'L_'+str(count)
        linha_texto = label+':'+linha_texto
        count +=1
        return linha_texto + 'j '+label
    else:
        label = 'L_'+str(count)
        linha_texto = label+':'+linha_texto
        text[linha] = linha_texto
        count+=1
    return 'j '+label