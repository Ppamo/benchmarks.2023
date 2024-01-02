# Benchmark HTMLStrip


## Objetivo

Tenemos aquí cinco diferentes funciones para eliminar las etiquetas HTML de los strings de texto.

Primero, **StripHtmlTagsWithRegexpP1** y **StripHtmlTagsWithRegexpP2** utilizan expresiones regulares con diferentes enfoques. **StripHtmlTagsWithRegexpP1** es el más sencillo de todos, Declarando simplemente una expresión regular <.*?> y reemplazando todos los patrones que la casen con un string vacío. Es interesante, pero su simplicidad es a la vez su mayor ventaja y su mayor desventaja. Algunas etiquetas HTML complejas podrían eludir esta regex debido a su simplicidad.

Por otro lado, **StripHtmlTagsWithRegexpP2** trata de ser más precisa al capturar y eliminar las etiquetas HTML. Su expresión regular (<\/?[a-zA-A]+?[^>]\*\/?>)\* es más estricta. Pero hay una sorpresa, ordena todas las etiquetas encontradas por longitud antes de eliminarlas, algo inusual pero inteligente para prevenir que las etiquetas más largas y anidadas se corten por las más cortas.

En el otro extremo del espectro, **StripHtmlTagsWithStringBuilderV1** y **StripHtmlTagsWithStringBuilderV2**, que deciden abandonar las expresiones regulares a favor de la manipulación de caracteres de base baja. Recorren el string carácter por carácter, buscando caracteres que indiquen el comienzo o el final de las etiquetas HTML y construyendo una cadena que solo incluye los caracteres fuera de las etiquetas. La diferencia clave aquí es que V2 se asegura de que el texto se convierta en un array de runas, permitiendo manejar de forma segura textos con caracteres especiales.

Finalmente, **StripHtmlTagsWithBlueMonday** utiliza un paquete de terceros llamado "bluemonday", famoso por su política de "Seguridad por defecto". Lo que hace es crear una política que simplemente despoja a todas las etiquetas y utiliza esa política para desinfectar el texto. Por supuesto, depende totalmente de cómo está codificada la política bluemonday.

En cuanto al rendimiento, todo depende de la situación específica. Las funciones de regex tienden a ser menos eficientes para largas secuencias de caracteres debido a la naturaleza de las expresiones regulares, mientras que las funciones StringBuilder son más rápidas y eficientes en la memoria para grandes strings de texto. Mientras tanto, Bluemonday es un arma de doble filo: puede ser extremadamente eficaz y rápida para etiquetas HTML sencillas y conocidas, pero podría ser más lenta si debe analizar y desinfectar etiquetas complejas o desconocidas.

&nbsp;

## Prueba

Para facilitar la ejecución de la prueba, se empaqueta el código Go en un Dockerfile, con un script llamado **runner.sh** que se ha creado para facilitar la construcción y el inicio del contenedor que ejecutará la prueba.   Éste script simplemente construye la imágen Docker si ésta no existe, y a continuación inicia un contenedor, el cual ejecutará directamente las pruebas.   La prueba simplemente toma un texto html de entrada y genera un texto sin los tags html, la función se ejecuta 1000 veces, para facilitar el análisis de los resultados.   El contenedor ejecutará el comando por defecto definido en el Dockerfile, el cual es:

```sh
$$ go test -bench=. -benchtime=100x ./...
```

A continuación podemos ver la salida del script y los resultados del benchmarking realizado al final.

```sh
$$ runner.sh

> Building image bm-striphtml:0.1.0
#0 building with "desktop-linux" instance using docker driver

#1 [internal] load build definition from Dockerfile
#1 transferring dockerfile:
#1 transferring dockerfile: 193B done
#1 DONE 0.0s

#2 [internal] load .dockerignore
#2 transferring context: 2B done
#2 DONE 0.0s

#3 [internal] load metadata for docker.io/library/golang:1.21.5-bullseye
#3 DONE 0.0s

#4 [1/2] FROM docker.io/library/golang:1.21.5-bullseye
#4 DONE 0.0s

#5 [2/2] WORKDIR /go/src/ppamo/striphtml
#5 CACHED

#6 exporting to image
#6 exporting layers done
#6 writing image sha256:ea4d3d5b876310c66b52fe943e8f3c98276d90d05a637624dae7bbb160f5a85e done
#6 naming to docker.io/library/bm-striphtml:0.1.0 done
#6 DONE 0.0s
< done!
> Running app:
go: downloading github.com/microcosm-cc/bluemonday v1.0.26
go: downloading github.com/stretchr/testify v1.8.4
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading github.com/aymerick/douceur v0.2.0
go: downloading golang.org/x/net v0.17.0
go: downloading github.com/gorilla/css v1.0.0
goos: linux
goarch: arm64
pkg: ppamo/striphtml/striphtml
Benchmark_StripHtmlTagsWithRegexpP1-8                100            492856 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           6866630 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             67551 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            287027 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            659591 ns/op
PASS
ok      ppamo/striphtml/striphtml       1.695s

```

&nbsp;

## Resultados

Partiendo desde la base, todas las funciones listadas aquí están diseñadas para realizar una tarea similar: eliminar las etiquetas HTML de una cadena de texto dada. A pesar de compartir el mismo objetivo, los métodos que cada función emplea y las librerías que usan son diferentes, lo que puede influir en su eficiencia y rendimiento.

Empezaremos hablando del rendimiento de cada función según la benchmarking proporcionado y explicaremos por qué ciertas funciones pueden rendir mejor que otras.

1. StripHtmlTagsWithRegexpP1: Esta función emplea la librería de regex para identificar y eliminar las etiquetas HTML en el texto. El uso de las expresiones regulares puede ser computacionalmente caro, debido a cómo deben ser interpretadas y ejecutadas. En este caso, la función ha mostrado un rendimiento de 492856 ns/op, lo que significa que, en promedio, cada operación lleva 492856 nanosegundos para ser completada.

2. StripHtmlTagsWithRegexpP2: Como la anterior, utiliza regex, pero en lugar de reemplazar las etiquetas en un solo paso, primero las identifica, las clasifica por longitud y luego las elimina una por una. Esta secuencia adicional de operaciones parece llevar a un tiempo de ejecución mucho más largo: 6866630 ns/op. Esto puede deberse a la sobrecarga añadida por la manipulación y iteración a través de la lista de grupos.

3. StripHtmlTagsWithStringBuilderV1 & StripHtmlTagsWithStringBuilderV2: Estas funciones utilizan el enfoque de StringBuilder para generar una nueva cadena mientras se recorre la cadena original. En lugar de buscar y reemplazar etiquetas, están construyendo activamente la cadena deseada a medida que avanzan. Su rendimiento (67551 ns/op y 287027 ns/op) es significativamente mejor que los enfoques de regex, lo que puede ser resultado de evitar la sobrecarga de las operaciones de regex y trabajar directamente con las representaciones de cadena. Nota: El enfoque V2 es ligeramente menos eficiente, debido a que está trabajando con una representación "rune" de la cadena que puede requerir más tiempo para procesar.

4. StripHtmlTagsWithBlueMonday: Esta función utiliza la librería bluemonday para eliminar las etiquetas HTML. BlueMonday se diseña específicamente para eliminar HTML y otros scripts no deseados de una manera eficiente y segura mediante la implementación de una Política de blanqueamiento, y su rendimiento (659591 ns/op) está en línea con esto.

Es importante tener en cuenta que, aunque el rendimiento es una consideración clave, no es la única que hay que tener en cuenta al seleccionar un enfoque para esta tarea. La precisión y seguridad de la eliminación de HTML también son factores importantes a considerar. Los enfoques basados en regex, por ejemplo, pueden ser vulnerables a errores si los patrones no están perfectamente definidos, mientras que las bibliotecas como bluemonday proporcionan garantías sólidas a ese respecto.

&nbsp;

## Anexos

A continuación se muestra el resultado de 10 diferentes ejecuciones del mismo comando anteriormente descrito, para verificar la calidad de los resultados analizados, la tendencia se mantiene:

*01*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            537309 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           6977976 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             67902 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            269218 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            654885 ns/op
```

*02*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            484826 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           6957376 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             72391 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            275817 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            646238 ns/op
```

*03*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            484955 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           6931575 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             70980 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            266409 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            638714 ns/op
```

*04*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            482492 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           7009362 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             69044 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            282045 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            669110 ns/op
```

*05*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            582355 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           7294747 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             87876 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            286010 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            706970 ns/op
```

*06*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            571620 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           7486661 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             81027 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            281950 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            770513 ns/op
```

*07*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            484528 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           7702265 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             71599 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            291977 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            637630 ns/op
```

*08*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            627720 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           7422970 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             74623 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            283887 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            693702 ns/op
```

*09*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            530898 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100          11707125 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             79306 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            331964 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100           1013548 ns/op
```

*10*
```sh
Benchmark_StripHtmlTagsWithRegexpP1-8                100            545114 ns/op
Benchmark_StripHtmlTagsWithRegexpP2-8                100           7651951 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV1-8         100             92815 ns/op
Benchmark_StripHtmlTagsWithStringBuilderV2-8         100            277378 ns/op
Benchmark_StripHtmlTagsWithBlueMonday-8              100            673960 ns/op
```
