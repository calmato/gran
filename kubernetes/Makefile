.PHONY: get-external-ip
get-external-ip:
	kubectl get service envoy -o jsonpath='{.status.loadBalancer.ingress[0].ip}'

.PHONY: generate-private-key
generate-private-key:
	openssl req -x509 -nodes \
		-newkey rsa:2048 -days 365 \
		-keyout privkey.pem -out cert.pem \
		-subj "/CN=${EXTERNAL_IP}"

.PHONY: generate-deployment
generate-deployment:
	cd ./${ENV}; sed 's/<PROJECT-ID>/${GCP_STORAGE_BUCKET_NAME}/' deployment-group-api.yaml.sample > deployment-group-api.yaml
	cd ./${ENV}; sed 's/<PROJECT-ID>/${GCP_STORAGE_BUCKET_NAME}/' deployment-todo-api.yaml.sample > deployment-todo-api.yaml
	cd ./${ENV}; sed 's/<PROJECT-ID>/${GCP_STORAGE_BUCKET_NAME}/' deployment-user-api.yaml.sample > deployment-user-api.yaml

.PHONY: update-deployment
update-deployment:
	kubectl patch deploy group-api -p "{\"spec\": {\"template\": {\"metadata\": {\"labels\": {\"date\": \"$(date +'%s')\" }}}}}"
	kubectl patch deploy todo-api -p "{\"spec\": {\"template\": {\"metadata\": {\"labels\": {\"date\": \"$(date +'%s')\" }}}}}"
	kubectl patch deploy user-api -p "{\"spec\": {\"template\": {\"metadata\": {\"labels\": {\"date\": \"$(date +'%s')\" }}}}}"
