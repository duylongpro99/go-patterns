# Pattern Definition
-   Encapsulate the creation logic of the object
-   Determine the specific product to instantiated through subclasses
-   Belongs to the **Creational Patterns**
-   Decouples the object creation from its usage

# Core features
- **Encapsulate of Creation Logic**: The client does not need to know the creation details.
- **Polymorphism Support**: Extend the products family through the interfaces.
- **Open-Closed Principle**: Add new products does not require modifying existing code.

# Advantages
- Seperate object creation from usage
- Easy to extend new products.
- Comply with the Dependency Inversion Principle

# Disadvantages
- Increase number of classes - each products needs a corresponding factory.
- The factory class may become to complex
- The client needs to know the abstrac interface of the product

# Scenario Application
- *Payment Gateway Integration*: Create different processors according to the payment method.
- *Cloud Storage Client*: Create corresponding API clients according to the storage service.
- *Message Queue Adapter*: Unify the creation logic of the message sending interface.
- *Third-party Login Service*: Dynamically generate authentication clients for different platforms.
