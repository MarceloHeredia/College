.text
.globl main
main:
lui $1,0x00001001
ori $8,$1,0x00000000
lw $9,0x00000000($8)
lui $1,0x00001001
ori $10,$1,0x00000004
lw $11,0x00000000($10)
addu $12,$9,$11
lui $1,0x00001001
ori $13,$1,0x00000008
L_1:
sw $12,0x00000000($13)
beq $9,$10,main
and $21,$20,$19
sll $22,$23,0x00000005
L_2:
srl $24,$25,0x00000009
slt $14,$15,$16
addiu $5,$3,0xfffffffb
andi $14,$0,0x0000000f
bne $21,$2,L_1
sw $2,0xfffffffc($20)
j L_1
j main
j L_2
jr $31