docker run -d \
-p 443:443 \
--name zipsvr \
-v /c/Users/dejik/Desktop/go/src/github.com/nhvtran/info344-in-class/zipsvr/tls:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
nhvtran/zipsvr