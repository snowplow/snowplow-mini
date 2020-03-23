CREATE USER snowplow WITH PASSWORD 'snowplow';
CREATE DATABASE iglu OWNER snowplow;
\connect iglu snowplow;
CREATE TYPE key_action AS ENUM ('CREATE', 'DELETE');
CREATE TYPE schema_action AS ENUM ('READ', 'BUMP', 'CREATE', 'CREATE_VENDOR');
CREATE TABLE iglu_permissions (
    apikey              UUID            NOT NULL,
    vendor              VARCHAR(128),
    wildcard            BOOL            NOT NULL,
    schema_action       schema_action,
    key_action          key_action[]    NOT NULL,
    PRIMARY KEY (apikey)
);
CREATE TABLE iglu_schemas (
    vendor      VARCHAR(128)  NOT NULL,
    name        VARCHAR(128)  NOT NULL,
    format      VARCHAR(128)  NOT NULL,
    model       INTEGER       NOT NULL,
    revision    INTEGER       NOT NULL,
    addition    INTEGER       NOT NULL,

    created_at  TIMESTAMP     NOT NULL,
    updated_at  TIMESTAMP     NOT NULL,
    is_public   BOOLEAN       NOT NULL,

    body        JSON          NOT NULL
);
CREATE TABLE iglu_drafts (
    vendor      VARCHAR(128) NOT NULL,
    name        VARCHAR(128) NOT NULL,
    format      VARCHAR(128) NOT NULL,
    version     INTEGER      NOT NULL,

    created_at  TIMESTAMP    NOT NULL,
    updated_at  TIMESTAMP    NOT NULL,
    is_public   BOOLEAN      NOT NULL,

    body        JSON         NOT NULL
);
