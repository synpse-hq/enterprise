# Synpse Enterprise 

Synpse Enterprise is a self-hosted version of Synpse, intended to run in your own datacenter or private cloud.

## Pre-requisites

1. License. Trial licenses are available for free. Just reach out to us via email (hello@synpse.net) or via the [online form](https://synpse.net/contact/)
1. DNS domain name. Example `synpse.faros.sh`. We are going to use this domain in the example below. Replace it with your own.
2. Single server with docker & docker-compose installed (we will provide Kubernetes installer too).
3. SSH key pair for SSH proxy. We provide key generator in this repository.

## Deployments

### docker-compose

docker-compose deployment models allows you to run all-in-one Synpse Enterprise deployment on a single server. When used in production, ensure that you perform regular backups of your Postgres database. We recommend using one of the managed services such as [GCP CloudSQL](https://cloud.google.com/sql), [AWS RDS](https://aws.amazon.com/rds/).

1. Configure DNS to point to your server. This is how it look on our CloudFlare account:

![CloudFlare](assets/cloudflare.png?raw=true "CloudFlare")

2. Create `.env` file using `.env.example` as blueprint. You should NOT need to modify `docker-compose.yaml` file

3. SSH into target server and clone repository:

```
git clone https://github.com/synpse-hq/enterprise 
cd enterprise
```

4. Generate/Add ssh key into assets/secrets folder:

```
# generate new key pair
make generate-ssh-key
```

5. Copy `.env` file into root of the project:

```
cp .env.example .env
```

6. Edit the `.env` file with your values.

7. Make sure DNS is already propagated. We use [Caddy](https://github.com/caddyserver/caddy), which handles TLS via Let's Encrypt. Caddy requires port 443 to solve Let's Encrypt challenge.

```
# both IP addresses should match
dig synpse.faros.sh
curl ifconfig.me
```

8. Deploy Synpse

```
docker-compose up -d
```

8. Open `synpse.faros.sh:8443` and login using ADMIN password you set in `.env` file 

### Kubernetes HELM

Comming soon.
