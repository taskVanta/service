git clone https://github.com/taskVanta/service.git
cd service

go mod tidy

go run main.go


├── main.go                  # Entry point
├── go.mod                   # Go module definition
├── handlers/                # Route handlers
├── models/                  # Data models
├── middleware/              # Custom middleware
├── routes/                  # Route definitions
├── utils/                   # Utility functions
├── database/                # DB connection and logic
└── README.md                # Project documentation


