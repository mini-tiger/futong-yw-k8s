package ksc

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/listers/apps/v1"
	autoscalingv1 "k8s.io/client-go/listers/autoscaling/v1"
	corev1 "k8s.io/client-go/listers/core/v1"
)

type KsCache struct {
	stopChan              chan struct{}
	sharedInformerFactory informers.SharedInformerFactory
}

func (c ClusterManager) Close() {
	resourceHandlerObj := c.Client.(*resourceHandler)
	close(resourceHandlerObj.ksCache.stopChan)
}

func buildKsCache(client *kubernetes.Clientset) (*KsCache, error) {
	stopChan := make(chan struct{})
	sharedInformerFactory := informers.NewSharedInformerFactory(client, defaultResyncPeriod)

	resourceMap := GetResourceMap()

	for _, value := range resourceMap {

		genericInformer, err := sharedInformerFactory.ForResource(value.GroupVersionResourceKind.GroupVersionResource)
		if err != nil {
			return nil, err
		}
		go genericInformer.Informer().Run(stopChan)
	}

	sharedInformerFactory.Start(stopChan)

	return &KsCache{
		stopChan:              stopChan,
		sharedInformerFactory: sharedInformerFactory,
	}, nil
}

func (c *KsCache) PodLister() corev1.PodLister {
	return c.sharedInformerFactory.Core().V1().Pods().Lister()
}

func (c *KsCache) EventLister() corev1.EventLister {
	return c.sharedInformerFactory.Core().V1().Events().Lister()
}

func (c *KsCache) DeploymentLister() appsv1.DeploymentLister {
	return c.sharedInformerFactory.Apps().V1().Deployments().Lister()
}

func (c *KsCache) NodeLister() corev1.NodeLister {
	return c.sharedInformerFactory.Core().V1().Nodes().Lister()
}

func (c *KsCache) EndpointLister() corev1.EndpointsLister {
	return c.sharedInformerFactory.Core().V1().Endpoints().Lister()
}

func (c *KsCache) HPALister() autoscalingv1.HorizontalPodAutoscalerLister {
	return c.sharedInformerFactory.Autoscaling().V1().HorizontalPodAutoscalers().Lister()
}
