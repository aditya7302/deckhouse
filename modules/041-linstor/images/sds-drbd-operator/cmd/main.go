package main

import (
	"context"
	"fmt"
	lclient "github.com/LINBIT/golinstor/client"
	"go.uber.org/zap/zapcore"
	v1 "k8s.io/api/core/v1"
	sv1 "k8s.io/api/storage/v1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiruntime "k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"os"
	goruntime "runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"st2/api/v1alpha1"
	"st2/config"
	"st2/pkg/controller"
	kubutils "st2/pkg/kubeutils"
)

const (
	OperatorNamespace = "d8-linstor"
)

var (
	resourcesSchemeFuncs = []func(*apiruntime.Scheme) error{
		v1alpha1.AddToScheme,
		clientgoscheme.AddToScheme,
		extv1.AddToScheme,
		v1.AddToScheme,
		sv1.AddToScheme,
	}
)

func main() {
	log := zap.New(zap.Level(zapcore.Level(-1)), zap.UseDevMode(true))
	// logf.SetLogger(zap.New(zap.Level(zapcore.Level(-1)), zap.UseDevMode(true)))
	// log := logf.Log.WithName("cmd")
	// log.WithName("cmd")

	ctx, _ := context.WithCancel(context.Background())

	log.Info(fmt.Sprintf("Go Version:%s ", goruntime.Version()))
	log.Info(fmt.Sprintf("OS/Arch:Go OS/Arch:%s/%s ", goruntime.GOOS, goruntime.GOARCH))

	cfgParams, err := config.NewConfig()
	if err != nil {
		log.Error(err, "error read configuration")
	}

	// Create default config Kubernetes client
	kConfig, err := kubutils.KubernetesDefaultConfigCreate()
	if err != nil {
		log.Error(err, "error read kubernetes configuration")
	}
	log.Info("read Kubernetes config")

	// Setup scheme for all resources
	scheme := runtime.NewScheme()
	for _, f := range resourcesSchemeFuncs {
		err := f(scheme)
		if err != nil {
			log.Error(err, "failed to add to scheme")
			os.Exit(1)
		}
	}
	log.Info("read scheme CR")

	// ---- server webhook ----------
	// myWebhookServer := webhook.NewServer(webhook.Options{
	//
	// })

	managerOpts := manager.Options{
		// LeaderElection:             true,
		// LeaderElectionNamespace:    "d8-storage-configurator",
		// LeaderElectionID:           "r", // uniq in cluster
		// LeaderElectionResourceLock: "leases",
		Scheme:             scheme,
		MetricsBindAddress: cfgParams.MetricsPort,
		Logger:             log,
		Namespace:          OperatorNamespace, // TODO: заменить на cache options
		// WebhookServer: myWebhookServer,
	}

	mgr, err := manager.New(kConfig, managerOpts)
	if err != nil {
		log.Error(err, "failed create manager")
		os.Exit(1)
	}

	log.Info("create kubernetes manager")

	lc, err := lclient.NewClient()

	// --------------------------------------------
	if _, err := controller.NewLinstorNode(ctx, mgr, lc, cfgParams.ConfigSecretName, cfgParams.ScanInterval); err != nil {
		log.Error(err, "failed create controller NewLinstorNode", err)
		os.Exit(1)
	}
	log.Info("controller NewLinstorNode start")

	if _, err := controller.NewLinstorStorageClass(ctx, mgr); err != nil {
		log.Error(err, "failed create controller NewLinstorStorageClass")
		os.Exit(1)
	}
	log.Info("controller NewLinstorStorageClass start")

	if _, err := controller.NewLinstorStoragePool(ctx, mgr, lc); err != nil {
		log.Error(err, "failed create controller NewLinstorStoragePool", err)
		os.Exit(1)
	}
	log.Info("controller NewLinstorStoragePool start")

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		log.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		log.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	err = mgr.Start(ctx)
	if err != nil {
		log.Error(err, "error start manager")
		os.Exit(1)
	}

	log.Info("starting the manager")
}
