# Pattern definition
-   Encapsulate algorithms into independent strategy classes
-   Support dynamic switching of strategies at runtime
-   **Behavioral Pattern**
-   Suitable with *multiple algorithms*

# Core features
- **Algorithm Encapsulation**: Each strategy independently implements the algorithm
- **Strategy Switching**: Dynamically switch strategies at runtime
- **Open-Closed Principle**: Adding new strategy does not require modifying existing code

# Advantages
- Decouple algorithm from client
- Facilitate algorithm extension and replacement
- Support combine strategies

# Disadvantages
- The client needs to know all strategies types
- The number of strategy classes may increase rapidly
- Difficult to share state between strategies

# Scenario Applications
- *Ecommerce Promotion Engine*: Dynamically switch the calculation logic according to the activity types
- *Logistics Distributed System*: Calculate the shipping fee according to user's choice.
- *Data Sorting Service*: The frontend request determine the sorting strategy.
- *Advertisement Placement Algorithm*: Adjust the placement strategy in real-time.
