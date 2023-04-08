influx:
	$ docker run -d -p 8086:8086 \
		  --name influx \
          -e DOCKER_INFLUXDB_INIT_MODE=setup \
          -e DOCKER_INFLUXDB_INIT_USERNAME=mhthrh \
          -e DOCKER_INFLUXDB_INIT_PASSWORD=P@ssw0rd \
          -e DOCKER_INFLUXDB_INIT_ORG=my0rg \
          -e DOCKER_INFLUXDB_INIT_BUCKET=BlackBucket \
          -e DOCKER_INFLUXDB_INIT_RETENTION=1w \
          -e DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=my-super-secret-auth-token \
          influxdb:2.7.0