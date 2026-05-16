# OpenDomain

开源的免费二级域名分发平台，集成 PowerDNS，支持多种 DNS 记录类型管理。

## 功能

- 域名注册与管理（实时可用性查询、自动续期）
- DNS 记录管理（A、AAAA、CNAME、MX、TXT、NS、SRV、CAA）
- 第三方登录（GitHub、Google、NodeLoc）
- 优惠券系统（配额券、VIP 券、续期券）
- 邀请机制与奖励
- 公告系统
- 域名健康扫描（DNS、HTTP、SSL、安全检测）
- 自定义页面管理
- 管理后台（用户、域名、订单、系统配置）
- Telegram Bot 通知
- FOSSBilling 集成（可选）

## 技术栈

- **后端**：Go + Gin + PostgreSQL + Redis
- **前端**：Vue 3 + Vite
- **DNS**：PowerDNS（MariaDB 后端）
- **部署**：Docker Compose

## 部署

### 前置要求

- Docker & Docker Compose

### 一键部署

```bash
git clone https://github.com/nodeloc/opendomain.git
cd opendomain

cp .env.example .env
# 编辑 .env，至少修改以下变量：
# DB_PASSWORD、JWT_SECRET、POWERDNS_API_KEY

./docker-deploy.sh
```

如果需要同时启动内置 PowerDNS，使用 `powerdns` profile：

```bash
# 额外修改 .env 中的 PDNS_DB_PASSWORD、PDNS_ADMIN_SECRET
docker compose --profile powerdns up -d
```

部署完成后访问：

| 服务 | 地址 | 说明 |
|------|------|------|
| API 后端 | http://localhost:8000 | 默认启用 |
| PowerDNS Admin | http://localhost:9191 | 需启用 powerdns profile |
| PowerDNS API | http://localhost:8081 | 需启用 powerdns profile |

### 服务说明

`docker-compose.yml` 包含以下服务：

- `postgres` — OpenDomain 应用数据库
- `redis` — 缓存
- `migrate` — 自动执行数据库迁移
- `app` — OpenDomain 后端应用
- `pdns-db` — PowerDNS 专用 MariaDB（`powerdns` profile）
- `powerdns` — 权威 DNS 服务器，暴露 53 端口（`powerdns` profile）
- `pdns-admin` — PowerDNS Web 管理面板（`powerdns` profile）

### 常用命令

```bash
# 查看日志
docker compose logs -f

# 查看应用日志
docker compose logs -f app

# 重启服务
docker compose restart

# 停止并删除容器
docker compose down
```

### 注意事项

- 53 端口在 Ubuntu 上可能被 `systemd-resolved` 占用，可在 `.env` 中设置 `DNS_PORT=5353` 改用其他端口
- 首次访问 PowerDNS Admin（`:9191`）需注册管理员账号，API URL 填 `http://powerdns:8081`

## 许可证

[MIT](LICENSE)
