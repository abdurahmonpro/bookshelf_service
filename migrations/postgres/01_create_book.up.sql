CREATE TABLE IF NOT EXISTS "book" (
    "id" UUID PRIMARY KEY,
    "isbn" VARCHAR(30) NOT NULL,
    "title" VARCHAR(30) NOT NULL,
    "cover" VARCHAR(17) NOT NULL UNIQUE,
    "author" VARCHAR(20) NOT NULL,
    "published" VARCHAR NOT NULL,
    "pages" INTEGER NOT NULL,
    -- 0 new, -- 1 reading, -- 2 finished
    "status" SMALLINT NOT NULL DEFAULT 0,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
); 