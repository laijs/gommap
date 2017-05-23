From alpine

COPY gommap /usr/bin/
RUN dd bs=50M count=1 if=/dev/urandom of=/random50M
CMD gommap /random50M
