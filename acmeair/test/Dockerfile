FROM python:3.9

# path to the config file
ENV TEST_CONFIG="/config.json"

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .

ENTRYPOINT [ "python3", "-m", "unittest" ]