#!/bin/sh

########################################################################
# Special Certificate Generation Command for Google IoT Cloud
#
# openssl req -x509 -newkey rsa:2048 -keyout rsa_private.pem -nodes \
#    -out rsa_cert.pem -subj "/CN=unused"
########################################################################

openssl req -x509 -newkey rsa:2048 -keyout rsa_private.pem -nodes \
    -out rsa_cert.pem -subj "/CN=unused"
