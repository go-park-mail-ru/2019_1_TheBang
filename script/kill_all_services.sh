#!bin/bash

kill -l $(lsof -i -P -n | grep main)