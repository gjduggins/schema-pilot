/*
Copyright 2025.

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
	"fmt"

	"k8s.io/apimachinery/pkg/types"

	schemapilotv1 "github.com/gjduggins/schemapilot-operator/api/v1"
	"github.com/gjduggins/schemapilot-operator/internal/db"
	"github.com/gjduggins/schemapilot-operator/internal/manager"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// SchemaPilotReconciler reconciles a SchemaPilot object
type SchemaPilotReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch
// +kubebuilder:rbac:groups=schemapilot.schemapilot.com,resources=schemapilots,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=schemapilot.schemapilot.com,resources=schemapilots/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=schemapilot.schemapilot.com,resources=schemapilots/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SchemaPilot object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *SchemaPilotReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	schemapilot := &schemapilotv1.SchemaPilot{}
	if err := r.Get(ctx, req.NamespacedName, schemapilot); err != nil {
		log.Error(err, "unable to fetch SchemaPilot")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("Reconciliation loop running", "SchemaPilot", req.NamespacedName, "DatabaseType", schemapilot.Spec.DatabaseType)

	connSecret := schemapilot.Spec.ConnectionString
	connStr, err := r.getSecretValue(ctx, req.Namespace, connSecret.Name, connSecret.Key)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed reading connectionSecret: %w", err)
	}
	// Create Oracle client connection
	client, err := db.Connect(
		ctx,
		connStr,
	)

	if err != nil {
		log.Error(err, "failed to connect to database")
		return ctrl.Result{}, err
	}
	defer client.Close()

	repoTokenRef := schemapilot.Spec.Repository.Token
	repoToken, err := r.getSecretValue(ctx, req.Namespace, repoTokenRef.Name, repoTokenRef.Key)
	if err != nil {
		log.Error(err, "failed reading repository.token")
		return ctrl.Result{}, err
	}
	adapter := db.NewGitHubAdapter(schemapilot.Spec.Repository.Url, schemapilot.Spec.Repository.Owner, schemapilot.Spec.Repository.Repo, repoToken)
	log.Info("ApplySchemaChanges starting")

	schemaManager := manager.NewSchemaManager(client, adapter)
	schemaManager.ApplySchemaChanges(ctx, schemapilot)

	log.Info("SchemaPilot reconciliation completed successfully", "SchemaPilot", req.NamespacedName)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SchemaPilotReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&schemapilotv1.SchemaPilot{}).
		Named("schemapilot").
		Complete(r)
}

func (r *SchemaPilotReconciler) getSecretValue(ctx context.Context, namespace, name, key string) (string, error) {
	secret := &corev1.Secret{}
	if err := r.Client.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, secret); err != nil {
		return "", err
	}
	val, ok := secret.Data[key]
	if !ok {
		return "", fmt.Errorf("key %q not found in Secret %q", key, name)
	}
	return string(val), nil
}
