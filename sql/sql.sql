-- 创建提现状态枚举类型
CREATE TYPE wallet_operator_status AS ENUM (
    'PENDING', -- 待处理
    'CREATE_UN_SIGN_ED', -- 创建未签名的交易完成
    'SIGN_ED', -- 签名完成
    'BUILD_SIGN_ED', -- 构建交易完成
    'SEND_TX_ED', -- 上链完成
    'SAFE_ED', -- 已安全
    'FINALLY_ED', -- 已最终状态
    'FAIL' -- 已取消
    );

-- 创建提现订单表
CREATE TABLE withdraw_order
(
    id             BIGSERIAL PRIMARY KEY,
    -- 基础信息
    order_id       VARCHAR(64) UNIQUE     NOT NULL,                   -- 订单号（业务ID）
    uid            BIGINT                 NOT NULL,                   -- 用户ID

    -- 代币信息
    token_symbol   VARCHAR(20)            NOT NULL,                   -- 代币符号
    token_address  VARCHAR(42),                                       -- 代币合约地址
    chain_id       BIGINT                 NOT NULL,                   -- 链ID

    -- 金额信息
    amount         DECIMAL(36, 18)        NOT NULL,                   -- 提现金额
    fee            DECIMAL(36, 18)        NOT NULL DEFAULT 0,         -- 手续费
    actual_amount  DECIMAL(36, 18)        NOT NULL,                   -- 实际到账金额（amount - fee）

    -- 地址信息
    from_address   VARCHAR(42),                                       -- 来源地址
    to_address     VARCHAR(42)            NOT NULL,                   -- 目标地址

    -- 交易信息
    tx_hash        VARCHAR(66),                                       -- 交易哈希
    nonce          BIGINT,                                            -- 交易nonce
    gas_price      DECIMAL(36, 18),                                   -- gas价格
    gas_limit      BIGINT,                                            -- gas限制

    -- 状态信息
    status         wallet_operator_status NOT NULL DEFAULT 'PENDING', -- 当前状态
    error_msg      TEXT,                                              -- 错误信息

    -- 业务信息
    memo           VARCHAR(255),                                      -- 备注
    operator       VARCHAR(50),                                       -- 操作人
    approve_time   TIMESTAMP,                                         -- 审批时间

    -- 时间戳
    created_at     TIMESTAMP              NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP              NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- 索引
    CONSTRAINT idx_withdraw_order_uid_token UNIQUE (uid, order_id)
);

-- 创建索引
CREATE INDEX idx_withdraw_order_uid ON withdraw_order (uid);
CREATE INDEX idx_withdraw_order_status ON withdraw_order (status);
CREATE INDEX idx_withdraw_order_created_at ON withdraw_order (created_at);
CREATE INDEX idx_withdraw_order_tx_hash ON withdraw_order (tx_hash);