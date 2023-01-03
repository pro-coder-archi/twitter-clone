#* K3D

create-cluster:
	k3d cluster create --config ./cluster.config.yaml

start-cluster:
	k3d cluster start twitter-clone-cluster

stop-cluster:
	k3d cluster stop twitter-clone-cluster

delete-cluster:
	k3d cluster delete twitter-clone-cluster

#* argoCD

get-argocd-server-password:
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo

show-argocd-ui:
	kubectl port-forward svc/argocd-server -n argocd 9080:443

#* goLang Migrate and cockroachDB

download-cockroachdb-cluster-cert:
	curl --create-dirs -o ./cockroachdb.root.crt -O https://cockroachlabs.cloud/clusters/7d6f23bf-3d6c-421b-ac2f-5d16c988bea7/cert

generate-migration-files:
	migrate create -ext sql -dir $(path) -seq $(sequence)

apply-migrations:
	chmod +x ./scripts/run-cockroachdb-migration.sh
	bash ./scripts/run-cockroachdb-migration.sh $(operation) $(path)