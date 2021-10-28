#!/bin/bash
#sudo chown $USER -R ./database/data && sudo chmod 755 -R ./database/data \
docker-compose up -d --build --force-recreate
