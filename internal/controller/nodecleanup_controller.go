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
)

// NodeCleanUpReconciler reconciles a NodeCleanUp object
type NodeCleanUpReconciler struct {
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
func (r *NodeCleanUpReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//_ = log.FromContext(ctx)

	node := &corev1.Node{}
	if err := r.Get(ctx, req.NamespacedName, node); err != nil {
			// You can add your handling logic here, like marking some other resource as deleted or triggering other actions.
			stlog.Print("Error getting Node")
			return ctrl.Result{}, err
	}

	if !node.ObjectMeta.DeletionTimestamp.IsZero() {
		stlog.Printf("Node %s has been deleted!\n", node.Name)
	}

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodeCleanUpReconciler) SetupWithManager(mgr ctrl.Manager) error {

	var DeletedNode *corev1.Node

	return ctrl.NewControllerManagedBy(mgr).
        For(&corev1.Node{}). // Node 리소스에 대한 워처를 설정합니다.
		//Watches(&corev1.Event{}, &handler.EnqueueRequestForObject{}).
        WithEventFilter(predicate.Funcs{
			UpdateFunc: func(e event.UpdateEvent) bool {
			 return false
			},
			DeleteFunc: func(e event.DeleteEvent) bool {

				DeletedNode = e.Object.(*corev1.Node)

				stlog.Print("Delete event received for Node\nNode Name: ",DeletedNode.ObjectMeta.Name, "\nNode site: ", DeletedNode.Labels["site"],"\nNode ip: ",DeletedNode.Status.Addresses[0].Address)
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
