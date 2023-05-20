FROM python:3.10
# path to the config file
ENV TEST_CONFIG="/config.json"

WORKDIR /usr/src/app
COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .
ENTRYPOINT [ "python3", "-m", "unittest" ]