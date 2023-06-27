CREATE TABLE IF NOT EXISTS authv1.role
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    create_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    delete_at timestamp with time zone,
    CONSTRAINT id PRIMARY KEY (id),
    CONSTRAINT name UNIQUE (name)
);
CREATE TABLE IF NOT EXISTS authv1.user_role
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    role_id uuid NOT NULL,
    create_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    delete_at timestamp with time zone,
    CONSTRAINT user_role_id PRIMARY KEY (id),
    CONSTRAINT user_role_rel UNIQUE (user_id, role_id)
);