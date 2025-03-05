### **Beer Pong Permissions Game – A Symbolic Transaction-Based Access Control System**  

#### **📌 Overview**  
The **Beer Pong Permissions Game** is an event-driven system where symbolic **payment transactions** (ball throws) are used to determine access control in a game-like environment. Instead of traditional role-based access systems, the system **gamifies authentication and validation** using a fake payment flow. Each transaction (throw) must be **authenticated, validated, and encrypted**, making it a **lightweight but extensible model**.  

The game follows the principles of **event-driven architecture**, real-time state validation, and challenge-based dispute resolution. Players interact in a **real-time UI**, where transactions visually "hop" from one cup to another, reflecting the **validation process** live.  

---

#### **🎯 Core Objectives**  
✅ **Symbolic Payment Flow:** Access control with gamified, lightweight transaction validation.  
✅ **Event-Driven Architecture:** Ensure **high scalability and performance** using Kafka, Redis, and WebSockets.  
✅ **Real-Time Transaction Feedback:** Provide instant visibility into **successful throws, failed validations, and challenges**.  
✅ **Tamper-Resistant Validation:** Every throw (transaction) is timestamped, signed, and **immutable** once validated.  
✅ **Low Overhead, High Scalability:** Optimize the system for **fast, lightweight validation** while keeping future adaptability for real financial transactions.  

---

#### **🛠️ How It Works**  
1️⃣ **A player attempts a throw (symbolic payment) at a cup (target endpoint).**  
2️⃣ **The system validates the throw** using authentication (JWT/PASETO) and stores it in Redis with a **5-minute TTL**.  
3️⃣ **Kafka processes the event asynchronously,** verifying the validity of the throw.  
4️⃣ **The cup owner can challenge the transaction** within **15 seconds**.  
5️⃣ **If a challenge occurs,** the system revalidates the throw, and a verdict is returned.  
6️⃣ **The game visually represents these transactions in real time,** showing throws, validations, and disputes.  

---

#### **🛠️ Tech Stack & Key Components**  

| **Component**  | **Purpose**  |
|--------------|------------|
| **Go Fiber API** | Manages game logic and API requests. |
| **WebSockets (Gorilla)** | Pushes real-time validation updates. |
| **Kafka** | Handles throw and challenge events asynchronously. |
| **Redis** | Stores active session data and game state. |
| **PostgreSQL** | Stores validated transactions and game history. |
| **JWT/PASETO** | Authenticates and verifies players. |
| **React UI** | Displays real-time transaction movements visually. |

---

#### **🔐 Security & Validation Measures**  
- **JWT / PASETO authentication** ensures that only verified players can throw.  
- **Timestamped & signed transactions** prevent replay attacks.  
- **15-second challenge window** ensures fairness while preventing indefinite delays.  
- **Redis TTL (Time-to-Live) caching** clears invalidated transactions automatically.  

---

#### **📈 Future Potential & Scalability**  
🚀 **Can be adapted for real transactions.**   
🚀 **Supports additional validation layers (e.g., blockchain signing).**  
🚀 **Easily scalable by adding more Kafka partitions & consumers.**  
🚀 **Can integrate with financial APIs for real payment simulations.**  

This system is designed to be **lightweight, scalable, and adaptable**—whether for gaming, access control, or even financial transaction simulations.  