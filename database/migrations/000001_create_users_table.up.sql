CREATE TABLE
    users (
        id SERIAL,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        email TEXT NOT NULL,
        status SMALLINT NOT NULL,
        created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(0) WITHOUT TIME ZONE,
        deleted_at TIMESTAMP(0) WITHOUT TIME ZONE,
        PRIMARY KEY (id)
    );