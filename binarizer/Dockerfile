FROM ubuntu:19.10
# Setup locale. This prevents Python 3 IO encoding issues.
ENV LANG C.UTF-8
ENV PYTHONUNBUFFERED 1

# installing SO deps
RUN apt-get update && apt-get install -yq  python3 python3-pip
COPY ./requirements.txt /opt/requirements.txt
RUN pip3 install -r /opt/requirements.txt

# Creating the user
ENV USER microservice
RUN groupadd $USER && useradd -g $USER $USER -s /bin/bash


# Copying the content to the image
ENV NAME binarizer
COPY ./ /opt/${NAME}
RUN chown ${USER}:${USER} /opt/${NAME}
WORKDIR /opt/${NAME}
EXPOSE 50053
USER ${USER}
ENV PYTHONPATH /opt/${NAME}
CMD ["/usr/bin/python3", "app/server.py"]