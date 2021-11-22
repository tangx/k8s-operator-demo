package helper

import (
	"context"
	"fmt"
	"time"

	appv1 "github.com/tangx/k8s-operator-demo/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CreateRedis 创建 redis pod
func CreateRedisPod(ctx context.Context, client client.Client, config *appv1.Redis) error {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%d", config.Name, time.Now().Unix()),
			Namespace: config.Namespace,
			// 创建 OwnerReference
			OwnerReferences: []metav1.OwnerReference{
				ownerReference(*config),
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "redis",
					Image: config.Spec.Image,
					Ports: []corev1.ContainerPort{
						{
							Name:          "redis-port",
							Protocol:      "TCP",
							ContainerPort: config.Spec.Port,
						},
					},
				},
			},
		},
	}

	// ctx := context.Background()
	return client.Create(ctx, pod)
}

// CreateRedis 创建 redis pod
func CreateRedisPod2(ctx context.Context, client client.Client, config *appv1.Redis) error {
	pod := &corev1.Pod{}
	pod.Name = config.Name
	pod.Namespace = config.Namespace
	pod.Spec.Containers = []corev1.Container{
		{
			Name:            config.Name,
			Image:           config.Spec.Image,
			ImagePullPolicy: corev1.PullIfNotPresent,
			Ports: []corev1.ContainerPort{
				{
					ContainerPort: config.Spec.Port,
				},
			},
		},
	}
	// ctx := context.Background()
	return client.Create(ctx, pod)
}

func ownerReference(config appv1.Redis) metav1.OwnerReference {
	return metav1.OwnerReference{
		APIVersion:         config.APIVersion,
		Kind:               config.Kind,
		Name:               config.Name,
		UID:                config.UID,
		Controller:         ptrBool(true),
		BlockOwnerDeletion: ptrBool(true),
	}
}

func ptrBool(b bool) *bool {
	return &b
}
