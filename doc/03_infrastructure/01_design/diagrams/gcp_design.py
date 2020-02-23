from diagrams import Diagram, Cluster
from diagrams.gcp.database import Firestore
from diagrams.gcp.network import CDN, DNS, LoadBalancing
from diagrams.gcp.storage import GCS
from diagrams.gcp.compute import GKE
from diagrams.onprem.compute import Server
from diagrams.oci.network import Internetgateway

with Diagram("Gran - GCP Design", show=False):
	server = Server("Client (Nuxt.js)")

	with Cluster("Google Cloud Platform"):
		firestore = Firestore("Firebase Firestore")
		cdn = CDN("Cloud CDN")
		dns = DNS("Cloud DNS")
		lb = LoadBalancing("Cloud Load Balancing")
		gcs = GCS("Cloud Storage")
		gke = GKE("Kubernetes Engine")
		net = Internetgateway("Gateway")

		dns >> cdn >> gcs
		dns >> lb >> gke >> firestore

	server >> net
	net >> dns
