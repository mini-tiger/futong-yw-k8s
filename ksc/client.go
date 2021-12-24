package ksc

import (
	"context"
	"fmt"
	"sync"
	"time"

	"ftk8s/base/cfg"
	"ftk8s/model"
	"ftk8s/storage"
	"ftk8s/util"

	jsoniter "github.com/json-iterator/go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	clientcmdapilatest "k8s.io/client-go/tools/clientcmd/api/latest"
	clientcmdapiv1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

const (
	// High enough QPS to fit all expected use cases.
	defaultQPS = 1e6
	// High enough Burst to fit all expected use cases.
	defaultBurst = 1e6
	// full resyc cache resource time
	defaultResyncPeriod = 30 * time.Second
)

var (
	ClusterManagerMap = &sync.Map{}
)

type ClusterManager struct {
	Cluster *model.Cluster
	Config  *rest.Config
	Client  IResource
}

func BuildApiServerClient() {
	clusters, err := storage.ListAllCluster()
	if err != nil {
		cfg.Mlog.Error("failed to list all cluster, error message: ", err.Error())
		return
	}

	changed := clusterChanged(clusters)
	if changed {
		cfg.Mlog.Info("cluster changed, so start to sync cluster...")

		shouldRemoveCluster(clusters)

		// build new ClusterManagerMap
		for i := 0; i < len(clusters); i++ {
			cluster := clusters[i]
			if cluster.ClusterAPI == "" {
				cfg.Mlog.Errorf("cluster(%d) ClusterAPI is empty", cluster.ID)
				continue
			}
			clientSet, config, err := BuildClient(cluster.ClusterAPI, cluster.K8sConfig)
			if err != nil {
				cfg.Mlog.Errorf("failed to build cluster(%d) client, error message: %v", cluster.ID, err.Error())
				continue
			}

			ksCache, err := buildKsCache(clientSet)
			if err != nil {
				cfg.Mlog.Errorf("clusterID(%d) failed to buildKsCache, error message: %v", cluster.ID, err.Error())
				continue
			}

			clusterManagerObj := &ClusterManager{
				Config:  config,
				Cluster: &cluster,
				Client:  NewResourceHandler(clientSet, ksCache),
			}

			clusterManagerInterface, ok := ClusterManagerMap.Load(cluster.ID)
			if ok {
				clusterManagerObj := clusterManagerInterface.(*ClusterManager)
				clusterManagerObj.Close()
			}

			ClusterManagerMap.Store(cluster.ID, clusterManagerObj)
		}
		cfg.Mlog.Info("Successfully to sync cluster")
	}
}

// shouldRemoveCluster deal with deleted cluster in ClusterManagerMap
func shouldRemoveCluster(changedClusters []model.Cluster) {
	changedClusterMap := make(map[int]struct{})
	for _, cluster := range changedClusters {
		changedClusterMap[cluster.ID] = struct{}{}
	}

	ClusterManagerMap.Range(func(key, value interface{}) bool {
		if _, ok := changedClusterMap[key.(int)]; !ok {
			clusterManagerInterface, _ := ClusterManagerMap.Load(key)
			clusterManagerObj := clusterManagerInterface.(*ClusterManager)
			clusterManagerObj.Close()
			ClusterManagerMap.Delete(key)
		}
		return true
	})
}

func clusterChanged(clusters []model.Cluster) bool {
	if util.GetLenOfSyncMap(ClusterManagerMap) != len(clusters) {
		cfg.Mlog.Infof("clusters length (%d) changed to (%d)", util.GetLenOfSyncMap(ClusterManagerMap), len(clusters))
		return true
	}

	for _, cluster := range clusters {
		clusterManagerInterface, ok := ClusterManagerMap.Load(cluster.ID)
		if !ok {
			return true
		}
		clusterManagerObj := clusterManagerInterface.(*ClusterManager)

		// 只要关心核心字段是否变化即可，其他字段不影响对k8s的apiserver发起请求
		if clusterManagerObj.Cluster.ClusterAPI != cluster.ClusterAPI {
			cfg.Mlog.Infof("cluster apiserver (%s) changed to (%s)", clusterManagerObj.Cluster.ClusterAPI, cluster.ClusterAPI)
			return true
		}
		if clusterManagerObj.Cluster.K8sConfig != cluster.K8sConfig {
			cfg.Mlog.Infof("cluster K8sConfig (%s) changed to (%s).", clusterManagerObj.Cluster.K8sConfig, cluster.K8sConfig)
			return true
		}
	}

	return false
}

func GetClusterManager(clusterID int) (*ClusterManager, error) {
	clusterManagerInterface, exist := ClusterManagerMap.Load(clusterID)

	if !exist {
		BuildApiServerClient()
		clusterManagerInterface, exist = ClusterManagerMap.Load(clusterID)
		if !exist {
			return nil, fmt.Errorf("clusterID(%d) not exist", clusterID)
		}
	}

	clusterManagerObj := clusterManagerInterface.(*ClusterManager)

	return clusterManagerObj, nil
}

func BuildClient(clusterAPI string, k8sConfig string) (*kubernetes.Clientset, *rest.Config, error) {
	configV1 := clientcmdapiv1.Config{}
	err := jsoniter.Unmarshal([]byte(k8sConfig), &configV1)
	if err != nil {
		cfg.Mlog.Error("failed to jsoniter unmarshal k8sConfig, error message: ", err.Error())
		return nil, nil, err
	}
	configObject, err := clientcmdapilatest.Scheme.ConvertToVersion(&configV1, clientcmdapi.SchemeGroupVersion)
	configInternal := configObject.(*clientcmdapi.Config)

	clientConfig, err := clientcmd.NewDefaultClientConfig(*configInternal, &clientcmd.ConfigOverrides{
		ClusterDefaults: clientcmdapi.Cluster{Server: clusterAPI},
	}).ClientConfig()

	if err != nil {
		cfg.Mlog.Error("failed to build client config, error message: ", err.Error())
		return nil, nil, err
	}

	clientConfig.QPS = defaultQPS
	clientConfig.Burst = defaultBurst

	clientSet, err := kubernetes.NewForConfig(clientConfig)

	if err != nil {
		cfg.Mlog.Errorf("failed to kubernetes.NewForConfig(%v) by clusterAPI(%s), error message: %v", clientConfig, clusterAPI, err.Error())
		return nil, nil, err
	}

	return clientSet, clientConfig, nil
}

func GetIResource(clusterID int) (IResource, error) {
	clusterManagerObj, err := GetClusterManager(clusterID)
	if err != nil {
		return nil, err
	}
	return clusterManagerObj.Client, err
}

func GetK8sClientSet(clusterID int) (*kubernetes.Clientset, error) {
	iResource, err := GetIResource(clusterID)
	if err != nil {
		return nil, err
	}
	resourceHandlerObj := iResource.(*resourceHandler)
	return resourceHandlerObj.client, err
}

func GetKsCache(clusterID int) (*KsCache, error) {
	iResource, err := GetIResource(clusterID)
	if err != nil {
		return nil, err
	}
	resourceHandlerObj := iResource.(*resourceHandler)
	return resourceHandlerObj.ksCache, err
}

func CheckConnectCluster(clusterAPI string, k8sConfig string) error {
	k8sCli, _, err := BuildClient(clusterAPI, k8sConfig)
	if err != nil {
		return err
	}

	_, err = k8sCli.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	return err
}
