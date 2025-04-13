package service

import (
	"fmt"
	"go-di-demo/pkg/common"
	"go-di-demo/pkg/common/dao"
)

// UserService 处理用户相关业务逻辑
type UserService struct {
	userDAO dao.UserDAO
	logger  common.Logger
}

// NewUserService 创建用户服务实例
func NewUserService(userDAO dao.UserDAO, logger common.Logger) *UserService {
	return &UserService{
		userDAO: userDAO,
		logger:  logger,
	}
}

// GetUser 获取用户信息
func (s *UserService) GetUser(id string) (string, error) {
	s.logger.Log(fmt.Sprintf("[Service] 获取用户信息，ID: %s", id))
	return s.userDAO.Get(id)
}

// SaveUser 保存用户信息
func (s *UserService) SaveUser(id string, data string) error {
	s.logger.Log(fmt.Sprintf("[Service] 保存用户信息，ID: %s", id))
	return s.userDAO.Save(id, data)
}

// ProductService 处理产品相关业务逻辑
type ProductService struct {
	productDAO dao.ProductDAO
	logger     common.Logger
}

// NewProductService 创建产品服务实例
func NewProductService(productDAO dao.ProductDAO, logger common.Logger) *ProductService {
	return &ProductService{
		productDAO: productDAO,
		logger:     logger,
	}
}

// GetProduct 获取产品信息
func (s *ProductService) GetProduct(id string) (string, error) {
	s.logger.Log(fmt.Sprintf("[Service] 获取产品信息，ID: %s", id))
	return s.productDAO.GetProduct(id)
}

// SaveProduct 保存产品信息
func (s *ProductService) SaveProduct(id string, data string) error {
	s.logger.Log(fmt.Sprintf("[Service] 保存产品信息，ID: %s", id))
	return s.productDAO.SaveProduct(id, data)
}

// NotificationService 发送通知
type NotificationService struct {
	logger common.Logger
}

// NewNotificationService 创建通知服务实例
func NewNotificationService(logger common.Logger) *NotificationService {
	return &NotificationService{
		logger: logger,
	}
}

// Notify 向用户发送通知
func (s *NotificationService) Notify(user string, message string) {
	s.logger.Log(fmt.Sprintf("[Service] 发送通知给 %s: %s", user, message))
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
