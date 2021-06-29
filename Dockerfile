FROM acinq/eclair
VOLUME "./path_on_host:/data"
VOLUME "./env/bitcoin.conf:/.eclair/bitcoin.conf"
EXPOSE 29000