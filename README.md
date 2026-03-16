# Flow
<img width="6250" height="3165" alt="User Creation Request Flow-2026-03-16-192143" src="https://github.com/user-attachments/assets/04098c57-9a09-4093-a30d-fa1c9a69231e" />



### Code

```bash
sequenceDiagram
participant Client
participant Middleware
participant Handler
participant Service
participant Repository
participant DB
Client->>Middleware: HTTP Request
Middleware->>Handler: forward request
Handler->>Service: CreateUser()
Service->>Repository: SaveUser()
Repository->>DB: INSERT USER
DB-->>Repository: OK
Repository-->>Service: user saved
Service-->>Handler: user response
Handler-->>Client: JSON Response
```

---

# Test

<p align="center">
<img width="1366" height="345" alt="Screenshot (631)" src="https://github.com/user-attachments/assets/dba5e0e2-2b60-4979-8cb1-ac4d76ad8557" />
  <img width="1062" height="629" alt="Screenshot (632)" src="https://github.com/user-attachments/assets/03c02ab7-084b-4b0b-925c-426038940217" />
<img width="1052" height="649" alt="Screenshot (633)" src="https://github.com/user-attachments/assets/faaf1df4-6bba-4536-a8e9-15a43a256e37" />
<img width="1366" height="564" alt="Screenshot (634)" src="https://github.com/user-attachments/assets/11e60db8-09fe-4cfa-8a0b-bb23db248733" />

</p>


# Database

<p align="center">
<img width="1366" height="550" alt="Screenshot (629)" src="https://github.com/user-attachments/assets/409aa76e-4a89-49cb-9d54-f57f8455acc7" />
  <img width="1342" height="550" alt="Screenshot (630)" src="https://github.com/user-attachments/assets/3cb23c82-bb97-4d9f-8233-f6835fe579bf" />

</p>

