CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    asset_id INTEGER NOT NULL,                             
    user_id UUID NOT NULL,                                 
    order_type VARCHAR(10) NOT NULL,
    amount NUMERIC(30, 10) NOT NULL,                      
    price NUMERIC(30, 10) NOT NULL,                       
    status VARCHAR(20) NOT NULL,      
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
)