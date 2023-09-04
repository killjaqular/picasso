-- Create users and assign privileges
-- This role is meant to be used by the server
DO
$$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'picasso') THEN
        CREATE ROLE picasso WITH LOGIN PASSWORD 'picasso';
    END IF;
END;
$$;
-- This role is meant to be used for troubleshooting and maintenance by directly
-- connecting to the database at its host:port
DO
$$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'adminuser') THEN
        CREATE ROLE adminuser WITH LOGIN PASSWORD 'adminpass';
    END IF;
END;
$$;

-- Grant permissions on the database
GRANT USAGE ON SCHEMA accounts TO picasso;
GRANT USAGE ON SCHEMA images TO picasso;
GRANT USAGE ON SCHEMA sessionKeys TO picasso;
GRANT USAGE ON SCHEMA uuids TO picasso;

GRANT USAGE ON SCHEMA accounts TO adminuser;
GRANT USAGE ON SCHEMA images TO adminuser;
GRANT USAGE ON SCHEMA sessionKeys TO adminuser;
GRANT USAGE ON SCHEMA uuids TO adminuser;

-- Grant permissions on tables
GRANT INSERT, SELECT ON accounts.user TO picasso;
GRANT DELETE, INSERT, SELECT ON images.image TO picasso;
GRANT DELETE, INSERT, SELECT ON sessionKeys.session TO picasso;
GRANT DELETE, INSERT, SELECT ON uuids.uuid TO picasso;

GRANT DELETE, INSERT, SELECT, UPDATE ON accounts.user TO adminuser;
GRANT DELETE, INSERT, SELECT, UPDATE ON images.image TO adminuser;
GRANT DELETE, INSERT, SELECT, UPDATE ON sessionKeys.session TO adminuser;
GRANT DELETE, INSERT, SELECT, UPDATE ON uuids.uuid TO adminuser;
