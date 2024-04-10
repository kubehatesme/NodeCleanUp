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
	stlog "log"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"fmt"
	//"sigs.k8s.io/controller-runtime/pkg/reconcile"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	//"k8s.io/apimachinery/pkg/types"
	//"k8s.io/apimachinery/pkg/api/errors"
	//"k8s.io/utils/pointer"

)

// NodeCleanUpReconciler reconciles a NodeCleanUp object
type NodeCleanUpReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	NodeInfo *corev1.Node
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
func (r *NodeCleanUpReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//log := log.FromContext(ctx)

	//_ = log.FromContext(ctx)
	// if r.NodeInfo.Name != "" || r.NodeInfo.Site != "" || r.NodeInfo.IpAddress != "" {
	// 	//stlog.Print("Delete event received for Node\nNode Name: ",r.NodeInfo.Name, "\nNode site: ", r.NodeInfo.Site,"\nNode ip: ",r.NodeInfo.IpAddress)

	// 	pod := NewPod(&r.NodeInfo)

	// 	_, errCreate := ctrl.CreateOrUpdate(ctx, r.Client, pod, func() error {
    //             return ctrl.SetControllerReference(&r.NodeInfo, pod, r.Scheme)
    //     })

    //     if errCreate != nil {
    //             log.Error(errCreate, "Error creating pod")
    //             return ctrl.Result{}, nil
    //     }
		
	// }
	infoName:=r.NodeInfo.Name
	infoAddress:=r.NodeInfo.Status.Addresses[0].Address
	infoSite:=r.NodeInfo.Labels["site"]

	stlog.Println(infoName, " ",infoAddress, " ",infoSite)


	labels := map[string]string{
			"DeletedNode": infoName,
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      infoName + "cleanup",
			Labels:    labels,
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "alpine",
					Image:   "alpine",
					Command:  []string{
						"/bin/sh",
                    	"-c",
                    	fmt.Sprintf(
                        "echo 'Delete event received for Node' && echo 'Node Name: %s' && echo 'Node site: %s' && echo 'Node ip: %s' && sleep 300", // 5 minutes
                        infoName, infoSite, infoAddress,
                    ),
					},
				},
			},
			RestartPolicy: corev1.RestartPolicyOnFailure,
		},
	}

    // Set NodeCleanUpReconciler as the owner and controller of the Pod
    // if err := controllerutil.SetControllerReference(r, pod, r.Scheme); err != nil {
	// 	stlog.Println("controllerutil error \n")
	// 	return ctrl.Result{}, err
	// }

	err := r.Create(ctx, pod)
		stlog.Println("trying pod create \n")
		if err != nil {
			return ctrl.Result{}, err
		}

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
func (r *NodeCleanUpReconciler) SetupWithManager(mgr ctrl.Manager) error {

	return ctrl.NewControllerManagedBy(mgr).
        For(&corev1.Node{}). // Node 리소스에 대한 워처를 설정합니다.
		//Watches(&corev1.Event{}, &handler.EnqueueRequestForObject{}).
        WithEventFilter(predicate.Funcs{
			UpdateFunc: func(e event.UpdateEvent) bool {
			 return false
			},
			DeleteFunc: func(e event.DeleteEvent) bool {
				r.NodeInfo=e.Object.(*corev1.Node)
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
