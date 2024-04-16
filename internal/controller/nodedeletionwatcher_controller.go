/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	//"sigs.k8s.io/controller-runtime/pkg/log"

	//nodecleanupcontrollerv1 "gitlab.arc.hcloud.io/ccp/hks/node-cleanup-controller/api/v1"

	//////hyunyoung added//////

	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	//"os/exec"

	"fmt"
	"sync"
	//"sigs.k8s.io/controller-runtime/pkg/reconcile"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	//"k8s.io/apimachinery/pkg/types"
	//"k8s.io/apimachinery/pkg/api/errors"
	//"k8s.io/utils/pointer"
)

var (
	cmdMutex sync.Mutex
	cmd      string
)

// NodeCleanUpReconciler reconciles a NodeCleanUp object
type NodeDeletionWatcher struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=nodecleanupcontroller.gitlab.arc.hcloud.io,resources=nodecleanups,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=nodecleanupcontroller.gitlab.arc.hcloud.io,resources=nodecleanups/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=nodecleanupcontroller.gitlab.arc.hcloud.io,resources=nodecleanups/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NodeCleanUp object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
func (r *NodeDeletionWatcher) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//log := log.FromContext(ctx)

	//_ = log.FromContext(ctx)
	//cmd := fmt.Sprintf("url=%s; site=%s; echo $(curl -sfL https://${url}/post-scripts/${site}/cleanup.sh | name=%s site=%s zone=%s address=%s)", r.NodeInfo.URL, r.NodeInfo.Site, r.NodeInfo.Name, r.NodeInfo.Site, r.NodeInfo.Zone, r.NodeInfo.Address)
	//cmd := fmt.Sprintf("url=%s; site=%s; echo $(curl -sfL https://${url}/post-scripts/${site}/cleanup.sh)", r.NodeInfo.URL, r.NodeInfo.Site)
	// out, err := exec.Command("bash", "-c", cmd).Output()
	// if err != nil {
	// 	fmt.Printf("error %s", err)
	// }
	// fmt.Println(string(out))

	///////////new pod///////////////
	// pod := &corev1.Pod{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		GenerateName: r.NodeInfo.Name + "cleanup-",
	// 		Labels:       labels,
	// 		Namespace:    "default",
	// 	},
	// 	Spec: corev1.PodSpec{
	// 		Containers: []corev1.Container{
	// 			{
	// 				Name:    "alpine",
	// 				Image:   "library/alpine/curl:3.15.4",
	// 				Command: []string{"sh", "-c"},
	// 				Args:    []string{cmd + "&& sleep 300"},
	// 				// Lifecycle: &corev1.Lifecycle{
	// 				// 	PostStart: &corev1.LifecycleHandler{
	// 				// 		Exec: &corev1.ExecAction{
	// 				// 			Command: []string{"sh", "-c", cmd},
	// 				// 		},
	// 				// 	},
	// 				// },
	// 				//{"apk add --no-cache curl","url=\"$r.NodeInfo.URL\" && echo $url",fmt.Sprintf("url=%s; site=%s; curl -sfL https://${url}/post-scripts/${site}/cleanup.sh | name=%s site=%s zone=%s address=%s sh - > res.txt",r.NodeInfo.URL, r.NodeInfo.Site, r.NodeInfo.Name, r.NodeInfo.Site, r.NodeInfo.Zone, r.NodeInfo.Address),"cat res.txt","sleep 300"},
	// 				// Env: []corev1.EnvVar{
	// 				// 	{
	// 				// 		Name: "CURLCMD",
	// 				// 		Value: fmt.Sprintf("url=%s; site=%s; curl -sfL https://${url}/post-scripts/${site}/cleanup.sh | name=%s site=%s zone=%s address=%s sh - > res.txt",r.NodeInfo.URL, r.NodeInfo.Site, r.NodeInfo.Name, r.NodeInfo.Site, r.NodeInfo.Zone, r.NodeInfo.Address),
	// 				// 	},
	// 				// },
	// 			},
	// 		},
	// 		RestartPolicy: corev1.RestartPolicyOnFailure,
	// 	},

	// err := r.Create(ctx, pod)
	// stlog.Println("trying pod create")
	// if err != nil {
	// 	return ctrl.Result{}, err
	// }
	///////////new pod///////////////

	// Set NodeCleanUpReconciler as the owner and controller of the Pod
	// if err := controllerutil.SetControllerReference(r, pod, r.Scheme); err != nil {
	// 	stlog.Println("controllerutil error \n")
	// 	return ctrl.Result{}, err
	// }

	// Check if this Pod already exists
	// found := &corev1.Pod{}
	// err := r.Get(ctx, types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	// if err != nil && errors.IsNotFound(err) {
	// 	// Pod not found, create it
	// 	err = r.Create(ctx, pod)
	// 	stlog.Println("trying pod create \n")
	// 	if err != nil {
	// 		return ctrl.Result{}, err
	// 	}
	// 	// Pod created successfully - don't requeue
	// 	return ctrl.Result{}, nil
	// } else if err != nil {
	// 	// Error getting the Pod
	// 	return ctrl.Result{}, err
	// }

	// Pod already exists - don't requeue
	return ctrl.Result{}, nil
}

// func NewPod(info *NodeInfo) *corev1.Pod {
// 	infoName:=info.Name
// 	infoAddress:=info.Status.Addresses[0].Address
// 	infoSite:=info.Labels["site"]

// 	labels := map[string]string{
// 			"DeletedNode": info.Name,
// 	}

// 	return &corev1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      infoName + "cleanup",
// 			Labels:    labels,
// 		},
// 		Spec: corev1.PodSpec{
// 			Containers: []corev1.Container{
// 				{
// 					Name:    "alpine",
// 					Image:   "alpine",
// 					Command:  []string{
// 						"/bin/sh",
// 						"-c",
// 						fmt.Sprintf(
// 							"echo 'Delete event received for Node' && echo 'Node Name: %s' && echo 'Node site: %s' && echo 'Node ip: %s'",
// 							infoName, infoSite, infoAddress,
// 						),
// 					},
// 				},
// 			},
// 			RestartPolicy: corev1.RestartPolicyOnFailure,
// 		},
// 	}
// }

// SetupWithManager sets up the controller with the Manager.
func (r *NodeDeletionWatcher) SetupWithManager(mgr ctrl.Manager) error {

	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Node{}). // Node 리소스에 대한 워처를 설정합니다.
		//Watches(&corev1.Event{}, &handler.EnqueueRequestForObject{}).
		WithEventFilter(predicate.Funcs{
			UpdateFunc: func(e event.UpdateEvent) bool {
				return false
			},
			DeleteFunc: func(e event.DeleteEvent) bool {
				node, ok := e.Object.(*corev1.Node)
				nodeName := node.Name
				nodeAddress := node.Status.Addresses[0].Address
				lenLabelSite := len(node.Labels["site"])         // to extract site & zone from label "site"
				nodeSite := node.Labels["site"][:lenLabelSite-4] //extract site from label "site"
				nodeZone := node.Labels["site"][lenLabelSite-3:] //extract zone from label "site"

				var url string
				if strings.Contains(nodeSite, "arc") {
					url = "gitlab.arc.hcloud.io/common/rancher/-/raw/master"
				} else if strings.Contains(nodeSite, "ap-northeast") {
					url = "/code.hcloud.hmc.co.kr/common/rancher/-/raw/master"
				} else if strings.Contains(nodeSite, "us-west") {
					url = "package-uswe.hcloud.hmc.co.kr/repository"
				} else if strings.Contains(nodeSite, "ap-southeast") {
					url = "package-sg.hcloud.hmc.co.kr/repository"
				} else if strings.Contains(nodeSite, "eu-central") {
					url = "package-eu.auto-hmg.io/repository"
				}

				if ok {
					cmdMutex.Lock()
					cmd = fmt.Sprintf("url=%s; site=%s; echo $(curl -sfL https://${url}/post-scripts/${site}/cleanup.sh | name=%s site=%s zone=%s address=%s sh -)", url, nodeSite, nodeName, nodeSite, nodeZone, nodeAddress)
					//probecmd := fmt.Sprintf("url=%s; site=%s; echo $(curl -sfL https://${url}/post-scripts/${site}/cleanup.sh | sh -)", url, nodeSite)
					cmdMutex.Unlock()
					fmt.Println("executing: ", cmd)

					deployment := createDeployment(nodeName)
					if err := r.Create(context.Background(), deployment); err != nil {
						fmt.Printf("Failed to create deployment: %v\n", err)
					} else {
						fmt.Printf("Created deployment for node %s\n", node.Name)
					}
				}

				return true
			},
			CreateFunc: func(e event.CreateEvent) bool {
				return false
			},
			GenericFunc: func(e event.GenericEvent) bool {
				return false
			},
		}).
		Complete(r)
}

func createDeployment(name string) *appsv1.Deployment {
	// Example deployment creation logic
	deployment := &appsv1.Deployment{
		// Fill in the deployment spec based on the node's labels
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: name + "-cleanup-",
			Namespace:    "default",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1), // Replica 수 설정
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "cleanup-" + name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "cleanup-" + name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "alpine",
							Image:   "library/alpine/curl:3.15.4",
							Command: []string{"sleep"},
							Args:    []string{"infinity"},
							ReadinessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{
										Command: []string{"sh", "-c", "response=\"$cmd\"; if [ -n \"$response\" ]; then echo \"$response\"; exit 0; else exit 1; fi"},
									},
								},
								InitialDelaySeconds: 20,
								PeriodSeconds:       10,
								FailureThreshold:    3,
							},
						},
					},
				},
			},
		},
	}
	return deployment
}

func int32Ptr(i int32) *int32 { return &i }
