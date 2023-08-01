---
title: "Модуль priority-class"
---

Модуль создает в кластере набор [priority class'ов](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass) и проставляет их компонентам установленным Deckhouse и приложениям в кластере.

[Priority Class](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption) — это функционал планировщика (scheduler'а), который позволяет учитывать приоритет Pod'а (из его принадлежности к классу) при планировании.

К примеру, при выкате в кластер Pod'ов с `priorityClassName: production-low`, если в кластере не будет доступных ресурсов для данного Pod'а, то Kubernetes начнет вытеснять Pod'ы с наименьшим приоритетом в кластере.
Т.е. сначала будут выгнаны все Pod'ы с `priorityClassName: develop`, потом — с `cluster-low`, и так далее.

При выставлении priority class очень важно понимать, к какому типу относится приложение и в каком окружении оно будет работать. Любой установленный `priorityClassName` не уменьшит приоритета Pod'а, т.к. если `priority-class` у Pod'а не установлен, планировщик считает его самым низким — `develop`. Очень важно правильно выставлять `priorityClassName`.

> **Внимание!** Нельзя использовать PriorityClasses: `system-node-critical`, `system-cluster-critical`, `cluster-medium`, `cluster-low`.

Устанавливаемые модулем priority class'ы (в порядке приоритета от большего к меньшему):

| Priority Class            | Описание                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Значение   |
|---------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------|
| `system-node-critical`    | Компоненты кластера, которые обязаны присутствовать на узле. Также полностью защищает от [вытеснения kubelet'ом](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/).<br>`node-exporter`, `csi`                                                                                                                                                                                                                                                   | 2000001000 |
| `system-cluster-critical` | Компоненты кластера, без которых его корректная работа невозможна. Этим PriorityClass в обязательном порядке помечаются MutatingWebhooks и Extension API servers. Также полностью защищает от [вытеснения kubelet'ом](https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource/).<br>`kube-dns`, `coredns`, `kube-proxy`, `flannel`, `kube-api-server`, `kube-controller-manager`, `kube-scheduler`, `cluster-autoscaler`, `machine-controller-manager`. | 2000000000 |
| `production-high`         | Stateful-приложения, отсутствие которых в production-окружении приводит к полной недоступности сервиса или потере данных (PostgreSQL, memcached, Redis, Mongo, ...).                                                                                                                                                                                                                                                                                                | 9000       |
| `cluster-medium`          | Компоненты кластера, влияющие на мониторинг (алерты, диагностика) кластера и автомасштабирование. Без мониторинга мы не можем оценить масштабы происшествия, без автомасштабирования — не сможем дать приложениям необходимые ресурсы.<br>`deckhouse`, `node-local-dns`, `kube-state-metrics`, `madison-proxy`, `node-exporter`, `trickster`, `grafana`, `kube-router`, `monitoring-ping`, `okmeter`, `smoke-mini`                                                  | 7000       |
| `production-medium`       | Основные stateless-приложения в production-окружении, которые отвечают за работу сервиса для посетителей.                                                                                                                                                                                                                                                                                                                                                           | 6000       |
| `deployment-machinery`    | Компоненты кластера, с помощью которых происходит сборка и деплой в кластер (Helm, werf).<br>`kube-system/tiller-deploy`                                                                                                                                                                                                                                                                                                                                            | 5000       |
| `production-low`          | Приложения в production-окружении (cron'ы, админки, batch-процессинг, ...), без которых можно прожить некоторое время. Если batch или cron никак нельзя прерывать, то он должен быть в production-medium, а не здесь.                                                                                                                                                                                                                                               | 4000       |
| `staging`                 | Staging-окружения для приложений.                                                                                                                                                                                                                                                                                                                                                                                                                                   | 3000       |
| `cluster-low`             | Компоненты кластера, без которых возможна эксплуатация кластера, но которые желательны. <br>`prometheus-operator`, `dashboard`, `dashboard-oauth2-proxy`, `cert-manager`, `prometheus`, `prometheus-longterm`                                                                                                                                                                                                                                                       | 2000       |
| `develop` (default)       | Develop-окружения для приложений. Класс по умолчанию, если не проставлены иные классы.                                                                                                                                                                                                                                                                                                                                                                              | 1000       |
| `standby`                 | Этот класс не предназначен для приложений. Используется в системных целях для резервирования узлов.                                                                                                                                                                                                                                                                                                                                                                 | -1         |