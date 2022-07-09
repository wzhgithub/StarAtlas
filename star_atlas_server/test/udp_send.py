import socket


if __name__ == '__main__':
    for i in range(0, 3):
        with open("sample_{}.bin".format(i), 'rb') as f:
            udp_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
            dest_address = ('127.0.0.1', 9191)
            send_data = f.read()
            udp_socket.sendto(send_data, dest_address)
            udp_socket.close()
