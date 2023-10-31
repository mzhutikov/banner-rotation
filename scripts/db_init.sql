CREATE TABLE slots (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL
);

CREATE TABLE banners (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL
);

CREATE TABLE user_groups (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL
);

CREATE TABLE banner_rotation (
    slot_id INTEGER REFERENCES slots(id),
    banner_id INTEGER REFERENCES banners(id),
    impressions INTEGER DEFAULT 0,
    clicks INTEGER DEFAULT 0,
    PRIMARY KEY (slot_id, banner_id)
);

CREATE TABLE statistics (
    id SERIAL PRIMARY KEY,
    event_type VARCHAR(10) CHECK(event_type IN ('click', 'impression')),
    slot_id INTEGER REFERENCES slots(id),
    banner_id INTEGER REFERENCES banners(id),
    user_group_id INTEGER REFERENCES user_groups(id),
    event_timestamp TIMESTAMP NOT NULL
);
