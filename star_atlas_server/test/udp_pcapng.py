import pcapng
import socket

if __name__ == '__main__':
    b = []
    with open('test.pcapng', 'rb') as f:
        scanner = pcapng.FileScanner(f)
        for block in scanner:
            if block.magic_number == 6:
                #print(block.packet_data)
                b.append(block.packet_data)
                # udp_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
                # dest_address = ('127.0.0.1', 9191)
                # udp_socket.sendto(block.packet_data, dest_address)
                # udp_socket.close()
    b = b"".join(b)
    with open('test.bin', 'wb') as f:
        f.write(b)
        
