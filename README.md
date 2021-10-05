# Descripción
Ejemplo de webservice de nivel 2 con Golang en Gin (Webservice), Swagger y Gorm (ORM)

# Requisitos

## Servidor MySql
Hay que tener levantado una base de datos puedes usar el siguiente comando para poner a punto a traves de docker la Base de datos:
```
MySQL:`docker run --name mysql5 --hostname mysql5 -v C:\Users\MaQuiNa1995\workspace\docker\mysql-go:/var/lib/mysql --network bridge -e MYSQL_ROOT_PASSWORD=pass -d -p 3306:3306 mysql:5`
```

## Base de datos 
Se requiere crear la base de datos de la aplicación en nuestro servidor Sql podemos usar a traves de docker el PhpMyAdmin (Las tablas serán auto generadas por Gorm asique no te preocupes por ellas)

```
PHPMyAdmin: `docker run --name phpmyadmin --hostname phpmyadmin -d --network bridge -p 8080:80 -e PMA_HOSTS=mysql5 --link mysql5:db -e PMA_VERBOSES="Interna" phpmyadmin/phpmyadmin:latest`
```

# Ejecutar el proyecto
Para ejecutar el proyecto:

```
swag init // para regenerar los documentos de swagger (solo si los modificas)
go run .\Main.go // Ejecuta el Main de la aplicación
```

## Ir a Swagger
despues de levantar la aplicación en la consola saldrá:
```
PS D:\Users\MaQuiNa1995\Workspace\go_workspace\Go_Webservice> go run .\Main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (1 handlers)
[GIN-debug] POST   /cancion                  --> maquina1995/webservice/controller.Create (1 handlers)
[GIN-debug] GET    /cancion                  --> maquina1995/webservice/controller.FindAll (1 handlers)
[GIN-debug] GET    /cancion/:id              --> maquina1995/webservice/controller.FindById (1 handlers)
[GIN-debug] PATCH  /cancion                  --> maquina1995/webservice/controller.Update (1 handlers)
[GIN-debug] DELETE /cancion/:id              --> maquina1995/webservice/controller.Delete (1 handlers)
[GIN-debug] Listening and serving HTTP on localhost:8081
```
Tendrás que ir a la siguiente ruta en tu navegador: http://localhost:8081/swagger/index.html para ir a swagger
o si lo prefieres ejecutar comandos curl apuntanto a los distintos endpoint (las rutas las teneis en el código de arriba)
