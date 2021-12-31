# mag-stadistics-dna-proccesed-function
Este servicio se encarga de generar las estadisticas de todas las secuencias de ADN procesadas en el servicio https://github.com/jpaomat/mag-find-mutants-function que determina si un humano es mutante.

## Instalación
***
installation.
```
$ git clone https://github.com/jpaomat/mag-stadistics-dna-processed-function.git
$ cd mag-stadistics-dna-processed-function
```
## Analisis de la calidad del código ⚙️

En sonarCloud puede ver el reporte que se genera al análizar la calidad del código:

[sonar-cloud-mag-stadistics-dna-processed-function](https://sonarcloud.io/summary/new_code?branch=feature%2Frefactor&id=jpaomat_mag-stadistics-dna-processed-function)

## Ejecutando linter ⚙️

Para la ejecución del linter para el análisis de la calidad del código use `npm run lint` este le mostrara si hay algún por en la sintaxis del código.

## Despliegue 📦

Este proyecto solo se puede probar una vez este desplegado en la nube de AWS, para esto solo tiene que subir los cambios que realice al repositorio de GIT con los comandos:

Hacer commit `git commit -m "text to commit"`
Subir cambios `git push origin feature/nombre_rama`

Es importante tener en cuenta que este proyecto esta configurado para que al hacer push previamente se ejecuten los comandos del linter y las pruebas unitarias, por lo que solo deja subir cambios si los anteriores comandos se ejecutan exitosamente.

Una vez se suban los cambios empieza a ejecutarse el pipeline que se encarga de desplegar el código en el servicios Lambda de AWS.

[mag-stadistics-dna-processed-function/actions](https://github.com/jpaomat/mag-stadistics-dna-processed-function/actions) workflows de GithubActions 

## Consumo del Servicio
En el siguiente Link encuentras la documentación para hacer el consumo del servicio /stast a traves de suagger o generar el curl para su ejecución en POSTMAN:

[Api mag-mutantns_api -swagger](https://app.swaggerhub.com/apis-docs/jpaomat/mag-mutantns_api/v1)
