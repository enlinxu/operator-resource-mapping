# operator-resource-mapping
Operator Resource Mapping (ORM) is a mechanism to achieve configuration consistency between running deployments and CustomResource, for example operator managed applications with VPA and HPA


The CRD is available at `config/crd/bases`

Install by

`kubectl apply -f .config/crd/bases/turbonomic.io_operatorresourcemappings.yaml`

Use `Makefile`

`make install` will generate a new CRD with updates (if any) and install it to the current kubectl context
