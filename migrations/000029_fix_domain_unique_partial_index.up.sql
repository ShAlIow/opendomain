-- 将 full_domain 唯一约束替换为 partial unique index
-- 只对未软删除（deleted_at IS NULL）的记录生效
-- 这样软删除的域名不占唯一索引，可以被重新注册

ALTER TABLE domains DROP CONSTRAINT IF EXISTS domains_full_domain_key;

CREATE UNIQUE INDEX IF NOT EXISTS domains_full_domain_active_key
    ON domains (full_domain)
    WHERE deleted_at IS NULL;
