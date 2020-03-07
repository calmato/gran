#################################################
# Common
#################################################
variable "location" {
  default = "asia-northeast1-a"
}

variable "tags" {
  default = {}
}

#################################################
# GKE Cluster
#################################################
variable "gke_cluster_name" {
  description = "GKE クラスタ名"
  default     = ""
}

variable "gke_cluster_description" {
  description = "GKE クラスタ説明"
  default     = ""
}

variable "gke_cluster_min_master_version" {
  description = "GKE クラスタ最低バージョン"
  default     = "1.14.10-gke.17"
}

variable "gke_cluster_ipv4_cidr" {
  description = "GKE PodのCIDR"
  default     = null
}

#################################################
# GKE Node
#################################################
variable "gke_node_configs" {
  description = "GKE ノード設定"
  type = list(object({
    name         = string # ノード名
    count        = number # ノード数
    preemptible  = bool   # プリエンプティブの利用
    machine_type = string # マシンタイプ e.g.) f1-micro, n1-standard-1, etc..
  }))
}
