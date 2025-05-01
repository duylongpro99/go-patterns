# Pattern definition

- Only one instance of a class is created globally.
- One unified access point
- Belongs to creational pattern
- Suitable for situations where global unique control is required.

# Core features

- **Unique instance**: 1 object instance globally
- **Self-initialization**: The class itself is responsible for creating its own instance.
- **Global access**: Provide a access entry through a static method or a global variable.

# Advantages
- Ensure global uniqueness of resource/states
- Reduce repeated creation of resources
- Provide a convenient global access point

# Disadvantages
- Violates the principle of high cohesion with hight module coupling
- Difficult to isolate unit test
- Damage the dependency injection Pattern

# Scenario Applications
- *Microservice Configuration Center*: Manage the configuration of distributed systems uniformly
- *Database Connection Pool*: Control the number of concurrent connections
- *Distributed Log Service*: Avoid repeated creation of log handlers
- *API Rate Limiter*: Share the state limiting state globally
