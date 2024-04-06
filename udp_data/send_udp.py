import socket

UDP_IP = '127.0.0.1'
UDP_PORT = 9191
with open('xxx.bin', 'rb') as f:
    MESSAGE = f.read();
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.sendto(MESSAGE, (UDP_IP, UDP_PORT))
    sock.close()