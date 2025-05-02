# Distributed Notification Delivery Service

## Overview

This project implements a Distributed Notification Delivery Service for Paper.Social. The system handles notifications for users when a post is published. It uses Go, gRPC, and GraphQL to simulate receiving posts, notifying followers, and querying notifications. 

### Objectives
1. **gRPC Service**: Implement the `PublishPost` endpoint to simulate receiving a new post and sending notifications.
2. **Notification Dispatch Queue**: Build an in-memory queue to dispatch notifications using Go routines and workers.
3. **GraphQL API**: Provide an API to retrieve the 20 most recent notifications for a user.

### Technologies Used
- Go (for backend development)
- gRPC (for communication)
- GraphQL (for querying notifications)
- In-memory data storage
- Go Routines and Worker Pool (for concurrency)

### Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone https://github.com/priyanka0123456/Distributed-notification-delivery-service/tree/master
   cd distributed-notification-delivery-service
