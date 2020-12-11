#author: Marcelo Heredia
import sys
from Crypto.Hash import SHA256

blocks = list()

def read_chunks(file, chunk_size = 1024):#default size = 1Kb
    while True:
        data = file.read(chunk_size)
        if not data: #end of file reached
            break
        yield data #return chunks as a 'generator' to the populate_blocks function

def populate_blocks(file): #save the chunks in a list
    for chunk in read_chunks(file):
        blocks.append(chunk)

def generate_hash():
    print(type(blocks))
    hash_obj = SHA256.new()
    hash_obj.update(blocks[-1])
    hn = hash_obj.digest() #gets the last block hash to start the loop

    for data in reversed(blocks[:-1]): #iterates the reversed list
        data += hn #append n+1 block's hash to the n block data
        hash_obj = SHA256.new(data) #initialize this block's hash
        hn = hash_obj.digest()  #save the hashcode

    return hash_obj.hexdigest() #returns the final hash (h0)

def main(argv):
    if (len(argv) < 1):  # need 1 argument (the file name plus extension)
        print("sorry, try again!")
        print("call example:\npython sha256.py testfile.mp4")
        exit()
    file = open(argv[0], "rb") #opens file as binary data

    populate_blocks(file)

    file.close()

    print(generate_hash())


if __name__ == '__main__':
    main(sys.argv[1:])
