job "nomad-api-proxy" {
  region      = "global"
  datacenters = ["dc1"]
  type        = "system"

  group "nomad-api-proxy" {
    count = 1

    network {
      mode = "bridge"

      port "api" {
        static = 4649
        to     = 4649
      }

    }

    service {
      name = "nomad-api-proxy"

//      check {
//        name     = "alive"
//        type     = "tcp"
//        port     = "http"
//        interval = "10s"
//        timeout  = "2s"
//      }
    }

    task "nomad-api-proxy" {
      driver = "docker"

      env {
        NOMAD_ADDR = "${attr.unique.network.ip-address}"
      }

      config {
        image        = "docker.io/library/nomad-api-proxy"
        network_mode = "host"
      }

      resources {
        cpu    = 100
        memory = 128
      }
    }
  }
}