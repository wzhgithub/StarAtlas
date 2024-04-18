import socket
import sys
import time

UDP_IP = '127.0.0.1'
UDP_PORT = 9191
with open(sys.argv[1], 'rb') as f:
    MESSAGE = f.read();
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    for i in range(0, 1):
        r = sock.sendto(MESSAGE, (UDP_IP, UDP_PORT))
        print(f"send len:{r}bytes")
        time.sleep(1)
    sock.close()
