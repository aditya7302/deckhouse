# [<en>] Section containing the parameters of instance class for worker nodes.
# [<en>] https://deckhouse.io/documentation/v1/modules/030-cloud-provider-aws/cr.html
# [<ru>] Секция, описывающая параметры инстанс-класса для узлов c компонентами, обеспечивающими рабочую нагрузку.
# [<ru>] https://deckhouse.ru/documentation/v1/modules/030-cloud-provider-aws/cr.html
apiVersion: deckhouse.io/v1
kind: AWSInstanceClass
metadata:
  name: worker
spec:
  # [<en>] VM disk size.
  # [<ru>] Размер диска для виртуальной машины.
  # [<en>] You might consider changing this.
  # [<ru>] Возможно, захотите изменить.
  diskSizeGb: 30
  # [<en>] VM disk type to use.
  # [<ru>] Используемый тип диска для виртуальной машины.
  # [<en>] You might consider changing this.
  # [<ru>] Возможно, захотите изменить.
  diskType: gp3
  # [<en>] Type of the instance.
  # [<ru>] Тип используемого инстанса.
  # [<en>] You might consider changing this.
  # [<ru>] Возможно, захотите изменить.
  instanceType: c5.xlarge
---
# [<en>] Section containing the parameters of worker node group.
# [<en>] https://deckhouse.io/documentation/v1/modules/040-node-manager/cr.html#nodegroup
# [<ru>] Секция, описывающая параметры группы узлов c компонентами, обеспечивающими рабочую нагрузку.
# [<ru>] https://deckhouse.ru/documentation/v1/modules/040-node-manager/cr.html#nodegroup
apiVersion: deckhouse.io/v1
kind: NodeGroup
metadata:
  name: worker
spec:
  cloudInstances:
    classReference:
      kind: AWSInstanceClass
      name: worker
    # [<en>] The maximum number of instances for the group in each zone (used by the autoscaler).
    # [<en>] You might consider changing this.
    # [<ru>] Максимальное количество инстансов в каждой зоне (используется при масштабировании).
    # [<ru>] Возможно, захотите изменить.
    maxPerZone: 1
    # [<en>] The minimum number of instances for the group in each zone.
    # [<ru>] Минимальное количество инстансов в каждой зоне.
    minPerZone: 1
    # [<en>] List of availability zones to create instances in.
    # [<ru>] Список зон, в которых создаются инстансы.
    # [<en>] You might consider changing this.
    # [<ru>] Возможно, захотите изменить.
    zones:
      - eu-central-1a
  disruptions:
    approvalMode: Automatic
  nodeType: CloudEphemeral
---
# [<en>] Section containing the parameters of NGINX Ingress controller.
# [<en>] https://deckhouse.io/documentation/v1/modules/402-ingress-nginx/cr.html
# [<ru>] Секция, описывающая параметры NGINX Ingress controller.
# [<ru>] https://deckhouse.ru/documentation/v1/modules/402-ingress-nginx/cr.html
apiVersion: deckhouse.io/v1
kind: IngressNginxController
metadata:
  name: nginx
spec:
  ingressClass: nginx
  inlet: LoadBalancer
  loadBalancer:
    annotations:
      service.beta.kubernetes.io/aws-load-balancer-type: nlb
  # [<en>] Describes on which nodes the Ingress Controller will be located. Label node.deckhouse.io/group: <NAME_GROUP_NAME> is set automatically.
  # [<ru>] Описывает, на каких узлах будет находиться Ingress-контроллер. Лейбл node.deckhouse.io/group: <NODE_GROUP_NAME> устанавливается автоматически.
  nodeSelector:
    node.deckhouse.io/group: worker
---
# [<en>] RBAC and authorization settings.
# [<en>] https://deckhouse.io/documentation/v1/modules/140-user-authz/cr.html#clusterauthorizationrule
# [<ru>] Настройки RBAC и авторизации.
# [<ru>] https://deckhouse.ru/documentation/v1/modules/140-user-authz/cr.html#clusterauthorizationrule
apiVersion: deckhouse.io/v1
kind: ClusterAuthorizationRule
metadata:
  name: admin
spec:
  subjects:
  - kind: User
    name: admin@deckhouse.io
  accessLevel: SuperAdmin
  portForwarding: true
---
# [<en>] Parameters of the static user.
# [<en>] https://deckhouse.io/documentation/v1/modules/150-user-authn/cr.html#user
# [<ru>] Данные статического пользователя.
# [<ru>] https://deckhouse.ru/documentation/v1/modules/150-user-authn/cr.html#user
apiVersion: deckhouse.io/v1
kind: User
metadata:
  name: admin
spec:
  # [<en>] User e-mail.
  # [<ru>] E-mail пользователя.
  email: admin@deckhouse.io
  # [<en>] This is a hash of the password <GENERATED_PASSWORD>, generated when loading the page of the Getting Started.
  # [<en>] Generate your own or use it at your own risk (for testing purposes)
  # [<en>] echo "<GENERATED_PASSWORD>" | htpasswd -BinC 10 "" | cut -d: -f2
  # [<ru>] Это хэш пароля <GENERATED_PASSWORD>, сгенерированного при загрузке страницы "Быстрого Старта".
  # [<ru>] Сгенерируйте свой или используйте этот, но только для тестирования
  # [<ru>] echo "<GENERATED_PASSWORD>" | htpasswd -BinC 10 "" | cut -d: -f2
  # [<en>] You might consider changing this.
  # [<ru>] Возможно, захотите изменить.
  password: <GENERATED_PASSWORD_HASH>
