package svc

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/zeromicro/go-zero/core/logx"

	"business-platform/app/common/grpcclient/multichain"
	"business-platform/app/common/grpcclient/sign"
	"business-platform/app/exchange/internal/config"
)

type ServiceContext struct {
	Config config.Config

	MultichainClient multichain.GrpcClient
	SignClient       sign.GrpcClient

	MasterDB *sql.DB
	SlaveDB  *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := context.Background()
	logCtx := logx.WithContext(ctx)

	multiChainClient, err := multichain.NewGrpcClient(c.MultiChainRpcConf.Endpoint)
	if err != nil {
		logCtx.Errorf("初始化 multichain 客户端失败: %v", err)
		panic(fmt.Sprintf("初始化 multichain 客户端失败: %v", err))
	}
	logCtx.Info("multichain 客户端初始化成功")

	signClient, err := sign.NewGrpcClient(c.SignRpcConf.Endpoint)
	if err != nil {
		logCtx.Errorf("初始化 sign 客户端失败: %v", err)
		panic(fmt.Sprintf("初始化 sign 客户端失败: %v", err))
	}
	logCtx.Info("sign 客户端初始化成功")

	masterDB, err := createDBConnection(c.DbConf.Master, "master", logCtx)
	if err != nil {
		logCtx.Errorf("连接主数据库失败: %v", err)
		panic(fmt.Sprintf("连接主数据库失败: %v", err))
	}
	logCtx.Info("主数据库连接成功")

	slaveDB, err := createDBConnection(c.DbConf.Slave, "slave", logCtx)
	if err != nil {
		logCtx.Errorf("连接从数据库失败: %v", err)
		panic(fmt.Sprintf("连接从数据库失败: %v", err))
	}
	logCtx.Info("从数据库连接成功")

	return &ServiceContext{
		Config: c,

		MultichainClient: multiChainClient,
		SignClient:       signClient,

		MasterDB: masterDB,
		SlaveDB:  slaveDB,
	}
}

func createDBConnection(dbConfig config.DbConfig, dbType string, logCtx logx.Logger) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbUsername, dbConfig.DbPassword, dbConfig.DbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("打开%s数据库连接失败: %v", dbType, err)
	}

	// Optionally, you can set connection pool settings here
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("测试%s数据库连接失败: %v", dbType, err)
	}

	logCtx.Infof("%s数据库连接池配置完成: MaxOpenConns=25, MaxIdleConns=25, MaxLifetime=5min", dbType)
	return db, nil
}
