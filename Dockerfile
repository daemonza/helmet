FROM alpine:3.4
MAINTAINER Werner Gillmer <werner.gillmer@gmail.com>

# Add Helm
ADD https://storage.googleapis.com/kubernetes-helm/helm-v2.2.0-linux-amd64.tar.gz /opt/

WORKDIR /opt
RUN  tar zxvf helm-v2.2.0-linux-amd64.tar.gz
RUN mv linux-amd64/helm /usr/local/bin

# Add Helmet
ADD helmet /opt/helmet
RUN chmod +x /opt/helmet

CMD /opt/helmet
