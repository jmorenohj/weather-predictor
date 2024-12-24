# Weather Predictor

**Weather Predictor** es un proyecto desarrollado en **GO** que utiliza **Fiber** para el servidor web y **MongoDB** para la persistencia de datos. Este proyecto tiene como objetivo predecir condiciones meteorológicas, específicamente días de sequía, días de lluvia y días con condiciones óptimas, a través de cálculos matemáticos y simulaciones. 

## Características principales

- **Días de Sequía**: El cálculo de los días de sequía se realiza de dos maneras:
  1. **Simulación**: Se simula el proceso completo para determinar los días de sequía.
  2. **Sistema de Congruencias**: Se resuelve el sistema de congruencias \( ax \mod 360 = bx \mod 360 = cx \mod 360 \), donde los valores de \( a \), \( b \), y \( c \) corresponden a 1, -5, y 3 respectivamente.

- **Días de Lluvia**: Para calcular los días de lluvia, se utiliza un **algoritmo geométrico** que determina si un punto se encuentra dentro de un polígono o no. Este algoritmo traza un segmento desde el punto de origen hacia el infinito y cuenta la cantidad de veces que se atraviesan otros segmentos del polígono. Si el número de intersecciones es impar, el punto se encuentra dentro del polígono.

- **Días con Condiciones Óptimas**: Los días con condiciones óptimas se calculan generando la ecuación de una recta a partir de los dos primeros puntos y verificando si el tercer punto se encuentra sobre dicha recta. Dado que las condiciones iniciales de las velocidades angulares no permiten coincidencias entre los puntos, los días óptimos en este modelo son cero. Esta conclusión fue corroborada de manera visual y manual.

  - Video de la simulación: [Ver video](https://drive.google.com/file/d/12effeu-wSEWJzOCxbTLBTLGvmGfYNMsI/view?usp=sharing).

## API

Para el desarrollo de la API se utilizó **Fiber**, que maneja las rutas y peticiones. La API interactúa con una base de datos **MongoDB** para almacenar y recuperar los datos necesarios para realizar las predicciones meteorológicas.

- **Router**: Se implementó un router encargado de gestionar las rutas y las peticiones HTTP.
- **Modelo de Datos**: El modelo utilizado para la interacción con la base de datos está diseñado para almacenar la información necesaria de manera eficiente.
  
Puedes probar las rutas y las peticiones utilizando el archivo de **Postman** incluido, que contiene ejemplos de las peticiones disponibles.

