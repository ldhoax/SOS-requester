CREATE TABLE
    requests (
        id UUID PRIMARY KEY,
        requester_id UUID NOT NULL,
        phone_number TEXT,
        email TEXT,
        short_description TEXT NOT NULL,
        location TEXT NOT NULL,
        latitude NUMERIC(10, 8),
        longitude NUMERIC(11, 8),
        description TEXT,
        emergency_level SMALLINT,
        created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        updated_at TIMESTAMP(0) WITHOUT TIME ZONE,
        deleted_at TIMESTAMP(0) WITHOUT TIME ZONE,
        FOREIGN KEY (requester_id) REFERENCES requesters(id)
    );