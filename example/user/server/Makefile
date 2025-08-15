
build:
	docker build -t gonic-user:v1 .

run:
	docker run -d -p 8000:8000 gonic-user:v1

create_user:
	curl --location 'http://localhost:8000/v1/user' --header 'Content-Type: application/json' --data '{"name":"Leo"}'

delete_user:
	curl --location --request DELETE 'http://localhost:8000/v1/user/10000'

modify_user:
	curl --location --request PUT 'http://localhost:8000/v1/user/10000' --header 'Content-Type: application/json' --data '{"name":"Leo"}'

update_user:
	curl --location --request PATCH 'http://localhost:8000/v1/user/10000' --header 'Content-Type: application/json' --data '{"id":"99999","name":"Leo"}'

get_user:
	curl --location 'http://localhost:8000/v1/user/10000'

list_user:
	curl --location 'http://localhost:8000/v1/users?page_num=1&page_size=10'
