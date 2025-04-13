package service

import (
	"fmt"
	"go-di-demo/pkg/common/dao"
	"time"
)

// UserService 处理用户相关业务逻辑
type UserService struct {
	userDAO dao.UserDAO
}

// NewUserService 创建用户服务实例
func NewUserService(userDAO dao.UserDAO) *UserService {
	return &UserService{
		userDAO: userDAO,
	}
}

// GetUser 获取用户信息
func (s *UserService) GetUser(id string) (string, error) {
	fmt.Printf("[%s] [Service] 获取用户信息，ID: %s\n", time.Now().Format(time.RFC3339), id)
	return s.userDAO.Get(id)
}

// SaveUser 保存用户信息
func (s *UserService) SaveUser(id string, data string) error {
	fmt.Printf("[%s] [Service] 保存用户信息，ID: %s\n", time.Now().Format(time.RFC3339), id)
	return s.userDAO.Save(id, data)
}

// ProductService 处理产品相关业务逻辑
type ProductService struct {
	productDAO dao.ProductDAO
}

// NewProductService 创建产品服务实例
func NewProductService(productDAO dao.ProductDAO) *ProductService {
	return &ProductService{
		productDAO: productDAO,
	}
}

// GetProduct 获取产品信息
func (s *ProductService) GetProduct(id string) (string, error) {
	fmt.Printf("[%s] [Service] 获取产品信息，ID: %s\n", time.Now().Format(time.RFC3339), id)
	return s.productDAO.GetProduct(id)
}

// SaveProduct 保存产品信息
func (s *ProductService) SaveProduct(id string, data string) error {
	fmt.Printf("[%s] [Service] 保存产品信息，ID: %s\n", time.Now().Format(time.RFC3339), id)
	return s.productDAO.SaveProduct(id, data)
}

// NotificationService 发送通知
type NotificationService struct{}

// NewNotificationService 创建通知服务实例
func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

// Notify 向用户发送通知
func (s *NotificationService) Notify(user string, message string) {
	fmt.Printf("[%s] [Service] 发送通知给 %s: %s\n", time.Now().Format(time.RFC3339), user, message)
}

// UserManager 综合管理用户和通知
type UserManager struct {
	userService         *UserService
	notificationService *NotificationService
}

// NewUserManager 创建用户管理器实例
func NewUserManager(userService *UserService, notificationService *NotificationService) *UserManager {
	return &UserManager{
		userService:         userService,
		notificationService: notificationService,
	}
}

// RegisterUser 注册用户并发送欢迎通知
func (m *UserManager) RegisterUser(id string, data string) error {
	if err := m.userService.SaveUser(id, data); err != nil {
		return err
	}
	m.notificationService.Notify(id, "欢迎加入我们的平台！")
	return nil
}

// GetUserDetails 获取用户详情
func (m *UserManager) GetUserDetails(id string) (string, error) {
	return m.userService.GetUser(id)
}
