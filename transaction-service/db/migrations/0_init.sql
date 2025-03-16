CREATE TABLE IF NOT EXISTS transactions (
	transaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	source_account_id UUID NOT NULL,
	target_account_id UUID NOT NULL,
	amount NUMERIC(12,2) NOT NULL,
	status VARCHAR(20) NOT NULL DEFAULT 'pending',
	timestamps TIMESTAMP DEFAULT NOW()
);