package apis

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *App) ToDeployment() *appsv1.Deployment {
	return &appsv1.Deployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      a.Name,
			Namespace: a.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &a.Replicas,
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{
					"app": a.Name,
					"namespace": a.Namespace,
				},
			},
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  a.Name,
							Image: a.Image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: a.Port,
								},
							},
						},
					},
				},
			},
		},
	}

}

func (a *App) ToService() *corev1.Service {
	return &corev1.Service{
		TypeMeta: v1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: a.Name,
			Namespace: a.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Port: a.Port,
					Name: "http",
				},
			},
			Selector: map[string]string{
				"app": a.Name,
				"namespace": a.Namespace,
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}
}
