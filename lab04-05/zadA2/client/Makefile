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

run:
	python3 client.py
