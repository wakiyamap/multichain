FROM ubuntu:xenial

RUN apt-get update --fix-missing && apt-get install --yes software-properties-common wget

RUN wget -c https://github.com/monacoinproject/monacoin/releases/download/monacoin-0.17.1/monacoin-0.17.1-x86_64-linux-gnu.tar.gz -O - | tar xz
RUN mv ./monacoin-0.17.1 /app
RUN chmod +x /app/bin/monacoind
RUN chmod +x /app/bin/monacoin-cli

COPY dogecoin.conf /root/.monacoin/monacoin.conf
COPY run.sh /root/run.sh
RUN chmod +x /root/run.sh

EXPOSE 19402

ENTRYPOINT ["./root/run.sh"]
