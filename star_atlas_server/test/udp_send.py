import socket


if __name__ == '__main__':
    bins =  [
        "test_udp-14da6df7-a955-4c83-bd5a-fc2af9eadbc8.bin",
        "test_udp-a71c333c-5cd3-4408-8467-309ec2e98360.bin"
    ]
    for i in bins:
        with open(i, 'rb') as f:
            udp_socket = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
            dest_address = ('127.0.0.1', 9191)
            send_data = f.read()
            udp_socket.sendto(send_data, dest_address)
            udp_socket.close()
