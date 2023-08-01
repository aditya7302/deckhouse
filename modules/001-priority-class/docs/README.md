---
title: "The priority-class module"
---

This module creates a set of [priority classes](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass) and assigns them to components installed by Deckhouse and applications in the cluster.

[Priority Class](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption) relates to the scheduler and allows it to schedule a Pod based on its priority (which is defined by the class the Pod belongs to).

Suppose we need to schedule a Pod belonging to the `priorityClassName: production-low` priority class. If the cluster does not have enough resources for this Pod, Kubernetes will start evicting Pods with the lowest priority to deploy our `production-low` Pod.
That is, Kubernetes will first evict all the `priorityClassName: develop` Pods, then proceed to `cluster-low` Pods, and so on.

When setting the priority class, it is crucial to understand what kind of application we have and what environment this application works in. Any `priorityClassName` set to a Pod cannot lower its priority because the scheduler considers Pods without `priority-class` as having the lowest (`develop`) priority. It is essential to set the `priorityClassName` correctly.

> **Caution!** You cannot use the following PriorityClasses: `system-node-critical`, `system-cluster-critical`, `cluster-medium`, `cluster-low`.

Below is the list of priority classes set by the module (sorted by the priority, starting with the higher one):

| Priority Class            | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Value      |
|---------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------|
| `system-node-critical`    | Cluster components that are must to be present on the node. This priority class fully protects components against [eviction by kubelet](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/).<br>`node-exporter`, `csi`                                                                                                                                                                                                                     | 2000001000 |
| `system-cluster-critical` | Cluster components that are critical to its correct operation. This PriorityClass is mandatory for MutatingWebhooks and Extension API servers. It also fully protects components against [eviction by kubelet](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/).<br>`kube-dns`, `coredns`, `kube-proxy`, `flannel`, `kube-api-server`, `kube-controller-manager`, `kube-scheduler`, `cluster-autoscaler`, `machine-controller-manager`. | 2000000000 |
| `production-high`         | Stateful applications in the production environment. Their unavailability leads to service downtime or data loss (postgresql, memcached, redis, mongo, etc.).                                                                                                                                                                                                                                                                                                | 9000       |
| `cluster-medium`          | Cluster components responsible for monitoring (alerts, diagnostic tools) and autoscaling. Monitoring tools help engineers assess the scale of incidents; autoscaling provides the necessary resources to applications.<br>`deckhouse`, `node-local-dns`, `kube-state-metrics`, `madison-proxy`, `node-exporter`, `trickster`, `grafana`, `kube-router`, `monitoring-ping`, `okmeter`, `smoke-mini`                                                           | 7000       |
| `production-medium`       | Main stateless applications in the production environment that are responsible for operating the service for end-users.                                                                                                                                                                                                                                                                                                                                      | 6000       |
| `deployment-machinery`    | Cluster components that are responsible for deploying/building (helm, werf).<br>`kube-system/tiller-deploy`                                                                                                                                                                                                                                                                                                                                                  | 5000       |
| `production-low`          | Non-critical, secondary applications in the production environment (crons, admin dashboards, batch processing). For important batch or cron jobs, consider assigning them the production-medium priority.                                                                                                                                                                                                                                                    | 4000       |
| `staging`                 | Staging environments for applications.                                                                                                                                                                                                                                                                                                                                                                                                                       | 3000       |
| `cluster-low`             | Cluster components that are desirable but not essential for proper cluster operation. <br>`prometheus-operator`, `dashboard`, `dashboard-oauth2-proxy`, `cert-manager`, `prometheus`, `prometheus-longterm`                                                                                                                                                                                                                                                  | 2000       |
| `develop` (default)       | Dev-environments for applications. The default class for a component (if other priority classes aren't set).                                                                                                                                                                                                                                                                                                                                                 | 1000       |
| `standby`                 | This class is not intended for applications. It is used for system purposes (reserving nodes).                                                                                                                                                                                                                                                                                                                                                               | -1         |