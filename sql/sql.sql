-- 创建数据库，如果不存在
-- DO
-- $$
--     BEGIN
--         IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'business_platform') THEN
--             CREATE DATABASE business_platform;
--         END IF;
--     END
-- $$;

-- 切换到新创建的数据库
\c business_platform

-- 创建表 public_key_info
CREATE TABLE IF NOT EXISTS public_key_info
(
    id         SERIAL PRIMARY KEY,
    public_key VARCHAR   NOT NULL,
    address    VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT idx_public_key_info_public_key_address UNIQUE (public_key, address)
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_public_key_info_public_key ON public_key_info (public_key);
CREATE INDEX IF NOT EXISTS idx_public_key_info_address ON public_key_info (address);