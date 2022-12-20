# Twitter Clone
Making a battle tested Twitter clone based on Kubernetes and the CloudNative ecosystem.

## Resources

- Figma - https://www.figma.com/file/WwJGhzeRHZ3nmZPTQ9KH7c/Twitter-Clone?t=3UfE9f8tjEobUHCr-6
- ProtoPie - https://cloud.protopie.io/p/75243de97087df8bc5e63f8d

## Installations

- Proto compiler (ProtoC) -

    ```bash
    sudo apt install -y protobuf-compiler -y

    # verify installation
    proto --version

    # compiler pluins for goLang

    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

    set -U fish_user_paths $HOME/go/bin $fish_user_paths
    set -U fish_user_paths /usr/local/go/bin $fish_user_paths
    ```
