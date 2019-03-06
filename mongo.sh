#!/bin/bash
docker run -d --rm -p 27017:27017 -v ~/rb_server_db:/data/db mongo:4.0