resource "google_compute_network" "vpc" {
  name                    = "minibank-vpc"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet" {
  name          = "minibank-subnet"
  ip_cidr_range = "10.0.0.0/24" # 10.0.0.0 ~ 10.0.0.255
  region        = "asia-northeast1"
  network       = google_compute_network.vpc.id
}