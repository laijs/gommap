From alpine

COPY gommap /usr/bin/
RUN dd bs=1M count=50 if=/dev/urandom of=/random50M
CMD gommap /random50M
