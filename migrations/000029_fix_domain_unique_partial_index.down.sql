-- 回滚：恢复原始全量唯一约束
DROP INDEX IF EXISTS domains_full_domain_active_key;

ALTER TABLE domains ADD CONSTRAINT domains_full_domain_key UNIQUE (full_domain);
