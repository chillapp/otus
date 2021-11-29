# Для запуска приложения применяем сервис, деплой и ингресс
kubectl apply -f otus-config.yaml -f mysql.yaml -f service.yaml -f deployment.yaml -f ingress.yaml
<br>
или 
<br>
skaffold run
<br>
или
<br>
helm install otus ./otus-chart # тут у меня проблема, не стартует зависимая пода с mysql
