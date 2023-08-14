# Database

### Schemas:

<ol>
<li>accounts</li>
<li>images</li>
<li>sessionKeys</li>
<li>uuids</li>
</ol>

### Tables:

```SQL
CREATE TABLE IF NOT EXISTS uuids.uuid (
    uuid UUID PRIMARY KEY,
    parentTable VARCHAR(64) -- The table that contains the actual asset of the UUID
);

CREATE TABLE IF NOT EXISTS accounts.user (
    uuid UUID PRIMARY KEY REFERENCES uuids.uuid(uuid),
    username VARCHAR(16) UNIQUE, -- No repeated usernames allowed
    password VARCHAR(256) -- SHA256
);

-- It is recommended images be stored on local file disk versus database
-- We will store metadata here, like the path on the local file system
CREATE TABLE IF NOT EXISTS images.image (
    uuid UUID PRIMARY KEY REFERENCES uuids.uuid(uuid),
    filePath VARCHAR(255), -- Maximum directory length in the ext4 file system, common in Ubuntu
    userUUID UUID REFERENCES accounts.user(uuid) -- Foreign key referencing user UUID
);

CREATE TABLE IF NOT EXISTS sessionKeys.session (
    uuid UUID PRIMARY KEY REFERENCES uuids.uuid(uuid),
    sessionKey VARCHAR(256), -- The session key of a user
    userUUID UUID REFERENCES accounts.user(uuid) -- Foreign key referencing user UUID
    expiration  TIMESTAMP WITH TIME ZONE -- Time sessionKey will expire and user must re-login
);
```