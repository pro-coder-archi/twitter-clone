resource "helm_release" "argocd" {
    name = "argocd"

    namespace = "argocd"
    create_namespace = true

    repository = "https://argoproj.github.io/argo-helm"
    chart = "argo-cd"
    version = "5.16.7"

    wait = true
}

#* creating the argocd application manager

data "kubectl_file_documents" "argocd_application_manager_documents" {
    content = file("./manifests/applications-manager.argocd-application.yaml")
}

resource "kubectl_manifest" "argocd_application_manager_manifest" {
    depends_on = [ helm_release.argocd ]

    for_each = data.kubectl_file_documents.argocd_application_manager_documents.manifests
    yaml_body = each.value

    wait = true
}