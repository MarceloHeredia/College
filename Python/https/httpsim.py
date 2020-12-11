#author: Marcelo Heredia
import sys
from Crypto.Hash import SHA256
from Crypto.Cipher import AES
from Crypto.Util.Padding import unpad
from Crypto.Util.Padding import pad
import Crypto.Random
import math

p = 0xb10b8f96a080e01dde92de5eae5d54ec52c99fbcfb06a3c69a6a9dca52d23b616073e28675a23d189838ef1e2ee652c013ecb4aea906112324975c3cd49b83bfaccbdd7d90c4bd7098488e9c219a73724effd6fae5644738faa31a4ff55bccc0a151af5f0dc8b4bd45bf37df365c1a65e68cfda76d4da708df1fb2bc2e4a4371
g = 0xa4d1cbd5c3fd34126765a442efb99905f8104dd258ac507fd6406cff14266d31266fea1e5c41564b777e690f5504f213160217b4b01b886a5e91547f9e2749f4d7fbd7d3b9a92ee1909d0d2263f80a76a6a24c087a091f531dbf0a0169b6a28ad662a4d18e73afa32d779d5918d08bc8858f4dcef97c2a24855e6eeb22b3b2e5
a = 0x30d0fe7f631de71f4aa542791ed4f16032427b30b575c89c90716f42e14039ee8c56d708fb6392c722325c093726fadb6a1e188e5b805aa3f4b12b1d2db7a3a9b1c11dcf5f3fcf37525b219efec12e2a0f82d1ff5497ef9204b54c07a691e738f06419dc8100f478ce189b769df3c173db9428072d65606bb37989abd8f31ac

#pow (g,a,p) (g^a)%p

#dados de exemplo:
#B = 0x10E16402D232A9675EC44224D070D08EBACB583813F64CDC738DAA1ADBA07F6E8598951C1D92A0775A5C323BD3765EF85196D9ADA2D014855D20F684693F53BDBC46E880DB874270412549FC02BB78348C7ACDD04E7D349291C9528A8E3B5030C05C6B46C596F6C69625EA57834DD1419C73A326FA1FE4C381DE61646503E224
#MSG1 = 0x7DF2DC1C0E9FB66EA5249C508B0B832AB8811F1472CD75F374F372161256F1F8F15BD7F3E0DCA9C14BF92C6FBBA9283BBC8A5691B8469EFC3A0A70F330978DFB798E67009CBF75CA882776E0374313139A9517608A8D069AA03FF1E388288ACCA92F3ED6A1DF19F5ABE6BAF640E44CE5FC10218773172E80E9528A486378C655
#MSG2 = 0x788AA899F6CDA8FCED4B9C3D17D3BE3077FAD4987EB446265550526D2A32CBAE8089F352FCCB1FB24629C7FB922FA561EDD1BA2FBF4846F46851BF6EE23CAF6EF6786E636A15504ADE77EDCC4007E92AEA928C5B30871DA29593ED38BD03D38C18EC1146D5BCF05604BA684AEC4AC1BA

def encrypt_msg(txt, S):
    revers = str(txt)[::-1] #reverse text
    encoded = revers.encode('utf-8') #encode
    to_cipher = bytearray(encoded) #turn into a byte array
    key = bytearray.fromhex(S) #convert key to bytearray
    iv = bytearray.fromhex(Crypto.Random.get_random_bytes(128).hex()[:32]) #generate a random 128bits IV

    cipher = AES.new(key, AES.MODE_CBC, iv)

    ct_bytes = cipher.encrypt(pad(to_cipher, AES.block_size))

    print(iv.hex() + ct_bytes.hex())

def decipher_msg(Msg, S):
    iv = bytearray.fromhex(Msg[:32]) #get IV [128]
    ct = bytearray.fromhex(Msg[32:]) #get ciphertext [rest of hex]

    key = bytearray.fromhex(S) #turn the key into byte array
    cipher = AES.new(key, AES.MODE_CBC, iv) #initialize AES
    plaintext = unpad(cipher.decrypt(ct), AES.block_size).decode("utf-8") #decipher
    print(plaintext)
    return plaintext

def generate_S(B):
    A = pow(g,a,p) #A = g^a mod p
    print("A = ", hex(A))
    V = pow(B,a,p) #V = B^a mod p
    nob = int(math.ceil(V.bit_length()/8))
    vbytes = V.to_bytes(nob, byteorder='big') #turn V into bytes array

    hash_obj = SHA256.new(vbytes)

    S = hash_obj.hexdigest()[:32]

    print("S = ", S)

    return S

def main(argv):
    if (len(argv) >= 1):  # need 1 argument (the file name plus extension)
        B = int(argv[0],16)     #to receive B as argument
    else:
        print("Type B value: ")
        B = int(input(),16)
    S = generate_S(B)

    print("Type the Message: ")
    MSG = hex(int(input(), 16))[2:]

    plaintext = decipher_msg(MSG, S)

    encrypt_msg(plaintext, S)

if __name__ == '__main__':
    main(sys.argv[1:])

#pow(g,a,p)