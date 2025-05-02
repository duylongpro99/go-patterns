# Pattern Definition
-   Dynamically add new function to objects
-   Extend object responsibility through composition rather than inheritance
-   Belongs to the **Structural Pattern**
-   Suitable for scenarios with *functional layering*

# Core features
-   **Transparent Extension**: Keep interface consistent, client unaware of changes
-   **Composition Feature**: Support nesting of multiple layers of decorators
-   **Runtime Binding**: Dynamically determine the combination of object functions

# Advantages
-   Avoid inheritance hierarchy explosion
-   Free combinations of functions
-   Comply with the Open/Closed Principle

# Disadvantages
-   Debugging dificult with multiple decorator layers
-   Excessive use will increase system complexity
-   The order of decorators may affect to execution result

# Scenario Application
-   **API Middleware**: Stack functions such as logging, authentication, rate limiting, etc.
-   **Ecommerce Order Processing**: Calculate the total order amount step by step
-   **Cloud Storage Service**: Dynamically combine storage functions
-   **Micoservice Gateway**: Request verification, response transformation, traffic monitoring
