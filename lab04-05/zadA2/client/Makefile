client_path = "client/"
grpcproject_path = "grpcproject/"

.PHONY: prepare
prepare:
	python -m pip install virtualenv
	virtualenv venv
	# source venv/bin/activate
	python -m pip install --upgrade pip
	python -m pip install grpcio grpcio-tools

proto:
	mkdir -p $(grpcproject_path)
	python -m grpc_tools.protoc -I	../$(grpcproject_path) --python_out=. --pyi_out=. --grpc_python_out=.   ../$(grpcproject_path)grpcproject.proto

run_client_0:
	python3 client.py 0

run_client_1:
	python3 client.py 1

run_client_2:
	python3 client.py 2

