# Proyecto 3 - Análisis y Diseño de Algoritmos

## Descripción

Este proyecto consiste en la implementación de los algoritmos **Move-To-Front (MTF)** e **Improved Move-To-Front (IMTF)** utilizando el lenguaje de programación Go.

El objetivo es analizar el comportamiento de listas autoorganizadas y comparar el desempeño de ambos algoritmos sobre diferentes secuencias de solicitudes, tal como se solicita en el Proyecto No. 3 del curso Análisis y Diseño de Algoritmos.

## Estructura del Proyecto

```text
.
├── main.go
├── mtf.go
├── imtf.go
├── analysis.go
├── go.mod
└── README.md
```

### Archivos

* **main.go**: Ejecuta todas las pruebas solicitadas en el proyecto.
* **mtf.go**: Implementación del algoritmo Move-To-Front (MTF).
* **imtf.go**: Implementación del algoritmo Improved Move-To-Front (IMTF).
* **analysis.go**: Funciones auxiliares para calcular y mostrar el mejor y peor caso de MTF.
* **go.mod**: Configuración del módulo de Go.

## Requisitos

* Go 1.20 o superior.

## Compilación y Ejecución

Clonar el repositorio:

```bash
git clone https://github.com/DufreyM/Proyecto3-ada.git
cd Proyecto3-ADA
```

Ejecutar el programa:

```bash
go run .
```

## Resultados Obtenidos

| Caso                   | Costo Total |
| ---------------------- | ----------- |
| Pregunta 1             | 90          |
| Pregunta 2             | 67          |
| Mejor Caso MTF         | 20          |
| Peor Caso MTF          | 100         |
| 20 veces el elemento 2 | 22          |
| 20 veces el elemento 3 | 23          |
| IMTF (mejor caso)      | 20          |
| IMTF (peor caso)       | 60          |

## Autor

Leonardo Dufrey Mejía Mejía

Universidad del Valle de Guatemala

Análisis y Diseño de Algoritmos
