
FROM goharbor/harbor-db-base:dev

VOLUME /var/lib/postgresql/data
#
# COPY ./make/photon/db/docker-entrypoint.sh /docker-entrypoint.sh
# COPY ./make/photon/db/initdb.sh /initdb.sh
# COPY ./make/photon/db/upgrade.sh /upgrade.sh
# COPY ./make/photon/db/docker-healthcheck.sh /docker-healthcheck.sh
# COPY ./make/photon/db/initial-notaryserver.sql /docker-entrypoint-initdb.d/
# COPY ./make/photon/db/initial-notarysigner.sql /docker-entrypoint-initdb.d/
# COPY ./make/photon/db/initial-registry.sql /docker-entrypoint-initdb.d/
#

COPY ./docker-entrypoint.sh /docker-entrypoint.sh
COPY ./initdb.sh /initdb.sh
COPY ./upgrade.sh /upgrade.sh
COPY ./docker-healthcheck.sh /docker-healthcheck.sh
COPY ./initial-notaryserver.sql /docker-entrypoint-initdb.d/
COPY ./initial-notarysigner.sql /docker-entrypoint-initdb.d/
COPY ./initial-registry.sql /docker-entrypoint-initdb.d/
RUN chown -R postgres:postgres /docker-entrypoint.sh /docker-healthcheck.sh /docker-entrypoint-initdb.d \
    && chmod u+x /docker-entrypoint.sh /docker-healthcheck.sh

ENTRYPOINT ["/docker-entrypoint.sh", "96", "13"]
HEALTHCHECK CMD ["/docker-healthcheck.sh"]

USER postgres
