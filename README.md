# k8s-operator-demo

# 目录

1. [使用 kuberbuilder 初始化项目](./docs/chapter01/01-kubebuilder-init-project.md)
2. [简单跑一跑](./docs/chapter01/02-simplest-redis-crd.md)
3. [发布 crd controller](./docs/chapter01/03-deploy-crd-controller.md)
4. [使用注解完整字段值约束](./docs/chapter02/04-filed-validation-by-comment.md)
5. [通过 webhook 进行字段验证](./docs/chapter02/05-filed-validation-by-webhook.md)
6. [使用 Operator 创建并发布一个 Pod](./docs/chapter02/06-create-pod-by-redis-operator.md)
7. K8S 父子资源删除管理
    1. [使用 OwnerReference 管理 redis operator 创建的 Pod](./docs/chapter02/07.1-delete-pod-by-redis-OwnerReference.md)
    2. [使用 finalizers 管理 redis operator 创建的 Pod](./docs/chapter02/07.2-delete-pod-by-finalizers.md)
8. [Pod 扩容与缩容](./docs/chapter02/08-scale-pod.md)
9. [监听 k8s 事件](./docs/chapter02/09-watch-k8s-event.md)
10. [重建被删除的 Pod](./docs/chapter02/10-recreate-deleted-pod.md)
11. [使用 controllerutil 优化代码](./docs/chapter02/11-official-package-optimize.md)
12. [增加 event 事件支持](./docs/chapter03/12-add-event.md)
13. [添加 Status 状态字段](./docs/chapter03/13-add-status.md)
13. [支持 kubectl scale 和 kubectl autoscale 命令](./docs/chapter03/14-kubectl-scale-autoscale-support.md)
