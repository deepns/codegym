import sys
print("Hello, World")
print(f'Running Python version {sys.version}')
# To run this script in a python container without building a docker image,
# docker run --rm --volume $PWD:/home python:latest python /home/app.py

