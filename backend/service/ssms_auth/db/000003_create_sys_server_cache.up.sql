CREATE TABLE IF NOT EXISTS sys.server_cache
(
    key character varying COLLATE pg_catalog."default" NOT NULL,
    value character varying COLLATE pg_catalog."default",
    expire_time timestamp with time zone,
    update_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    delete_at timestamp with time zone,
    CONSTRAINT server_cache_pkey PRIMARY KEY (key)
);
ALTER TABLE IF EXISTS sys.server_cache
    OWNER to clans;