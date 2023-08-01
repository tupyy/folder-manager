BEGIN;

CREATE TABLE IF NOT EXISTS folders (
    id TEXT PRIMARY KEY,
    resource_id VARCHAR(50),
    name TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT (now() AT TIME ZONE 'UTC') NOT NULL,
    owner_id TEXT NOT NULL,
    bucket TEXT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS TABLE labels (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS folders_labels (
    folder_id TEXT REFERENCES folders(id) ON DELETE CASCADE,
    label_id TEXT REFERENCES labels(id) ON DELETE CASCADE,
    CONSTRAINT folders_labels_pk PRIMARY KEY (
        folder_id,
        label_id
    ) 
);

COMMIT;
