FROM ghcr.io/u-yas/ennbu:0.0.3

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]