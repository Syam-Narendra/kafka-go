# Kafka Cart Abandonment Notifier 🚀

This project uses **Kafka (Sarama)** to track **cart events** and **order events** in an eCommerce system. If a user **adds an item to the cart but does not place an order within 30 minutes**, a **notification** is triggered.

## 🔹 Features
- ✅ **Real-time tracking** of user actions (cart additions, orders).
- ✅ **Kafka-based event-driven architecture** for scalability.
- ✅ **Automated notifications** for abandoned carts.
- ✅ **In-memory or Redis-based tracking** for efficient monitoring.

## 🔹 Architecture
1. **Kafka Producers** → Publish \`add_to_cart\` and \`order_placed\` events.
2. **Kafka Consumer (\`cart_watcher_service\`)** → Monitors user cart activity.
3. **Abandoned Cart Checker** → Runs every 30 minutes to detect inactive carts.
4. **Notification System** → Sends email/SMS/push reminders to users.

## 🔹 Tech Stack
- **Go** (Sarama Kafka Client)
- **Apache Kafka** (Event Streaming)
- **Redis** (Optional, for session storage)
- **Docker** (For containerization)

## 🔹 Setup Instructions
1. **Start Kafka** using Docker:
   ```sh
   docker-compose up -d
   ```
2. **Run the Consumer Service**:
   ```sh
   go run main.go
   ```
3. **Produce Sample Events** (using a Kafka producer or test script).

## 🔹 How It Works
- A user **adds a product to the cart** → Event stored in `cart_events` Kafka topic.
- A user **places an order** → Event stored in `order_events` Kafka topic.
- If no order is placed within **30 minutes**, the system **sends a notification**.

## 🔹 Example Output
```
User user_123 added to cart: prod_001
User user_456 added to cart: prod_002
User user_123 placed an order: order_789
Sending reminder to user_456 for product prod_002
🔔 Notification sent to user user_456: 'Complete your purchase for prod_002!'
```

## 🔹 Next Steps
✅ **Use Redis instead of in-memory storage** for scalability.  
✅ **Use Kafka's Streams API** for better real-time processing.  
✅ **Integrate an actual notification system** like Twilio (SMS), Firebase (push notifications), or email services.

🚀 Now your eCommerce backend can send real-time reminders for abandoned carts!

