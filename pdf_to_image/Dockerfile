FROM ubuntu:19.10
# Setup locale. This prevents Python 3 IO encoding issues.
ENV LANG C.UTF-8
ENV PYTHONUNBUFFERED 1

# installing SO deps
RUN apt-get update && apt-get install -yq  python3 python3-pip poppler-utils
COPY ./requirements.txt /opt/requirements.txt
RUN pip3 install -r /opt/requirements.txt

# Creating the user
ENV USER microservice
RUN groupadd $USER && useradd -g $USER $USER -s /bin/bash


# Copying the content to the image
ENV NAME pdf_to_image
COPY ./ /opt/${NAME}
RUN chown ${USER}:${USER} /opt/${NAME}
WORKDIR /opt/${NAME}

# Installing python dependencies
EXPOSE 50052
USER ${USER}
ENV PYTHONPATH /opt/${NAME}
CMD ["/usr/bin/python3", "app/server.py"]