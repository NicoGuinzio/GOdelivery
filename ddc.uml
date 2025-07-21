classDiagram
    class User {
        +string ID
        +string Name
        +string Rol
    }
    class Task {
        +string ID
        +string UserID
        +string Title
        +TaskStatus Status
    }
    class TaskStatus {
        <<enumeration>>
        Pendiente
        En~progreso
        Finalizada
    }
    User "1" <|-- "*" Task : owns
    Task o-- TaskStatus : status
    
    class UserRepository {
        +Save(User) error
        +FindByID(id string) (User, error)
        +List() []User
    }
    class TaskRepository {
        +Save(Task) error
        +FindByID(id string) (Task, error)
        +ListByUser(userID string) []Task
        +Update(Task) error
    }
    class InMemoryUserRepo {
        -map[string]User users
        +Save(User) error
        +FindByID(id string) (User, error)
        +List() []User
    }
    class InMemoryTaskRepo {
        -map[string]Task tasks
        +Save(Task) error
        +FindByID(id string) (Task, error)
        +ListByUser(userID string) []Task
        +Update(Task) error
    }
    UserRepository <|.. InMemoryUserRepo
    TaskRepository <|.. InMemoryTaskRepo
    TaskRepository <.. Task
    UserRepository <.. User