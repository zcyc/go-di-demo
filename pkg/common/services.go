package common

import (
	"fmt"
	"time"
)

// Logger is a simple logging interface
type Logger interface {
	Log(message string)
}

// SimpleLogger implements the Logger interface
type SimpleLogger struct{}

func NewSimpleLogger() *SimpleLogger {
	return &SimpleLogger{}
}

func (l *SimpleLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), message)
}

// Database is a simple database interface
type Database interface {
	Get(id string) (string, error)
	Save(id string, data string) error
}

// InMemoryDB implements the Database interface
type InMemoryDB struct {
	data map[string]string
	log  Logger
}

func NewInMemoryDB(logger Logger) *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]string),
		log:  logger,
	}
}

func (db *InMemoryDB) Get(id string) (string, error) {
	db.log.Log(fmt.Sprintf("Getting data for ID: %s", id))
	if val, ok := db.data[id]; ok {
		return val, nil
	}
	return "", fmt.Errorf("not found: %s", id)
}

func (db *InMemoryDB) Save(id string, data string) error {
	db.log.Log(fmt.Sprintf("Saving data for ID: %s", id))
	db.data[id] = data
	return nil
}

// UserService uses Database and Logger
type UserService struct {
	db     Database
	logger Logger
}

func NewUserService(db Database, logger Logger) *UserService {
	return &UserService{
		db:     db,
		logger: logger,
	}
}

func (s *UserService) GetUser(id string) (string, error) {
	s.logger.Log(fmt.Sprintf("UserService: Getting user %s", id))
	return s.db.Get(id)
}

func (s *UserService) SaveUser(id string, data string) error {
	s.logger.Log(fmt.Sprintf("UserService: Saving user %s", id))
	return s.db.Save(id, data)
}

// NotificationService sends notifications
type NotificationService struct {
	logger Logger
}

func NewNotificationService(logger Logger) *NotificationService {
	return &NotificationService{
		logger: logger,
	}
}

func (s *NotificationService) Notify(user string, message string) {
	s.logger.Log(fmt.Sprintf("Notification to %s: %s", user, message))
}

// UserManager uses both UserService and NotificationService
type UserManager struct {
	userService         *UserService
	notificationService *NotificationService
}

func NewUserManager(userService *UserService, notificationService *NotificationService) *UserManager {
	return &UserManager{
		userService:         userService,
		notificationService: notificationService,
	}
}

func (m *UserManager) RegisterUser(id string, data string) error {
	if err := m.userService.SaveUser(id, data); err != nil {
		return err
	}
	m.notificationService.Notify(id, "Welcome to our platform!")
	return nil
}

func (m *UserManager) GetUserDetails(id string) (string, error) {
	return m.userService.GetUser(id)
}
