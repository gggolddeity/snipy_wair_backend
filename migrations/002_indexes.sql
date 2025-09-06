CREATE EXTENSION IF NOT EXISTS citext;
CREATE INDEX IF NOT EXISTS idx_refresh_user ON refresh_tokens(user_id);
CREATE INDEX IF NOT EXISTS idx_refresh_expires ON refresh_tokens(expires_at);
CREATE INDEX IF NOT EXISTS idx_reset_user   ON password_reset_tokens(user_id);
CREATE INDEX IF NOT EXISTS idx_reset_expires ON password_reset_tokens(expires_at);