CREATE TABLE IF NOT EXISTS public.room_users
(
    user_id uuid NOT NULL,
    room_id uuid NOT NULL,
    CONSTRAINT room_key FOREIGN KEY (room_id)
        REFERENCES public.rooms (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);