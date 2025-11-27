resource "google_container_cluster" "primary" {
  name     = "minibank-cluster"
  location = "asia-northeast1"

  network    = google_compute_network.vpc.name
  subnetwork = google_compute_subnetwork.subnet.name

  enable_autopilot = true

  deletion_protection = false
}