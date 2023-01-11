terraform {
    backend "local" { }

    required_providers {

        kubectl = {
            source = "gavinbunney/kubectl"
            version = "1.14.0"
        }

        helm = {
            source = "hashicorp/helm"
            version = "2.8.0"
        }
    }
}

provider "kubectl" {

    config_path = local.kubeconfig.path
    config_context = local.kubeconfig.context
}

provider "helm" {

    kubernetes {
        config_path = local.kubeconfig.path
        config_context = local.kubeconfig.context
    }
}