CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW(),
                       is_verified BOOLEAN DEFAULT FALSE NOT NULL
);

CREATE TABLE kyc_verification (
                                  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
                                  ssn_last4 CHAR(4) NOT NULL,
                                  dob DATE NOT NULL,
                                  address TEXT NOT NULL,
                                  kyc_status bool DEFAULT FALSE NOT NULL,
                                  submitted_at TIMESTAMP NOT NULL,
                                  verified_at TIMESTAMP,
                                  PRIMARY KEY(user_id)
);

CREATE TABLE wallets (
                         user_id BIGINT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
                         balance_cents BIGINT DEFAULT 0 NOT NULL,
                         locked_cents BIGINT DEFAULT 0 NOT NULL,
                         updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE markets (
                         id BIGSERIAL PRIMARY KEY,
                         title TEXT NOT NULL,
                         description TEXT NOT NULL,
                         category VARCHAR(100) NOT NULL,
                         status VARCHAR(20) NOT NULL,
                         created_at TIMESTAMP DEFAULT NOW() NOT NULL,
                         closes_at TIMESTAMP NOT NULL,
                         resolves_at TIMESTAMP NOT NULL
);

CREATE TABLE contracts (
                           id BIGSERIAL PRIMARY KEY,
                           market_id BIGINT REFERENCES markets(id) ON DELETE CASCADE,
                           contract_type VARCHAR(10) CHECK (contract_type IN ('YES', 'NO')) NOT NULL,
                           price_cents INT NOT NULL ,
                           volume BIGINT DEFAULT 0 NOT NULL
);

CREATE TABLE orders (
                        id BIGSERIAL PRIMARY KEY,
                        user_id BIGINT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
                        contract_id BIGINT REFERENCES contracts(id) ON DELETE CASCADE NOT NULL,
                        order_type VARCHAR(10) CHECK (order_type IN ('buy', 'sell')) NOT NULL,
                        order_style VARCHAR(10) CHECK (order_style IN ('market', 'limit')) NOT NULL,
                        price_cents INT NOT NULL,
                        quantity INT NOT NULL,
                        status VARCHAR(10) DEFAULT 'open' NOT NULL,
                        created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE trades (
                        id BIGSERIAL PRIMARY KEY,
                        buy_order_id BIGINT NOT NULL,
                        sell_order_id BIGINT NOT NULL,
                        contract_id BIGINT REFERENCES contracts(id) ON DELETE CASCADE NOT NULL,
                        price_cents INT NOT NULL,
                        quantity INT NOT NULL,
                        executed_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE TABLE transactions (
                              id BIGSERIAL PRIMARY KEY,
                              user_id BIGINT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
                              type VARCHAR(20) NOT NULL,
                              amount_cents BIGINT NOT NULL,
                              balance_after BIGINT NOT NULL,
                              created_at TIMESTAMP DEFAULT NOW() NOT NULL,
                              reference_id UUID DEFAULT gen_random_uuid()
);

CREATE TABLE market_resolution (
                                   market_id BIGINT PRIMARY KEY REFERENCES markets(id) ON DELETE CASCADE NOT NULL,
                                   outcome VARCHAR(10) CHECK (outcome IN ('YES', 'NO'))NOT NULL ,
                                   resolved_by BIGINT REFERENCES users(id) NOT NULL ,
                                   resolved_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE TABLE audit_logs (
                            id BIGSERIAL PRIMARY KEY,
                            user_id BIGINT REFERENCES users(id),
                            action TEXT NOT NULL,
                            metadata JSONB,
                            ip_address INET,
                            created_at TIMESTAMP DEFAULT NOW()
);
