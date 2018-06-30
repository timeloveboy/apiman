FROM timeloveboy/moedocker
MAINTAINER timeloveboy(734991033@qq.com)
ADD apiman /usr/local/bin/
VOLUME ["/web"]
EXPOSE 8080
CMD apiman -port=8080 -root=/web