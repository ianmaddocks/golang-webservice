#webapp.tf

provider "kubernetes" {
  config_path = pathexpand(var.kind_cluster_config_path)
  alias = "alias"
}

variable "cloudflare_zone_id" {
    type = string
    sensitive = true
}

resource "kubernetes_deployment" "webapp1_deployment" {
  depends_on = [helm_release.traefik]
  
  metadata {
    name = "webapp1-deploy"
    labels = {
      app = "webapp1"
    }
    namespace = "default"
  }
  spec {
    replicas = 5
    selector {
      match_labels = {
        app = "webapp1"
      }
    }
    min_ready_seconds   = "5"
    strategy {
        type            = "RollingUpdate"
        rolling_update {
          max_surge        = "1"
          max_unavailable  = "0"
        }
    }
    template {
      metadata {
        labels = {
           app = "webapp1"
        }
      }
      spec {
        container {
          image = "ianmaddocks/webapp1:latest"
          name  = "webapp1"
          port {
            container_port = 80
          }
          liveness_probe {
            http_get {
              path = "/healthz"
              port = 80

              http_header {
                name  = "X-Custom-Header"
                value = "Awesome"
              }
            }
            initial_delay_seconds = 3
            period_seconds        = 3
          }
        }
      }
    }
  }
}

resource "kubernetes_ingress_v1" "webapp1_ingress" {
  depends_on = [kubernetes_deployment.webapp1_deployment]

  metadata {
    name = "webapp1"
  }
  spec { 
    rule {
      http {
        path {
          backend {
            service {
              name = "webapp1-svc"
              port {
                number = 80
              }
            }
          }
          path = "/version"
        }
        path {
          backend {
            service {
              name = "webapp1-svc"
              port {
                number = 80
              }
            }
          }
          path = "/whoami"
        }
      }
    }
    tls {
      secret_name = "webapp1"
      hosts = ["civo.maddocks.name"]
    }
  }
}

resource "kubernetes_service_v1" "webapp1" {
  depends_on = [kubernetes_deployment.webapp1_deployment]

  metadata {
    name = "webapp1-svc"
  }
  spec {
    selector = {
      app = kubernetes_deployment.webapp1_deployment.metadata.0.labels.app
    }
    port {
      port  = 80
    }
    type = "ClusterIP"
  }
}

resource "kubectl_manifest" "webapp1-certificate" {
    depends_on = [time_sleep.wait_for_clusterissuer]

    yaml_body = <<YAML
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: webapp1
  namespace: default
spec:
  secretName: webapp1
  issuerRef:
    name: cloudflare-prod
    kind: ClusterIssuer
  dnsNames:
  - 'civo.maddocks.name'
YAML
}

resource "cloudflare_record" "clcreative-main-cluster" {
    zone_id = var.cloudflare_zone_id #"your-zone-id"
    name = "civo.maddocks.name"
    value =  data.civo_loadbalancer.traefik_lb.public_ip #the public IP
    type = "A"
    proxied = false
}

output "public_ip_addr" {
  value       = data.civo_loadbalancer.traefik_lb.public_ip
  description = "The public IP address of the civo server instance."
}
