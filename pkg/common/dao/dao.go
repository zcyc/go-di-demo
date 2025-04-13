package dao

import (
	"fmt"
	"time"
)

// UserDAO 定义用户数据访问接口
type UserDAO interface {
	Get(id string) (string, error)
	Save(id string, data string) error
}

// InMemoryUserDAO 实现基于内存的用户数据访问层
type InMemoryUserDAO struct {
	data map[string]string
}

// NewInMemoryUserDAO 创建新的内存用户DAO实例
func NewInMemoryUserDAO() *InMemoryUserDAO {
	return &InMemoryUserDAO{
		data: make(map[string]string),
	}
}

// Get 获取用户数据
func (dao *InMemoryUserDAO) Get(id string) (string, error) {
	fmt.Printf("[%s] [DAO] 获取用户数据，ID: %s\n", time.Now().Format(time.RFC3339), id)
	if val, ok := dao.data[id]; ok {
		return val, nil
	}
	return "", fmt.Errorf("用户不存在: %s", id)
}

// Save 保存用户数据
func (dao *InMemoryUserDAO) Save(id string, data string) error {
	fmt.Printf("[%s] [DAO] 保存用户数据，ID: %s\n", time.Now().Format(time.RFC3339), id)
	dao.data[id] = data
	return nil
}

// ProductDAO 定义产品数据访问接口
type ProductDAO interface {
	GetProduct(id string) (string, error)
	SaveProduct(id string, data string) error
}

// InMemoryProductDAO 实现基于内存的产品数据访问层
type InMemoryProductDAO struct {
	data map[string]string
}

// NewInMemoryProductDAO 创建新的内存产品DAO实例
func NewInMemoryProductDAO() *InMemoryProductDAO {
	return &InMemoryProductDAO{
		data: make(map[string]string),
	}
}

// GetProduct 获取产品数据
func (dao *InMemoryProductDAO) GetProduct(id string) (string, error) {
	fmt.Printf("[%s] [DAO] 获取产品数据，ID: %s\n", time.Now().Format(time.RFC3339), id)
	if val, ok := dao.data[id]; ok {
		return val, nil
	}
	return "", fmt.Errorf("产品不存在: %s", id)
}

// SaveProduct 保存产品数据
func (dao *InMemoryProductDAO) SaveProduct(id string, data string) error {
	fmt.Printf("[%s] [DAO] 保存产品数据，ID: %s\n", time.Now().Format(time.RFC3339), id)
	dao.data[id] = data
	return nil
}
