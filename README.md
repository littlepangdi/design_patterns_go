design_patterns_go

## Design patterns with Golang

### What will we do here:

#### Stage 1:

Golang demos of 22 Design Patterns with several elegant skills

|    type    |     design pattern      |                                                                                                                notes                                                                                                                 |
|:----------:|:-----------------------:|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
|   others   |    Producer-consumer    |                                                                                                      connect through mq\channel                                                                                                      |
|  Creative  |        Singleton        |                                                                                     one global instance,only accessed through specific functions                                                                                     |
|            |         Factory         |                                                     Provide a method to create an object in the superclass, allowing subclasses to determine the type of object to instantiate.                                                      |
|            |    Abstract Factory     |                                                                    different product series ✖ different products + abstract_factory interface = abstract factory                                                                     |
|            |        Prototype        |                                                                                enable you to clone objects without depending on the class they belong                                                                                |
|            |         Builder         |                                                                                   director use builder to manage and complete building procedures                                                                                    |
| Structural |         Adapter         |                                                                             Convert the interface of a class to another interface that the client wants.                                                                             |
|            |         Bridge          |                                                       Split a large class or series of closely related classes into two separate hierarchies of abstraction and implementation                                                       |
|            |        Composite        |                                                                  Use it to group objects into tree structures and work with them as if they were separate objects.                                                                   |
|            |        Decorator        |                                                                     Allows adding new functionality to an existing object without changing its structure(pizza).                                                                     |
|            |         Facade          |                                                                            facade wraps several components and provides add/deduct operations for clients                                                                            |
|            |        Flyweight        |                            eliminates the need to store all data in each object, allowing you to load more objects into a limited amount of memory by sharing the same state shared by multiple objects.                             |
|            |          Proxy          |                                                  The proxy controls access to the original object and allows some processing before and after submitting the request to the object.                                                  |
|  Behavior  | Chain of responsibility |                                                                                         allows you to send requests down a chain of handlers                                                                                         |
|            |         Command         |                                                                                           build a single link between sender and receiver                                                                                            |
|            |        Iterator         |                                                                                 range-over collections without concerning about underlying structure                                                                                 |
|            |        Mediator         |                                                                  restricts direct interaction between objects, forcing them to cooperate through a mediator object.                                                                  |
|            |         Memento         |                                                                                   save and restore states from memento(under caretaker's control)                                                                                    |
|            |          State          |                               abstracts state-related behavior into separate state classes, and letting the original object delegate work to instance of these classes instead of handling it itself.                                |
|            |        Strategy         |                                              define a series of algorithms and put each algorithm into a separate class so that the algorithm objects can be replaced with each other.                                               |
|            |     Template Method     |                                                                         Arrange algorithm steps in advance and initialize with different algorithm instances                                                                         |
|            |         Visitor         |                                                                                           get and process class information outside class                                                                                            |
|            |        Observer         |                                                        A subscription mechanism that notifies multiple other objects that "observe" an object when an object's events occur.                                                         |


#### Stage 2:
demo of mainstream frameworks / packages
- [x] 


