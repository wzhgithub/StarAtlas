[test]
step 1: install netcat
yum install netcat
nc -ul 127.0.0.1 8080

step 2: 
./utest /dev/stdout | xargs --null echo > /dev/udp/127.0.0.1/8080

OR

for ((i=0; i<=0; i++)); do
  ./utest /dev/stdout | xargs --null echo > /dev/udp/127.0.0.1/8080
done

