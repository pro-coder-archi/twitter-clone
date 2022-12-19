#* K3D

create-cluster:
	k3d cluster create --config ./cluster.config.yaml

start-cluster:
	k3d cluster start cloudnative-tools-testing

stop-cluster:
	k3d cluster stop cloudnative-tools-testing

delete-cluster:
	k3d cluster delete cloudnative-tools-testing

#* argoCD

get-argocd-server-password:
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo

show-argocd-ui:
	kubectl port-forward svc/argocd-server -n argocd 9080:443