import sha3
import binascii

def bytes32(i):
    return binascii.unhexlify('%064x' % i)

def keccak256(x):
    return sha3.keccak_256(x).hexdigest()

h = binascii.b2a_hex(s)