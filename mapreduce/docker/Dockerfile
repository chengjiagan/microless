FROM python:3.10-slim
# name of the service
ARG SERVICE
# path to the config file
ENV SERVICE_CONFIG="/config.json"
# opentelemetry settings
ENV OTEL_TRACES_EXPORTER="otlp_proto_http"
ENV OTEL_METRICS_EXPORTER="none"
ENV OTEL_EXPORTER_OTLP_ENDPOINT="http://jaeger:4318"
ENV OTEL_SERVICE_NAME="${SERVICE}"

WORKDIR /usr/src/app
COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .
ENTRYPOINT [ "opentelemetry-instrument", "python", "main.py" ]
