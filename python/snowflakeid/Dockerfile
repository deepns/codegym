# Use an official Python runtime as the base image
FROM python:3.10-slim

# Set the working directory in the container
WORKDIR /app

# Copy application code
COPY app.py /app

# Install dependencies
RUN pip install pymongo

# Define environment variables (can be overridden in Kubernetes)
ENV MONGO_URI=mongodb://mongo:27017/
ENV MACHINE_ID=0

# Run the application
CMD ["python", "app.py"]
