# Fun with RabbitMQ

[#59 - Fun with RabbitMQ (screencast)](https://sysadmincasts.com/episodes/59-fun-with-rabbitmq)

## Usage

Launch RabbitMQ in Docker

```
$ docker run -d -p 5671:5671 -p 5672:5672 -p 15672:15672 rabbitmq:3.6.6-management
Unable to find image 'rabbitmq:3.6.6-management' locally
3.6.6-management: Pulling from library/rabbitmq
693502eb7dfb: Pull complete
7eb18686cc46: Pull complete
ae00e0021d4f: Pull complete
9c40b0d0b2a9: Pull complete
2fb8e146207a: Pull complete
3d218990416b: Pull complete
e6a2b2fe78c0: Pull complete
fe9045f6bf09: Pull complete
2811fbe50640: Pull complete
34de23e5443b: Pull complete
57fc62f25d65: Pull complete
130407bb1e30: Pull complete
861305534fee: Pull complete
166ee531bc38: Pull complete
bdc4fbb675c6: Pull complete
Digest: sha256:60fa2366e203f1515b99082c4a2f3dcb157d836eacedb31dfdeb97d3fd9dd1ee
Status: Downloaded newer image for rabbitmq:3.6.6-management
40e6b32c0d721cf673b8a2afb94cd1e1aea87043e95f24df3ff2ee131827bdf3
```

Connect to the RabbitMQ Management Interface at http://localhost:15672/.

Run through `01_hello` and `02_wordlist` by launching the producer and consumers.
