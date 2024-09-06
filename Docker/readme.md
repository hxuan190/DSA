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

Khi làm việc với Docker, đặc biệt là khi chạy các container, bạn cần hiểu rõ về khái niệm port mapping giữa cổng Docker (cổng bên trong container) và cổng Local (cổng trên máy host). Điều này rất quan trọng để ứng dụng trong container có thể giao tiếp với thế giới bên ngoài.
1. Cổng Local (Host Port)

   Đây là cổng trên máy host (máy mà bạn đang sử dụng để chạy container). Đây là nơi mà người dùng hoặc các dịch vụ khác trên máy tính của bạn có thể truy cập ứng dụng trong container.
   Ví dụ: Trên máy tính của bạn, bạn có thể chạy một ứng dụng ở cổng localhost:8080.

2. Cổng Docker (Container Port)

   Đây là cổng mà ứng dụng bên trong container đang chạy. Mỗi container có môi trường mạng riêng và không thể tự động truy cập trực tiếp vào các dịch vụ trên máy host hoặc các container khác mà không có cấu hình mạng phù hợp.
   Ví dụ: Một container có thể chạy ứng dụng tại cổng 3000 bên trong container.

3. Port Mapping (Ánh xạ cổng)

Khi chạy container, bạn có thể chỉ định ánh xạ giữa cổng Docker và cổng Local bằng cách sử dụng tùy chọn -p trong lệnh docker run. Điều này cho phép chuyển tiếp lưu lượng từ cổng Local của máy chủ sang cổng của container bên trong.

