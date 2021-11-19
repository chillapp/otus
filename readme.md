# Для запуска приложения применяем сервис, деплой и ингресс
kubectl apply -f service.yaml -f deployment.yaml -f ingress.yaml
<br>
или 
<br>
skaffold run