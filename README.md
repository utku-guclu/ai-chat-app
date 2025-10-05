┌──────────────────┐               
│      Client      │               
│ (Browser/React)  │               
└───────┬──────────┘               
        │ (HTTP/S)                  
        │ (WebSocket Upgrade)       
        ▼                           
┌──────────────────┐               
│      Nginx       │               
│ (Reverse Proxy)  │               
└───────┬──────────┘               
        │ 1. /          (Forward to React:3000)
        │ 2. /api/      (Forward to Go:8080)
        │ 3. /ws/       (Forward/Upgrade to Go:8080)
        ▼                           
┌──────────────────┐  (Internal Docker Network) 
│    Go Backend    │                          
│ (API, WS Server) │                          
└───────┬─────┬────┘                          
        │     │                             
        │     ▼ 1. Data Persistence (SQL)
        │   ┌──────────────────┐          
        │   │    PostgreSQL    │          
        │   │    (Database)    │          
        │   └──────────────────┘          
        │                                 
        ▼ 2. Real-Time Messaging (Pub/Sub)
      ┌──────────────────┐                
      │       Redis      │                
      │ (Message Broker) │                
      └──────────────────┘
