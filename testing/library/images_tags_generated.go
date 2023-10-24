// Code generated by "tools/images_tags.go" DO NOT EDIT.
// To generate run 'make generate'
package library

var DefaultImagesDigests = map[string]interface{}{
	"admissionPolicyEngine": map[string]interface{}{
		"constraintExporter": "imageHash-admissionPolicyEngine-constraintExporter",
		"gatekeeper":         "imageHash-admissionPolicyEngine-gatekeeper",
		"trivyProvider":      "imageHash-admissionPolicyEngine-trivyProvider",
	},
	"basicAuth": map[string]interface{}{
		"nginx": "imageHash-basicAuth-nginx",
	},
	"cephCsi": map[string]interface{}{
		"cephcsi": "imageHash-cephCsi-cephcsi",
	},
	"certManager": map[string]interface{}{
		"certManagerAcmeSolver": "imageHash-certManager-certManagerAcmeSolver",
		"certManagerCainjector": "imageHash-certManager-certManagerCainjector",
		"certManagerController": "imageHash-certManager-certManagerController",
		"certManagerWebhook":    "imageHash-certManager-certManagerWebhook",
	},
	"chrony": map[string]interface{}{
		"chrony": "imageHash-chrony-chrony",
	},
	"ciliumHubble": map[string]interface{}{
		"relay":      "imageHash-ciliumHubble-relay",
		"uiBackend":  "imageHash-ciliumHubble-uiBackend",
		"uiFrontend": "imageHash-ciliumHubble-uiFrontend",
	},
	"cloudProviderAws": map[string]interface{}{
		"cloudControllerManager124": "imageHash-cloudProviderAws-cloudControllerManager124",
		"cloudControllerManager125": "imageHash-cloudProviderAws-cloudControllerManager125",
		"cloudControllerManager126": "imageHash-cloudProviderAws-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderAws-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderAws-cloudControllerManager128",
		"cloudDataDiscoverer":       "imageHash-cloudProviderAws-cloudDataDiscoverer",
		"ebsCsiPlugin":              "imageHash-cloudProviderAws-ebsCsiPlugin",
		"nodeTerminationHandler":    "imageHash-cloudProviderAws-nodeTerminationHandler",
	},
	"cloudProviderAzure": map[string]interface{}{
		"azurediskCsi":              "imageHash-cloudProviderAzure-azurediskCsi",
		"cloudControllerManager124": "imageHash-cloudProviderAzure-cloudControllerManager124",
		"cloudControllerManager125": "imageHash-cloudProviderAzure-cloudControllerManager125",
		"cloudControllerManager126": "imageHash-cloudProviderAzure-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderAzure-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderAzure-cloudControllerManager128",
		"cloudDataDiscoverer":       "imageHash-cloudProviderAzure-cloudDataDiscoverer",
	},
	"cloudProviderGcp": map[string]interface{}{
		"cloudControllerManager124": "imageHash-cloudProviderGcp-cloudControllerManager124",
		"cloudControllerManager125": "imageHash-cloudProviderGcp-cloudControllerManager125",
		"cloudControllerManager126": "imageHash-cloudProviderGcp-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderGcp-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderGcp-cloudControllerManager128",
		"cloudDataDiscoverer":       "imageHash-cloudProviderGcp-cloudDataDiscoverer",
		"pdCsiPlugin":               "imageHash-cloudProviderGcp-pdCsiPlugin",
	},
	"cloudProviderOpenstack": map[string]interface{}{
		"cinderCsiPlugin124":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin124",
		"cinderCsiPlugin125":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin125",
		"cinderCsiPlugin126":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin126",
		"cinderCsiPlugin127":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin127",
		"cinderCsiPlugin128":        "imageHash-cloudProviderOpenstack-cinderCsiPlugin128",
		"cloudControllerManager124": "imageHash-cloudProviderOpenstack-cloudControllerManager124",
		"cloudControllerManager125": "imageHash-cloudProviderOpenstack-cloudControllerManager125",
		"cloudControllerManager126": "imageHash-cloudProviderOpenstack-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderOpenstack-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderOpenstack-cloudControllerManager128",
		"cloudDataDiscoverer":       "imageHash-cloudProviderOpenstack-cloudDataDiscoverer",
	},
	"cloudProviderVsphere": map[string]interface{}{
		"cloudControllerManager124": "imageHash-cloudProviderVsphere-cloudControllerManager124",
		"cloudControllerManager125": "imageHash-cloudProviderVsphere-cloudControllerManager125",
		"cloudControllerManager126": "imageHash-cloudProviderVsphere-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderVsphere-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderVsphere-cloudControllerManager128",
		"vsphereCsiPlugin":          "imageHash-cloudProviderVsphere-vsphereCsiPlugin",
		"vsphereCsiPluginLegacy":    "imageHash-cloudProviderVsphere-vsphereCsiPluginLegacy",
	},
	"cloudProviderYandex": map[string]interface{}{
		"cloudControllerManager124": "imageHash-cloudProviderYandex-cloudControllerManager124",
		"cloudControllerManager125": "imageHash-cloudProviderYandex-cloudControllerManager125",
		"cloudControllerManager126": "imageHash-cloudProviderYandex-cloudControllerManager126",
		"cloudControllerManager127": "imageHash-cloudProviderYandex-cloudControllerManager127",
		"cloudControllerManager128": "imageHash-cloudProviderYandex-cloudControllerManager128",
		"cloudMetricsExporter":      "imageHash-cloudProviderYandex-cloudMetricsExporter",
		"yandexCsiPlugin":           "imageHash-cloudProviderYandex-yandexCsiPlugin",
	},
	"cniCilium": map[string]interface{}{
		"builderArtifact":        "imageHash-cniCilium-builderArtifact",
		"builderRuntimeArtifact": "imageHash-cniCilium-builderRuntimeArtifact",
		"cilium":                 "imageHash-cniCilium-cilium",
		"operator":               "imageHash-cniCilium-operator",
		"virtCilium":             "imageHash-cniCilium-virtCilium",
	},
	"cniFlannel": map[string]interface{}{
		"flanneld": "imageHash-cniFlannel-flanneld",
	},
	"cniSimpleBridge": map[string]interface{}{
		"simpleBridge": "imageHash-cniSimpleBridge-simpleBridge",
	},
	"common": map[string]interface{}{
		"alpine":                    "imageHash-common-alpine",
		"checkKernelVersion":        "imageHash-common-checkKernelVersion",
		"csiExternalAttacher124":    "imageHash-common-csiExternalAttacher124",
		"csiExternalAttacher125":    "imageHash-common-csiExternalAttacher125",
		"csiExternalAttacher126":    "imageHash-common-csiExternalAttacher126",
		"csiExternalAttacher127":    "imageHash-common-csiExternalAttacher127",
		"csiExternalAttacher128":    "imageHash-common-csiExternalAttacher128",
		"csiExternalProvisioner124": "imageHash-common-csiExternalProvisioner124",
		"csiExternalProvisioner125": "imageHash-common-csiExternalProvisioner125",
		"csiExternalProvisioner126": "imageHash-common-csiExternalProvisioner126",
		"csiExternalProvisioner127": "imageHash-common-csiExternalProvisioner127",
		"csiExternalProvisioner128": "imageHash-common-csiExternalProvisioner128",
		"csiExternalResizer124":     "imageHash-common-csiExternalResizer124",
		"csiExternalResizer125":     "imageHash-common-csiExternalResizer125",
		"csiExternalResizer126":     "imageHash-common-csiExternalResizer126",
		"csiExternalResizer127":     "imageHash-common-csiExternalResizer127",
		"csiExternalResizer128":     "imageHash-common-csiExternalResizer128",
		"csiExternalSnapshotter124": "imageHash-common-csiExternalSnapshotter124",
		"csiExternalSnapshotter125": "imageHash-common-csiExternalSnapshotter125",
		"csiExternalSnapshotter126": "imageHash-common-csiExternalSnapshotter126",
		"csiExternalSnapshotter127": "imageHash-common-csiExternalSnapshotter127",
		"csiExternalSnapshotter128": "imageHash-common-csiExternalSnapshotter128",
		"csiLivenessprobe124":       "imageHash-common-csiLivenessprobe124",
		"csiLivenessprobe125":       "imageHash-common-csiLivenessprobe125",
		"csiLivenessprobe126":       "imageHash-common-csiLivenessprobe126",
		"csiLivenessprobe127":       "imageHash-common-csiLivenessprobe127",
		"csiLivenessprobe128":       "imageHash-common-csiLivenessprobe128",
		"csiNodeDriverRegistrar124": "imageHash-common-csiNodeDriverRegistrar124",
		"csiNodeDriverRegistrar125": "imageHash-common-csiNodeDriverRegistrar125",
		"csiNodeDriverRegistrar126": "imageHash-common-csiNodeDriverRegistrar126",
		"csiNodeDriverRegistrar127": "imageHash-common-csiNodeDriverRegistrar127",
		"csiNodeDriverRegistrar128": "imageHash-common-csiNodeDriverRegistrar128",
		"distroless":                "imageHash-common-distroless",
		"iptablesWrapper":           "imageHash-common-iptablesWrapper",
		"kubeRbacProxy":             "imageHash-common-kubeRbacProxy",
		"nginxStatic":               "imageHash-common-nginxStatic",
		"pause":                     "imageHash-common-pause",
		"shellOperator":             "imageHash-common-shellOperator",
	},
	"containerizedDataImporter": map[string]interface{}{
		"cdiApiserver":    "imageHash-containerizedDataImporter-cdiApiserver",
		"cdiCloner":       "imageHash-containerizedDataImporter-cdiCloner",
		"cdiController":   "imageHash-containerizedDataImporter-cdiController",
		"cdiImporter":     "imageHash-containerizedDataImporter-cdiImporter",
		"cdiOperator":     "imageHash-containerizedDataImporter-cdiOperator",
		"cdiUploadproxy":  "imageHash-containerizedDataImporter-cdiUploadproxy",
		"cdiUploadserver": "imageHash-containerizedDataImporter-cdiUploadserver",
	},
	"controlPlaneManager": map[string]interface{}{
		"controlPlaneManager124":   "imageHash-controlPlaneManager-controlPlaneManager124",
		"controlPlaneManager125":   "imageHash-controlPlaneManager-controlPlaneManager125",
		"controlPlaneManager126":   "imageHash-controlPlaneManager-controlPlaneManager126",
		"controlPlaneManager127":   "imageHash-controlPlaneManager-controlPlaneManager127",
		"controlPlaneManager128":   "imageHash-controlPlaneManager-controlPlaneManager128",
		"etcd":                     "imageHash-controlPlaneManager-etcd",
		"kubeApiserver124":         "imageHash-controlPlaneManager-kubeApiserver124",
		"kubeApiserver125":         "imageHash-controlPlaneManager-kubeApiserver125",
		"kubeApiserver126":         "imageHash-controlPlaneManager-kubeApiserver126",
		"kubeApiserver127":         "imageHash-controlPlaneManager-kubeApiserver127",
		"kubeApiserver128":         "imageHash-controlPlaneManager-kubeApiserver128",
		"kubeApiserverHealthcheck": "imageHash-controlPlaneManager-kubeApiserverHealthcheck",
		"kubeControllerManager124": "imageHash-controlPlaneManager-kubeControllerManager124",
		"kubeControllerManager125": "imageHash-controlPlaneManager-kubeControllerManager125",
		"kubeControllerManager126": "imageHash-controlPlaneManager-kubeControllerManager126",
		"kubeControllerManager127": "imageHash-controlPlaneManager-kubeControllerManager127",
		"kubeControllerManager128": "imageHash-controlPlaneManager-kubeControllerManager128",
		"kubeScheduler124":         "imageHash-controlPlaneManager-kubeScheduler124",
		"kubeScheduler125":         "imageHash-controlPlaneManager-kubeScheduler125",
		"kubeScheduler126":         "imageHash-controlPlaneManager-kubeScheduler126",
		"kubeScheduler127":         "imageHash-controlPlaneManager-kubeScheduler127",
		"kubeScheduler128":         "imageHash-controlPlaneManager-kubeScheduler128",
		"kubernetesApiProxy":       "imageHash-controlPlaneManager-kubernetesApiProxy",
	},
	"dashboard": map[string]interface{}{
		"dashboard":      "imageHash-dashboard-dashboard",
		"metricsScraper": "imageHash-dashboard-metricsScraper",
	},
	"deckhouse": map[string]interface{}{
		"webhookHandler": "imageHash-deckhouse-webhookHandler",
	},
	"delivery": map[string]interface{}{
		"argocd":               "imageHash-delivery-argocd",
		"argocdImageUpdater":   "imageHash-delivery-argocdImageUpdater",
		"redis":                "imageHash-delivery-redis",
		"werfArgocdCmpSidecar": "imageHash-delivery-werfArgocdCmpSidecar",
	},
	"descheduler": map[string]interface{}{
		"descheduler": "imageHash-descheduler-descheduler",
	},
	"documentation": map[string]interface{}{
		"web": "imageHash-documentation-web",
	},
	"extendedMonitoring": map[string]interface{}{
		"certExporter":               "imageHash-extendedMonitoring-certExporter",
		"eventsExporter":             "imageHash-extendedMonitoring-eventsExporter",
		"extendedMonitoringExporter": "imageHash-extendedMonitoring-extendedMonitoringExporter",
		"imageAvailabilityExporter":  "imageHash-extendedMonitoring-imageAvailabilityExporter",
	},
	"flantIntegration": map[string]interface{}{
		"flantPricing": "imageHash-flantIntegration-flantPricing",
		"grafanaAgent": "imageHash-flantIntegration-grafanaAgent",
		"madisonProxy": "imageHash-flantIntegration-madisonProxy",
	},
	"ingressNginx": map[string]interface{}{
		"controller11":          "imageHash-ingressNginx-controller11",
		"controller16":          "imageHash-ingressNginx-controller16",
		"kruise":                "imageHash-ingressNginx-kruise",
		"kruiseStateMetrics":    "imageHash-ingressNginx-kruiseStateMetrics",
		"kubeRbacProxy":         "imageHash-ingressNginx-kubeRbacProxy",
		"nginxExporter":         "imageHash-ingressNginx-nginxExporter",
		"protobufExporter":      "imageHash-ingressNginx-protobufExporter",
		"proxyFailover":         "imageHash-ingressNginx-proxyFailover",
		"proxyFailoverIptables": "imageHash-ingressNginx-proxyFailoverIptables",
	},
	"istio": map[string]interface{}{
		"apiProxy":          "imageHash-istio-apiProxy",
		"kiali":             "imageHash-istio-kiali",
		"metadataDiscovery": "imageHash-istio-metadataDiscovery",
		"metadataExporter":  "imageHash-istio-metadataExporter",
		"operatorV1x12x6":   "imageHash-istio-operatorV1x12x6",
		"operatorV1x13x7":   "imageHash-istio-operatorV1x13x7",
		"operatorV1x16x2":   "imageHash-istio-operatorV1x16x2",
		"pilotV1x12x6":      "imageHash-istio-pilotV1x12x6",
		"pilotV1x13x7":      "imageHash-istio-pilotV1x13x7",
		"pilotV1x16x2":      "imageHash-istio-pilotV1x16x2",
		"proxyv2V1x12x6":    "imageHash-istio-proxyv2V1x12x6",
		"proxyv2V1x13x7":    "imageHash-istio-proxyv2V1x13x7",
		"proxyv2V1x16x2":    "imageHash-istio-proxyv2V1x16x2",
	},
	"keepalived": map[string]interface{}{
		"keepalived": "imageHash-keepalived-keepalived",
	},
	"kubeDns": map[string]interface{}{
		"coredns":                           "imageHash-kubeDns-coredns",
		"resolvWatcher":                     "imageHash-kubeDns-resolvWatcher",
		"stsPodsHostsAppenderInitContainer": "imageHash-kubeDns-stsPodsHostsAppenderInitContainer",
		"stsPodsHostsAppenderWebhook":       "imageHash-kubeDns-stsPodsHostsAppenderWebhook",
	},
	"kubeProxy": map[string]interface{}{
		"initContainer": "imageHash-kubeProxy-initContainer",
		"kubeProxy124":  "imageHash-kubeProxy-kubeProxy124",
		"kubeProxy125":  "imageHash-kubeProxy-kubeProxy125",
		"kubeProxy126":  "imageHash-kubeProxy-kubeProxy126",
		"kubeProxy127":  "imageHash-kubeProxy-kubeProxy127",
		"kubeProxy128":  "imageHash-kubeProxy-kubeProxy128",
	},
	"linstor": map[string]interface{}{
		"drbdReactor":               "imageHash-linstor-drbdReactor",
		"linstorAffinityController": "imageHash-linstor-linstorAffinityController",
		"linstorCsi":                "imageHash-linstor-linstorCsi",
		"linstorDrbdWait":           "imageHash-linstor-linstorDrbdWait",
		"linstorPoolsImporter":      "imageHash-linstor-linstorPoolsImporter",
		"linstorSchedulerAdmission": "imageHash-linstor-linstorSchedulerAdmission",
		"linstorSchedulerExtender":  "imageHash-linstor-linstorSchedulerExtender",
		"linstorServer":             "imageHash-linstor-linstorServer",
		"piraeusOperator":           "imageHash-linstor-piraeusOperator",
		"spaas":                     "imageHash-linstor-spaas",
	},
	"localPathProvisioner": map[string]interface{}{
		"helper":               "imageHash-localPathProvisioner-helper",
		"localPathProvisioner": "imageHash-localPathProvisioner-localPathProvisioner",
	},
	"logShipper": map[string]interface{}{
		"vector": "imageHash-logShipper-vector",
	},
	"loki": map[string]interface{}{
		"loki": "imageHash-loki-loki",
	},
	"metallb": map[string]interface{}{
		"controller": "imageHash-metallb-controller",
		"speaker":    "imageHash-metallb-speaker",
	},
	"monitoringKubernetes": map[string]interface{}{
		"ebpfExporter":                      "imageHash-monitoringKubernetes-ebpfExporter",
		"kubeStateMetrics":                  "imageHash-monitoringKubernetes-kubeStateMetrics",
		"kubeletEvictionThresholdsExporter": "imageHash-monitoringKubernetes-kubeletEvictionThresholdsExporter",
		"nodeExporter":                      "imageHash-monitoringKubernetes-nodeExporter",
	},
	"monitoringPing": map[string]interface{}{
		"monitoringPing": "imageHash-monitoringPing-monitoringPing",
	},
	"networkGateway": map[string]interface{}{
		"dnsmasq": "imageHash-networkGateway-dnsmasq",
		"snat":    "imageHash-networkGateway-snat",
	},
	"networkPolicyEngine": map[string]interface{}{
		"kubeRouter": "imageHash-networkPolicyEngine-kubeRouter",
	},
	"nodeLocalDns": map[string]interface{}{
		"coredns":      "imageHash-nodeLocalDns-coredns",
		"iptablesLoop": "imageHash-nodeLocalDns-iptablesLoop",
	},
	"nodeManager": map[string]interface{}{
		"bashibleApiserver":        "imageHash-nodeManager-bashibleApiserver",
		"capiControllerManager":    "imageHash-nodeManager-capiControllerManager",
		"capsControllerManager":    "imageHash-nodeManager-capsControllerManager",
		"clusterAutoscaler124":     "imageHash-nodeManager-clusterAutoscaler124",
		"clusterAutoscaler125":     "imageHash-nodeManager-clusterAutoscaler125",
		"clusterAutoscaler126":     "imageHash-nodeManager-clusterAutoscaler126",
		"clusterAutoscaler127":     "imageHash-nodeManager-clusterAutoscaler127",
		"clusterAutoscaler128":     "imageHash-nodeManager-clusterAutoscaler128",
		"earlyOom":                 "imageHash-nodeManager-earlyOom",
		"machineControllerManager": "imageHash-nodeManager-machineControllerManager",
	},
	"openvpn": map[string]interface{}{
		"easyrsaMigrator": "imageHash-openvpn-easyrsaMigrator",
		"openvpn":         "imageHash-openvpn-openvpn",
		"ovpnAdmin":       "imageHash-openvpn-ovpnAdmin",
		"pmacct":          "imageHash-openvpn-pmacct",
	},
	"operatorPrometheus": map[string]interface{}{
		"prometheusConfigReloader": "imageHash-operatorPrometheus-prometheusConfigReloader",
		"prometheusOperator":       "imageHash-operatorPrometheus-prometheusOperator",
	},
	"operatorTrivy": map[string]interface{}{
		"nodeCollector": "imageHash-operatorTrivy-nodeCollector",
		"operator":      "imageHash-operatorTrivy-operator",
		"reportUpdater": "imageHash-operatorTrivy-reportUpdater",
		"trivy":         "imageHash-operatorTrivy-trivy",
	},
	"podReloader": map[string]interface{}{
		"podReloader": "imageHash-podReloader-podReloader",
	},
	"prometheus": map[string]interface{}{
		"alertmanager":                "imageHash-prometheus-alertmanager",
		"alertsReceiver":              "imageHash-prometheus-alertsReceiver",
		"grafana":                     "imageHash-prometheus-grafana",
		"grafanaDashboardProvisioner": "imageHash-prometheus-grafanaDashboardProvisioner",
		"prometheus":                  "imageHash-prometheus-prometheus",
		"trickster":                   "imageHash-prometheus-trickster",
	},
	"prometheusMetricsAdapter": map[string]interface{}{
		"k8sPrometheusAdapter":   "imageHash-prometheusMetricsAdapter-k8sPrometheusAdapter",
		"prometheusReverseProxy": "imageHash-prometheusMetricsAdapter-prometheusReverseProxy",
	},
	"prometheusPushgateway": map[string]interface{}{
		"pushgateway": "imageHash-prometheusPushgateway-pushgateway",
	},
	"registrypackages": map[string]interface{}{
		"containerd1624":   "imageHash-registrypackages-containerd1624",
		"crictl124":        "imageHash-registrypackages-crictl124",
		"crictl125":        "imageHash-registrypackages-crictl125",
		"crictl126":        "imageHash-registrypackages-crictl126",
		"crictl127":        "imageHash-registrypackages-crictl127",
		"crictl128":        "imageHash-registrypackages-crictl128",
		"d8Curl821":        "imageHash-registrypackages-d8Curl821",
		"drbd":             "imageHash-registrypackages-drbd",
		"jq16":             "imageHash-registrypackages-jq16",
		"kubeadm12417":     "imageHash-registrypackages-kubeadm12417",
		"kubeadm12515":     "imageHash-registrypackages-kubeadm12515",
		"kubeadm12610":     "imageHash-registrypackages-kubeadm12610",
		"kubeadm1277":      "imageHash-registrypackages-kubeadm1277",
		"kubeadm1283":      "imageHash-registrypackages-kubeadm1283",
		"kubectl12417":     "imageHash-registrypackages-kubectl12417",
		"kubectl12515":     "imageHash-registrypackages-kubectl12515",
		"kubectl12610":     "imageHash-registrypackages-kubectl12610",
		"kubectl1277":      "imageHash-registrypackages-kubectl1277",
		"kubectl1283":      "imageHash-registrypackages-kubectl1283",
		"kubelet12417":     "imageHash-registrypackages-kubelet12417",
		"kubelet12515":     "imageHash-registrypackages-kubelet12515",
		"kubelet12610":     "imageHash-registrypackages-kubelet12610",
		"kubelet1277":      "imageHash-registrypackages-kubelet1277",
		"kubelet1283":      "imageHash-registrypackages-kubelet1283",
		"kubernetesCni120": "imageHash-registrypackages-kubernetesCni120",
		"tomlMerge01":      "imageHash-registrypackages-tomlMerge01",
	},
	"runtimeAuditEngine": map[string]interface{}{
		"falco":         "imageHash-runtimeAuditEngine-falco",
		"falcosidekick": "imageHash-runtimeAuditEngine-falcosidekick",
		"rulesLoader":   "imageHash-runtimeAuditEngine-rulesLoader",
	},
	"snapshotController": map[string]interface{}{
		"snapshotController":        "imageHash-snapshotController-snapshotController",
		"snapshotValidationWebhook": "imageHash-snapshotController-snapshotValidationWebhook",
	},
	"terraformManager": map[string]interface{}{
		"baseTerraformManager":      "imageHash-terraformManager-baseTerraformManager",
		"terraformManagerAws":       "imageHash-terraformManager-terraformManagerAws",
		"terraformManagerAzure":     "imageHash-terraformManager-terraformManagerAzure",
		"terraformManagerGcp":       "imageHash-terraformManager-terraformManagerGcp",
		"terraformManagerOpenstack": "imageHash-terraformManager-terraformManagerOpenstack",
		"terraformManagerVsphere":   "imageHash-terraformManager-terraformManagerVsphere",
		"terraformManagerYandex":    "imageHash-terraformManager-terraformManagerYandex",
	},
	"upmeter": map[string]interface{}{
		"smokeMini": "imageHash-upmeter-smokeMini",
		"status":    "imageHash-upmeter-status",
		"upmeter":   "imageHash-upmeter-upmeter",
		"webui":     "imageHash-upmeter-webui",
	},
	"userAuthn": map[string]interface{}{
		"crowdBasicAuthProxy":   "imageHash-userAuthn-crowdBasicAuthProxy",
		"dex":                   "imageHash-userAuthn-dex",
		"dexAuthenticator":      "imageHash-userAuthn-dexAuthenticator",
		"dexAuthenticatorRedis": "imageHash-userAuthn-dexAuthenticatorRedis",
		"kubeconfigGenerator":   "imageHash-userAuthn-kubeconfigGenerator",
		"selfSignedGenerator":   "imageHash-userAuthn-selfSignedGenerator",
	},
	"userAuthz": map[string]interface{}{
		"webhook": "imageHash-userAuthz-webhook",
	},
	"verticalPodAutoscaler": map[string]interface{}{
		"admissionController": "imageHash-verticalPodAutoscaler-admissionController",
		"recommender":         "imageHash-verticalPodAutoscaler-recommender",
		"updater":             "imageHash-verticalPodAutoscaler-updater",
	},
	"virtualization": map[string]interface{}{
		"imageAlpineBionic": "imageHash-virtualization-imageAlpineBionic",
		"imageAlpineFocal":  "imageHash-virtualization-imageAlpineFocal",
		"imageAlpineJammy":  "imageHash-virtualization-imageAlpineJammy",
		"imageUbuntuBionic": "imageHash-virtualization-imageUbuntuBionic",
		"imageUbuntuFocal":  "imageHash-virtualization-imageUbuntuFocal",
		"imageUbuntuJammy":  "imageHash-virtualization-imageUbuntuJammy",
		"libguestfs":        "imageHash-virtualization-libguestfs",
		"virtApi":           "imageHash-virtualization-virtApi",
		"virtController":    "imageHash-virtualization-virtController",
		"virtExportproxy":   "imageHash-virtualization-virtExportproxy",
		"virtExportserver":  "imageHash-virtualization-virtExportserver",
		"virtHandler":       "imageHash-virtualization-virtHandler",
		"virtLauncher":      "imageHash-virtualization-virtLauncher",
		"virtOperator":      "imageHash-virtualization-virtOperator",
		"vmiRouter":         "imageHash-virtualization-vmiRouter",
	},
}
