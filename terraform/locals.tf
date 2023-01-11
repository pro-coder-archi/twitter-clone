locals {
    kubeconfig= {

        path = "~/.kube/config"
        context = "k3d-twitter-clone-cluster"
    }
}