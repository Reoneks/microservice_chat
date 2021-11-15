CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.messages
(
    id uuid DEFAULT uuid_generate_v4(),
    text text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    created_by uuid NOT NULL,
    room_id uuid NOT NULL,
    status smallint NOT NULL DEFAULT 1,
    PRIMARY KEY (id),
    CONSTRAINT user_key FOREIGN KEY (created_by)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID,
);