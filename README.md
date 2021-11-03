# Synpse Enterprise 

Synpse Enterprise is self hosted version of Synpse intended to run in your own Datacenter.
To run Synpse self-hosted version, you will need to acquire Synpse license. You can get testing license 
by reaching to use on Discord or via form (TBD)

## Pre-requisites

1. DNS domain name. Example `synpse.yourcorp.com`
2. Kubernetes cluster OR single server with docker & docker-compose installed.
3. ssh-key pair for SSH proxy. We provide generator in this repository.

## Deployments

### docker-compose

docker-compose deployment models allows you to run all-in-one Synpse Enterprise deployment on single server.
It should not be used in production deployment, as it does not give you full HA capabilities.

> For production use Kubernetes deployment

1. Configure DNS to point to your server. This is how it look on our CloudFlare account

![CloudFlare](assets/cloudflare.pnk?raw=true "CloudFlare")

2. Create `.env` file using `.env.example` as blueprint.

3. SSH into target server and clone repository:

```
git clone https://github.com/synpse-hq/enterprise 
cd enterprise
```

4. Generate/Add ssh key into assets/secrets folder.

```
# generate new key pair
make generate-ssh-key
```

5. Copy `.env` file into root of the project

```
ls -la .env*
[root@unknown enterprise]$ ls -la .env*
-rw-rw-r--. 1 root root 1181 Nov  2 19:02 .env
-rw-rw-r--. 1 root root 1955 Nov  3 10:26 .env.example
```

6. Make sure DNS is already propagated. We run caddy, which handles TLS via Let's Encrypt.

```
# both IP addresses should match
dig synpse.yourcorp.com
curl ifconfig.me
```

7. Deploy Synpse

```
docker-compose up 
# for daemon mode run:
docker-compose up -d
```

### Kubernetes HELM

Soon.
