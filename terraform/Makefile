.PHONY: init
init:
	cd env/${ENV}; terraform init

fmt:
	cd env/${ENV}; terraform fmt -check=true
	cd env/${ENV}; terraform validate

plan:
	cd env/${ENV}; terraform plan -parallelism=60

apply:
	cd env/${ENV}; terraform apply -auto-approve -parallelism=60

destroy:
	cd env/${ENV}; terraform destroy -auto-approve -parallelism=60
