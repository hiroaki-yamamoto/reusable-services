FROM python
ENV PYTHONUNBUFFERED=1
RUN mkdir -p /opt/code
VOLUME [ "/opt/code" ]
WORKDIR /opt/code
ENTRYPOINT [ "./test.sh" ]
