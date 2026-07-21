# Base
FROM node:22-alpine3.20 AS base
WORKDIR /app
RUN corepack enable
# esbuild's build script is not needed (pnpm installs the platform binary as an
# optional dep), so don't fail the install over unapproved build scripts.
# Set as env because "pnpm build" re-checks the deps status before running.
ENV PNPM_CONFIG_STRICT_DEP_BUILDS=false
COPY ./mc-frontend/package.json ./mc-frontend/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
COPY ./mc-frontend/ ./
RUN pnpm build

# Binary
FROM nginx:1.27-alpine AS binary
COPY --from=base /app/dist /usr/share/nginx/html
COPY ./mc-frontend/nginx.conf /etc/nginx/conf.d/default.conf
