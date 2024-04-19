import struct

def float_to_bytes(f):
    return struct.pack('>f', f)

# Example usage
int_cp = 256
float_cp = 1
int_cp_byte = float_to_bytes(int_cp)
float_cp_byte = float_to_bytes(float_cp)
error_ratio = float_to_bytes(0.33)
s = 0xeb.to_bytes(1, 'big')
v = s + int_cp_byte + float_cp_byte + error_ratio
v = v + int(235).to_bytes(1, 'big', signed=False)
filename = "computer_power_data.bin"
with open(filename, "wb") as f:
    f.write(v)

with open(filename, "rb") as f:
    bs = f.read()
    print(len(bs))
    s = int.from_bytes(bytes=bs[0:1], byteorder='big')
    print(s == 0xeb)
    f1 = struct.unpack('>f',bs[1:5])[0]
    print(f1)
    f1 = struct.unpack('>f',bs[5:9])[0]
    print(f1)
    f1 = struct.unpack('>f',bs[9:13])[0]
    print(f1)
    f1 = int.from_bytes(bytes=bs[13:14], byteorder='big')
    print(f1)
