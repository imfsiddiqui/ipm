# Use the official Alpine Linux image
FROM alpine:latest

# Copy the contents of the dist directory to the root directory of the container
COPY dist/bin/ /ipm/

# Set the working directory to the root directory
WORKDIR /ipm/ipm-linux-amd64/

# Set the entrypoint to the executable
CMD ["/ipm/ipm-linux-amd64/ipm"]
