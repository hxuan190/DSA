Docker Overview

Docker is a platform designed to simplify the development, shipping, and running of applications by using containerization. Containers package software with all its dependencies, ensuring consistency across multiple environments.
Key Concepts

    Container: A lightweight, standalone, executable package that includes everything needed to run a piece of software, including code, runtime, libraries, and system tools.

    Image: A lightweight, immutable file that contains the source code, libraries, and dependencies needed to run the application. Containers are created from Docker images.

    Dockerfile: A text file that contains instructions on how to build a Docker image.

    Docker Hub: A public repository where you can find and share container images.
Basic Docker Commands
Build an Image: docker build -t <image_name> .

Run a Container: docker run -d --name <container_name> <image_name>

List Running Containers: docker ps

Stop a Running Container: docker stop <container_id>

Remove a Container: docker rm <container_id>

Pull an Image from Docker Hub: docker pull <image_name>

