### Strategy Pattern

Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.

### Observer Pattern

Observer is a behavioral design pattern. In this pattern there are two roles: observer, and observable (or in this case, it's called subject). The subject keep track of observer, and which topic are they subscribed to, everytime something change in that topic the subject can then notify every subscriber by invoking a method in their interface.

### Decorator Pattern

Decorator is a structural design pattern where there is a base implementation with addon implementation that "is" the same interface but also "has a" same interface. this way, a base implementation can be decorated with multiple add ons that add new functionality to the same method from the same interface.

### Factory Method Pattern

Factory method is a creational design pattern. In this pattern a product can be instantiated by a factory. A factory is an interface that implements a creation/instantiation, but this instantiation process is deferred to its subclasses, or in this case to its implementation. A simple factory pattern is exactly this except it directly uses concrete implementation of a factory instead of using interface. 

### References

- [Refactoring Guru](https://refactoring.guru/)
- [Design Patterns in Object Oriented Programming by Christopher Okhravi](https://www.youtube.com/playlist?list=PLrhzvIcii6GNjpARdnO4ueTUAVR9eMBpc)
- Head First Design Patterns Second Edition by Eric Freeman, Elisabeth Robson, Bert Bates, and Kathy Sierra
- Design Pattern Elements of Reusable Object-Oriented Software by Erich Gamma, John Vlissides, Richard Helm, and Ralph Johnson
