################################################################
# Builds a Docker image for the Picasso database.
################################################################
FROM postgres:latest

# Copy in configuration files
COPY *.sql /docker-entrypoint-initdb.d/
#COPY postgresql.conf /etc/postgresql.conf
# Copy in ssl files
#COPY ssl/ /var/lib/postgresql/ssl/
# Copy in utility scripts
COPY wait-for-it.sh /

# Set directory for the following RUN commands
#WORKDIR "/var/lib/postgresql/"
#RUN ["chown", "999", "./ssl/server.key"]
#RUN ["chmod", "0600", "./ssl/server.key"]

# Set directory for the following RUN commands
WORKDIR "/"
RUN ["chmod", "+x", "wait-for-it.sh"]
ENV PGDATA='/data/postgres-picasso'

# Checks if the container is up and running
HEALTHCHECK --interval=2s --timeout=60s --start-period=10s --retries=10 CMD bash /wait-for-it.sh localhost:5432
