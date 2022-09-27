# Code auto-generation for fanout and chain topologies
## Details
all the connections are gRPC  
## Usage
1. Define configuration file in `configs` directory
2. generate code
```
ansible-playbook main.yaml --extra-vars "config_file=name_of_config.yaml"
```
3. generated code is in `generated` directory

## TODO
1. K8s deployment with skaffold
  - Prepare k8s manifest including all the services
  - Prepare skaffold manifest for building and delpoying all the services
2. Allow the choice between single db instance or db per serivce COMPLETED
3. make the root service use REST COMPLETED
