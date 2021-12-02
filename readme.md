# Запускаем
```shell
kubectl apply -f otus-config.yaml -f mysql.yaml -f service.yaml -f deployment.yaml -f ingress.yaml
```
ИЛИ
```shell
skaffold run
```
ИЛИ
```shell
helm install otus ./otus-chart
```
