# Guía: Árboles Binarios y Recorridos

## Ejercicios

1. Implementar el árbol binario genérico definido en [`BinaryTree`](binarytree/binarytree.go) y [`BinaryNodo`](binarytree/binarynode.go).

2. Escribir una función recursivo que calcule (y devuelva) la suma de todos los valores pares de los nodos en un árbol binario de enteros. El método debe devolver 0 si el árbol está vacío.

    ![Árbol binario de ejemplo](imagenes/arbol-binario.drawio.svg)

    > 2 + 2 + 6 + 4 =  14

3. Escribir una función recursivo que calcule (y devuelva) la suma de los nodos cuyo valor es mayor que un $k$ dado, de un árbol binario de enteros. El método debe devolver 0 si el árbol está vacío.

    ![Árbol binario de ejemplo](imagenes/arbol-binario.drawio.svg)

    > Mayores que 8, daría 20

4. Escribir una función recursivo que calcule (y devuelva) la mayor altura entre el sub-árbol izquierdo y derecho de un árbol binario de enteros. El método debe devolver -1 si el árbol está vacío.

    ![Árbol binario ejercicio 4](imagenes/arbol-binario-ej04.drawio.svg)

    > 3

5. Escribir una función recursivo que calcule (y devuelva) la suma de todas las hojas de un árbol binario de enteros que son hijas izquierdas de un nodo. El método debe devolver 0 si el árbol está vacío

    ![Árbol binario ejercicio 5](imagenes/arbol-binario-ej05.drawio.svg)

    > 17 + 27 = 44

6. Escribir una función recursivo para un árbol binario de enteros, que calcule (y devuelva) la suma de todos los nodos derechos cuyo valor es impar. El método debe devolver 0 si el árbol está vacío

    ![Árbol binario ejercicio 6](imagenes/arbol-binario-ej06.drawio.svg)

    > 29 + 15 + 11 = 55

7. Escribir una función que permita recorrer el árbol por niveles. Pista usar una cola para registrar los hijos de un nodo que se está procesando y que se visitarán luego de terminar de visitar todos los nodos del nivel actual.

    ![Árbol binario ejercicio 7](imagenes/arbol-binario-ej07.drawio.svg)

    > A, B, C, D, E, F, G,H, I, J, K, L

### Desafio

El ladrón ha encontrado un nuevo lugar para sus robos una vez más. Solo hay una entrada a esta área, llamada raíz.

Salvo por la raíz, cada casa tiene una y solo una casa padre. Después de un recorrido, el inteligente ladrón se dio cuenta de que todas las casas en este lugar forman un árbol binario. Se contactará automáticamente con la policía si se roban dos casas directamente conectadas durante la misma noche.

Dado la raíz del árbol binario, devolver la cantidad máxima de dinero que el ladrón puede robar sin alertar a la policía.

Un ejemplo:

![Árbol binario desafio 1](imagenes/arbol-binario-desafio01.drawio.svg)

- **Salida**: 7
- **Explicación**: La cantidad máxima de dinero que el ladrón puede robar &rarr; 3 + 3 + 1 = 7.

otro:

![Árbol binario desafio 1](imagenes/arbol-binario-desafio02.drawio.svg)

- **Salida**: 9
- **Explicación**: La cantidad máxima de dinero que el ladrón puede robar &rarr; 4 + 5 = 9.
