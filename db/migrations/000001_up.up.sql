CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       username VARCHAR(50) UNIQUE,
                       created_at TIMESTAMP DEFAULT NOW(),
                       is_verified BOOLEAN DEFAULT FALSE
);

CREATE TABLE kyc_verification (
                                  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
                                  ssn_last4 CHAR(4),
                                  dob DATE,
                                  address TEXT,
                                  kyc_status VARCHAR(20),
                                  submitted_at TIMESTAMP,
                                  verified_at TIMESTAMP,
                                  PRIMARY KEY(user_id)
);

CREATE TABLE wallets (
                         user_id BIGINT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
                         balance_cents BIGINT DEFAULT 0,
                         locked_cents BIGINT DEFAULT 0,
                         updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE markets (
                         id BIGSERIAL PRIMARY KEY,
                         title TEXT NOT NULL,
                         description TEXT,
                         category VARCHAR(100),
                         status VARCHAR(20),
                         created_at TIMESTAMP DEFAULT NOW(),
                         closes_at TIMESTAMP,
                         resolves_at TIMESTAMP
);

CREATE TABLE contracts (
                           id BIGSERIAL PRIMARY KEY,
                           market_id BIGINT REFERENCES markets(id) ON DELETE CASCADE,
                           contract_type VARCHAR(10) CHECK (contract_type IN ('YES', 'NO')),
                           price_cents INT,
                           volume BIGINT DEFAULT 0
);

CREATE TABLE orders (
                        id BIGSERIAL PRIMARY KEY,
                        user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
                        contract_id BIGINT REFERENCES contracts(id) ON DELETE CASCADE,
                        order_type VARCHAR(10) CHECK (order_type IN ('buy', 'sell')),
                        order_style VARCHAR(10) CHECK (order_style IN ('market', 'limit')),
                        price_cents INT,
                        quantity INT,
                        status VARCHAR(10) DEFAULT 'open',
                        created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE trades (
                        id BIGSERIAL PRIMARY KEY,
                        buy_order_id BIGINT,
                        sell_order_id BIGINT,
                        contract_id BIGINT REFERENCES contracts(id),
                        price_cents INT,
                        quantity INT,
                        executed_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transactions (
                              id BIGSERIAL PRIMARY KEY,
                              user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
                              type VARCHAR(20),
                              amount_cents BIGINT,
                              balance_after BIGINT,
                              created_at TIMESTAMP DEFAULT NOW(),
                              reference_id UUID DEFAULT gen_random_uuid()
);

CREATE TABLE market_resolution (
                                   market_id BIGINT PRIMARY KEY REFERENCES markets(id),
                                   outcome VARCHAR(10) CHECK (outcome IN ('YES', 'NO')),
                                   resolved_by BIGINT REFERENCES users(id),
                                   resolved_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE audit_logs (
                            id BIGSERIAL PRIMARY KEY,
                            user_id BIGINT REFERENCES users(id),
                            action TEXT,
                            metadata JSONB,
                            ip_address INET,
                            created_at TIMESTAMP DEFAULT NOW()
);
