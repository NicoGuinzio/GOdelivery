package main

import (
	"bufio"
	"fmt"
	"os"
)

// ================= CAPA DE DOMINIO =================
// Usuario representa un usuario del sistema (Aggregate Root)
type Usuario struct {
	ID     string // Identificador único del usuario
	Nombre string // Nombre del usuario
	Rol    string // Rol del usuario (ej: admin, usuario)
}

// Tarea representa una tarea asociada a un usuario
type Tarea struct {
	ID        string      // Identificador único de la tarea
	UsuarioID string      // ID del usuario dueño de la tarea
	Titulo    string      // Título o descripción de la tarea
	Estado    EstadoTarea // Estado actual de la tarea
}

// EstadoTarea es un Value Object que define el estado de una tarea
type EstadoTarea string

const (
	Pendiente  EstadoTarea = "Pendiente"
	EnProgreso EstadoTarea = "En progreso"
	Finalizada EstadoTarea = "Finalizada"
)

// ================= CAPA DE APLICACIÓN =================
// Interfaces de repositorios para abstracción de persistencia
type RepositorioUsuario interface {
	Guardar(usuario Usuario) error          // Guarda un usuario
	BuscarPorID(id string) (Usuario, error) // Busca un usuario por ID
	Listar() []Usuario                      // Lista todos los usuarios
}

type RepositorioTarea interface {
	Guardar(tarea Tarea) error                 // Guarda una tarea
	BuscarPorID(id string) (Tarea, error)      // Busca una tarea por ID
	ListarPorUsuario(usuarioID string) []Tarea // Lista tareas de un usuario
	Actualizar(tarea Tarea) error              // Actualiza una tarea
}

// Caso de uso: Crear un usuario
func CrearUsuario(repo RepositorioUsuario, nombre string, rol string) (Usuario, error) {
	id := generarID()
	usuario := Usuario{ID: id, Nombre: nombre, Rol: rol}
	if err := repo.Guardar(usuario); err != nil {
		return Usuario{}, err
	}
	return usuario, nil
}

// Caso de uso: Crear una tarea para un usuario
func CrearTarea(repo RepositorioTarea, usuarioID, titulo string) (Tarea, error) {
	id := generarID()
	tarea := Tarea{ID: id, UsuarioID: usuarioID, Titulo: titulo, Estado: Pendiente}
	if err := repo.Guardar(tarea); err != nil {
		return Tarea{}, err
	}
	return tarea, nil
}

// Caso de uso: Marcar una tarea como finalizada
func FinalizarTarea(repo RepositorioTarea, tareaID string) error {
	tarea, err := repo.BuscarPorID(tareaID)
	if err != nil {
		return err
	}
	tarea.Estado = Finalizada
	return repo.Actualizar(tarea)
}

// Caso de uso: Listar tareas de un usuario
func ListarTareas(repo RepositorioTarea, usuarioID string) []Tarea {
	return repo.ListarPorUsuario(usuarioID)
}

// ================= CAPA DE INFRAESTRUCTURA =================
// Implementación en memoria del repositorio de usuarios
type RepositorioUsuarioMem struct {
	usuarios map[string]Usuario // Mapa de usuarios por ID
}

func NuevoRepositorioUsuarioMem() *RepositorioUsuarioMem {
	return &RepositorioUsuarioMem{usuarios: make(map[string]Usuario)}
}

func (r *RepositorioUsuarioMem) Guardar(usuario Usuario) error {
	r.usuarios[usuario.ID] = usuario
	return nil
}

func (r *RepositorioUsuarioMem) BuscarPorID(id string) (Usuario, error) {
	usuario, ok := r.usuarios[id]
	if !ok {
		return Usuario{}, fmt.Errorf("usuario no encontrado")
	}
	return usuario, nil
}

func (r *RepositorioUsuarioMem) Listar() []Usuario {
	usuarios := []Usuario{}
	for _, u := range r.usuarios {
		usuarios = append(usuarios, u)
	}
	return usuarios
}

// Implementación en memoria del repositorio de tareas
type RepositorioTareaMem struct {
	tareas map[string]Tarea // Mapa de tareas por ID
}

func NuevoRepositorioTareaMem() *RepositorioTareaMem {
	return &RepositorioTareaMem{tareas: make(map[string]Tarea)}
}

func (r *RepositorioTareaMem) Guardar(tarea Tarea) error {
	r.tareas[tarea.ID] = tarea
	return nil
}

func (r *RepositorioTareaMem) BuscarPorID(id string) (Tarea, error) {
	tarea, ok := r.tareas[id]
	if !ok {
		return Tarea{}, fmt.Errorf("tarea no encontrada")
	}
	return tarea, nil
}

func (r *RepositorioTareaMem) ListarPorUsuario(usuarioID string) []Tarea {
	tareas := []Tarea{}
	for _, t := range r.tareas {
		if t.UsuarioID == usuarioID {
			tareas = append(tareas, t)
		}
	}
	return tareas
}

func (r *RepositorioTareaMem) Actualizar(tarea Tarea) error {
	_, ok := r.tareas[tarea.ID]
	if !ok {
		return fmt.Errorf("tarea no encontrada")
	}
	r.tareas[tarea.ID] = tarea
	return nil
}

// ================= CAPA DE INTERFAZ =================
// CLI para interactuar con la aplicación y probar los casos de uso
func main() {
	repoUsuario := NuevoRepositorioUsuarioMem() // Repositorio de usuarios en memoria
	repoTarea := NuevoRepositorioTareaMem()     // Repositorio de tareas en memoria
	lector := bufio.NewReader(os.Stdin)         // Lector de entrada estándar

	for {
		fmt.Println("\n--- To-Do CLI ---")
		fmt.Println("1. Crear usuario")
		fmt.Println("2. Crear tarea")
		fmt.Println("3. Finalizar tarea")
		fmt.Println("4. Listar tareas por usuario")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")
		opcion, _ := lector.ReadString('\n')
		switch opcion[0] {
		case '1':
			fmt.Print("Nombre de usuario: ")
			nombre, _ := lector.ReadString('\n')
			fmt.Print("Rol de usuario: ")
			rol, _ := lector.ReadString('\n')
			usuario, err := CrearUsuario(repoUsuario, limpiar(nombre), limpiar(rol))
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Usuario creado con ID:", usuario.ID)
			}
		case '2':
			fmt.Print("ID de usuario: ")
			usuarioID, _ := lector.ReadString('\n')
			fmt.Print("Título de la tarea: ")
			titulo, _ := lector.ReadString('\n')
			tarea, err := CrearTarea(repoTarea, limpiar(usuarioID), limpiar(titulo))
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Tarea creada con ID:", tarea.ID)
			}
		case '3':
			fmt.Print("ID de tarea: ")
			tareaID, _ := lector.ReadString('\n')
			err := FinalizarTarea(repoTarea, limpiar(tareaID))
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Tarea finalizada.")
			}
		case '4':
			fmt.Print("ID de usuario: ")
			usuarioID, _ := lector.ReadString('\n')
			tareas := ListarTareas(repoTarea, limpiar(usuarioID))
			if len(tareas) == 0 {
				fmt.Println("No hay tareas para este usuario.")
			} else {
				fmt.Println("Tareas:")
				for _, t := range tareas {
					fmt.Printf("- [%s] %s (ID: %s)\n", t.Estado, t.Titulo, t.ID)
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
// generarID genera un ID único simple (no apto para producción)
func generarID() string {
	return fmt.Sprintf("%d", os.Getpid()+int(os.Getuid())+int(os.Geteuid())+int(os.Getppid())+int(os.Getegid()))
}

// limpiar elimina saltos de línea y espacios
func limpiar(s string) string {
	return string([]byte(s)[:len(s)-1])
}
