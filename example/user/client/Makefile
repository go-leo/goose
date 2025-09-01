create_user:
	curl --location 'http://localhost:8000/v1/user' --header 'Content-Type: application/json' --data '{"name":"Leo"}'

	echo "create_user done\n"
delete_user:
	curl --location --request DELETE 'http://localhost:8000/v1/user/10000'

	echo "delete_user done\n"
modify_user:
	curl --location --request PUT 'http://localhost:8000/v1/user/10000' --header 'Content-Type: application/json' --data '{"name":"Leo"}'

	echo "modify_user done\n"
update_user:
	curl --location --request PATCH 'http://localhost:8000/v1/user/10000' --header 'Content-Type: application/json' --data '{"id":"99999","name":"Leo"}'

	echo "update_user done\n"
get_user:
	curl --location 'http://localhost:8000/v1/user/10000'

	echo "get_user done\n"
list_user:
	curl --location 'http://localhost:8000/v1/users?page_num=1&page_size=10'

	echo "list_user done\n"
run_client: create_user delete_user modify_user update_user get_user list_user
	echo "run_client done"
