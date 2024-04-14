![alt text](image.png)

![alt text](image-1.png)

![alt text](image-2.png)

The above issue should be ideally solved with - https://go.dev/doc/tutorial/generics
But here we will use interfaces to solve it (Generics were introduced recently)

Updated:- The above shuffle should actuall be solved with interfaces only since it is not receiving any values as input and is acting as a method available on type so solving it with interface is a proper method.

![alt text](image-3.png)
![alt text](image-4.png)

![alt text](image-5.png)