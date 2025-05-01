# Pattern Definition
-   Define 1-to-many relationship between objects
-   Automatically notify all observers when the state of the subject changes
-   Belongs to the Behavioral Pattern category
-   Suitable for event driven scenarios

# Core features
-   **Event-driven**: The change of subject state triggers the update of the observers
-   **Loose Coupling**: The subject and the observers only interact through abstract interfaces, promoting decoupling and flexibility
-   **Dynamic Subscription**: Observers can register or unregister anytime.

# Advantages
-   Support broadcast style for event notifitcations.
-   Easy to expands new observer types.
-   Comply with the Open/Closed Principle.

# Disadvantages
-   A large number of observers may lead to performance problems
-   Circular dependencies may lead to memory leaks
-   The subject needs to maintain the list of observers

# Scenario Application

- *Realtime notification system*: Users subscribe to the subject to receive the real-time notifitcations
- *Stock Quotation Platform*: Investors subscribe to stock codes to get price changes
- *E-commerce Inventory Monitoring*: Users subscribe to product to receive inventory change reminders
- *Distributed system monitoring*: The monitoring service subscribes to the state changes of microservices
