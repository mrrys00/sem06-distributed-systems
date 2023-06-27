prepare_system:
	go get -u -v all
	go list

prepare_users:
	sudo rabbitmqctl add_user admin admin
	sudo rabbitmqctl set_user_tags admin administrator
	sudo rabbitmqctl set_permissions -p / admin ".*" ".*" ".*"

	sudo rabbitmqctl add_user agency agency
	sudo rabbitmqctl set_user_tags agency agency
	sudo rabbitmqctl set_permissions -p / agency ".*" ".*" ".*"

	sudo rabbitmqctl add_user speditor speditor
	sudo rabbitmqctl set_user_tags speditor speditor
	sudo rabbitmqctl set_permissions -p / speditor ".*" ".*" ".*"

run_admin:
	go run -p 3 admin/admin.go

run_agency1:
	go run -p 3 agency/agency.go 1

run_agency2:
	go run -p 3 agency/agency.go 2

run_speditor1:
	go run -p 3 speditor/speditor.go 1

run_speditor2:
	go run -p 3 speditor/speditor.go 2
