CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.rooms
(
    id uuid DEFAULT uuid_generate_v4(),
    name character varying(32) NOT NULL,
    status smallint NOT NULL DEFAULT 1,
    created_at timestamp without time zone NOT NULL,
    created_by uuid NOT NULL,
    PRIMARY KEY (id),
);