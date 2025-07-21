package main

import (
	"bufio"
	"fmt"
	"os"
)

// ================= DOMAIN LAYER =================
// Entidad User representa un usuario del sistema.
type User struct {
	ID   string // Identificador único
	Name string
	Rol  string
}

// Entidad Task representa una tarea asociada a un usuario.
type Task struct {
	ID     string // Identificador único
	UserID string // Referencia al usuario dueño de la tarea
	Title  string
	Status TaskStatus
}

// Value Object: TaskStatus define el estado de una tarea.
type TaskStatus string

const (
	TaskPending    TaskStatus = "Pendiente"
	TaskInProgress TaskStatus = "En progreso"
	TaskCompleted  TaskStatus = "Finalizada"
)

// ================= APPLICATION LAYER =================
// Casos de uso: lógica de aplicación que orquesta entidades y repositorios.

type UserRepository interface {
	Save(user User) error
	FindByID(id string) (User, error)
	List() []User
}

type TaskRepository interface {
	Save(task Task) error
	FindByID(id string) (Task, error)
	ListByUser(userID string) []Task
	Update(task Task) error
}

// Use case: Crear usuario
func CreateUser(repo UserRepository, name string) (User, error) {
	id := generateID()
	user := User{ID: id, Name: name}
	if err := repo.Save(user); err != nil {
		return User{}, err
	}
	return user, nil
}

// Use case: Crear tarea para un usuario
func CreateTask(repo TaskRepository, userID, title string) (Task, error) {
	id := generateID()
	task := Task{ID: id, UserID: userID, Title: title, Status: TaskPending}
	if err := repo.Save(task); err != nil {
		return Task{}, err
	}
	return task, nil
}

// Use case: Completar tarea
func CompleteTask(repo TaskRepository, taskID string) error {
	task, err := repo.FindByID(taskID)
	if err != nil {
		return err
	}
	task.Status = TaskCompleted
	return repo.Update(task)
}

// Use case: Listar tareas por usuario
func ListTasks(repo TaskRepository, userID string) []Task {
	return repo.ListByUser(userID)
}

// ================= INFRASTRUCTURE LAYER =================
// Implementaciones en memoria de los repositorios.

type InMemoryUserRepo struct {
	users map[string]User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{users: make(map[string]User)}
}

func (r *InMemoryUserRepo) Save(user User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) FindByID(id string) (User, error) {
	user, ok := r.users[id]
	if !ok {
		return User{}, fmt.Errorf("usuario no encontrado")
	}
	return user, nil
}

func (r *InMemoryUserRepo) List() []User {
	users := []User{}
	for _, u := range r.users {
		users = append(users, u)
	}
	return users
}

type InMemoryTaskRepo struct {
	tasks map[string]Task
}

func NewInMemoryTaskRepo() *InMemoryTaskRepo {
	return &InMemoryTaskRepo{tasks: make(map[string]Task)}
}

func (r *InMemoryTaskRepo) Save(task Task) error {
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepo) FindByID(id string) (Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("tarea no encontrada")
	}
	return task, nil
}

func (r *InMemoryTaskRepo) ListByUser(userID string) []Task {
	tasks := []Task{}
	for _, t := range r.tasks {
		if t.UserID == userID {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func (r *InMemoryTaskRepo) Update(task Task) error {
	_, ok := r.tasks[task.ID]
	if !ok {
		return fmt.Errorf("tarea no encontrada")
	}
	r.tasks[task.ID] = task
	return nil
}

// ================= INTERFACE LAYER =================
// CLI simple para interactuar con la aplicación y probar los casos de uso.

func main() {
	userRepo := NewInMemoryUserRepo()
	taskRepo := NewInMemoryTaskRepo()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- To-Do CLI ---")
		fmt.Println("1. Crear usuario")
		fmt.Println("2. Crear tarea")
		fmt.Println("3. Completar tarea")
		fmt.Println("4. Listar tareas por usuario")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")
		opt, _ := reader.ReadString('\n')
		switch opt[0] {
		case '1':
			fmt.Print("Nombre de usuario: ")
			name, _ := reader.ReadString('\n')
			user, err := CreateUser(userRepo, trim(name))
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Usuario creado con ID:", user.ID)
			}
		case '2':
			fmt.Print("ID de usuario: ")
			userID, _ := reader.ReadString('\n')
			fmt.Print("Título de la tarea: ")
			title, _ := reader.ReadString('\n')
			task, err := CreateTask(taskRepo, trim(userID), trim(title))
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Tarea creada con ID:", task.ID)
			}
		case '3':
			fmt.Print("ID de tarea: ")
			taskID, _ := reader.ReadString('\n')
			err := CompleteTask(taskRepo, trim(taskID))
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Tarea completada.")
			}
		case '4':
			fmt.Print("ID de usuario: ")
			userID, _ := reader.ReadString('\n')
			tasks := ListTasks(taskRepo, trim(userID))
			if len(tasks) == 0 {
				fmt.Println("No hay tareas para este usuario.")
			} else {
				fmt.Println("Tareas:")
				for _, t := range tasks {
					fmt.Printf("- [%s] %s (ID: %s)\n", t.Status, t.Title, t.ID)
				}
			}
		case '5':
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

// ================= UTILIDADES =================
// Función simple para generar IDs únicos (no apto para producción)
func generateID() string {
	return fmt.Sprintf("%d", os.Getpid()+int(os.Getuid())+int(os.Geteuid())+int(os.Getppid())+int(os.Getegid()))
}

// trim elimina saltos de línea y espacios
func trim(s string) string {
	return string([]byte(s)[:len(s)-1])
}
