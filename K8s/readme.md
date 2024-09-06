1. Cluster

   Cluster là nhóm các máy chủ (nodes) mà Kubernetes quản lý. Mỗi cluster gồm ít nhất một master node (điều khiển) và nhiều worker nodes (nơi các ứng dụng chạy).

2. Node

   Node là một máy chủ vật lý hoặc ảo trong cluster Kubernetes, nơi các ứng dụng được chạy. Có hai loại node:
   Master Node: Quản lý toàn bộ cluster, bao gồm lịch trình Pods, cập nhật trạng thái cluster, và xử lý các hoạt động như scale, load balancing.
   Worker Node: Nơi chứa các container ứng dụng trong các Pods.

3. Pod

   Pod là đơn vị triển khai nhỏ nhất trong Kubernetes. Mỗi Pod chứa một hoặc nhiều container và chia sẻ chung mạng, không gian lưu trữ.
   Pod thường là wrapper xung quanh một hoặc nhiều container, thường được sử dụng để chạy các ứng dụng đơn lẻ hoặc liên quan.

4. Service

   Service cung cấp một cách tiếp cận để truy cập vào các Pod, giúp quản lý việc load balancing và duy trì IP tĩnh cho các Pod. Một Service có thể:
   Định tuyến các yêu cầu đến một nhóm Pod cụ thể.
   Cung cấp một điểm truy cập cố định cho các Pod, ngay cả khi Pods bị xóa hoặc tái triển khai.

5. Deployment

   Deployment là cách quản lý và kiểm soát việc triển khai ứng dụng trong Kubernetes. Nó giúp kiểm soát số lượng bản sao (replicas) của Pod, triển khai bản cập nhật mới và đảm bảo rằng ứng dụng luôn ở trạng thái hoạt động.

6. Namespace

   Namespace cung cấp một cơ chế để chia cluster Kubernetes thành nhiều không gian logic độc lập, giúp dễ dàng quản lý tài nguyên và môi trường khác nhau.

7. ConfigMap và Secret

   ConfigMap: Lưu trữ các cấu hình không bí mật, có thể truyền vào các container dưới dạng biến môi trường hoặc file cấu hình.
   Secret: Lưu trữ thông tin nhạy cảm như mật khẩu, khóa API và cung cấp chúng một cách an toàn cho Pods.

8. Ingress

   Ingress giúp cấu hình các quy tắc định tuyến HTTP và HTTPS để cung cấp truy cập bên ngoài vào các dịch vụ trong cluster Kubernetes.

9. Volume

   Volume là cơ chế lưu trữ dữ liệu cho container, giúp giữ lại dữ liệu ngay cả khi container dừng. Kubernetes cung cấp nhiều loại volume như emptyDir, hostPath, PersistentVolume, và PersistentVolumeClaim.

10. Kubelet

    Kubelet là một thành phần chạy trên mỗi worker node, chịu trách nhiệm giao tiếp giữa master node và worker node, đảm bảo rằng các container trong Pod đang chạy.

Cách bắt đầu với Kubernetes:

    Cài đặt Minikube: Công cụ này giúp bạn tạo ra một cluster Kubernetes đơn giản để học và thử nghiệm trên máy cá nhân.
    Sử dụng kubectl: Đây là công cụ dòng lệnh để quản lý cluster Kubernetes.
        Các lệnh cơ bản:
            kubectl get pods: Liệt kê các Pod đang chạy.
            kubectl apply -f <file.yaml>: Áp dụng file cấu hình YAML để tạo các tài nguyên trong Kubernetes.
            kubectl describe pod <pod_name>: Xem chi tiết về một Pod.