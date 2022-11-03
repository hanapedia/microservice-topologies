FROM --platform=linux/amd64 otel/opentelemetry-collector-contrib:0.45.0
COPY conf.yaml .
EXPOSE 1888
EXPOSE 8888
EXPOSE 8889
EXPOSE 13133
EXPOSE 4317
EXPOSE 55670
CMD [ "--config=conf.yaml" ]

