cd sr_server
pip3 install -r requirements.txt
python3 asrserver_http.py

cd sr_go_client
go build
./asrt-sdk-go 
