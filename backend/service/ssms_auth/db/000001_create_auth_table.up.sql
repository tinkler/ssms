CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS authv1."user"
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    username character varying COLLATE pg_catalog."default" NOT NULL,
    password character varying COLLATE pg_catalog."default" NOT NULL,
    create_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    delete_at timestamp with time zone,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT username UNIQUE (username)
)