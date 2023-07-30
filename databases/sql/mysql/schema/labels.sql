CREATE TABLE labels (
    uuid VARCHAR(36) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE labels_data (
    uuid VARCHAR(36) NOT NULL,
    customer VARCHAR(32) NOT NULL,
    family VARCHAR(8) NOT NULL,
    model VARCHAR(16) NOT NULL,
    part_number VARCHAR(16) NOT NULL,
    station VARCHAR(32) NOT NULL,
    label TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (uuid) REFERENCES labels(uuid)
);

CREATE TABLE labels_deletes (
    uuid VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (uuid) REFERENCES labels(uuid)
);
