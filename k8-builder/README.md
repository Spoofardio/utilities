# Kubernetes YAML Builder

Create Yaml files from combining template and config files

## How to use it


### General Form
This will create a filled out yaml file at `target/example_dev.yaml`

if main points to p1()
```
go run main.go example.yaml example_dev.conf
```

if main points to p2()
```
go run main.go v0.1 staging
```

### Example

example_dev.conf
```
PORT=5672
NAME=rabbitmq
```

example.yaml
```
apiVersion: v1
kind: Service
metadata:
  name: {{NAME}}
  namespace: default
  labels:
    app: {{NAME}}
spec:
  type: NodePort
  ports:
  - port: {{PORT}}
    name: amqp
  selector:
    app: {{NAME}}
```

## Authors

* **Zachary Spofford** - *Initial work* - [Spoofardio](https://github.com/Spoofardio)