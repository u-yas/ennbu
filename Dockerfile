FROM ghcr.io/u-yas/ennbu:0.0.5

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]