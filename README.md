# Imersão Full Stack & FullCycle - Codelivery

![Imersão Full Stack && Full Cycle](https://events-fullcycle.s3.amazonaws.com/events-fullcycle/static/site/img/grupo_4417.png)

## Description

Backend repository made with Golang

**Important**: The Apache Kafka application must be running first.

## To set up /etc/hosts

Communication between applications takes place directly through the machine's network.

For this it is necessary to configure an address that all Docker containers can access.

Add to your /etc/hosts (for Windows the path is C:\Windows\system32\drivers\etc\hosts):

```
127.0.0.1 host.docker.internal
```

In all operating systems, it is necessary to open the program to edit _hosts_ as Administrator of the machine or root.

## Run the application

Run the commands:

```
docker-compose up -d
# Enter the container
docker-compose exec app bash
# Run the application Golang
go run main.go
```

### For Windows

Remember to install WSL2 and Docker. I see the video:: [https://www.youtube.com/watch?v=usF0rYCcj-E](https://www.youtube.com/watch?v=usF0rYCcj-E)

Follow the quick installation guide: [https://github.com/codeedu/wsl2-docker-quickstart](https://github.com/codeedu/wsl2-docker-quickstart)
