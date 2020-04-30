from diagrams import Diagram, Cluster
from diagrams.gcp.database import Firestore
from diagrams.onprem.network import Envoy
from diagrams.onprem.compute import Server
from diagrams.oci.network import Internetgateway
from diagrams.k8s.compute import Pod
from diagrams.k8s.network import Ingress, Service

with Diagram("Gran - API System - Design", show=False):
	igw = Internetgateway("Gateway")

	with Cluster("Google Cloud Platform"):
		with Cluster("Google Kubernetes Engine"):
			net = Ingress("calmato.work")
			pods = [
				Pod("user service"),
				Pod("group service"),
				Pod("todo service")
			]

			net >> Service("svc") >> Envoy("lb") >> pods


		with Cluster("Firebase"):
			auth = Server("firebase authentication")
			firestore = Firestore("firestore")

			pods >> auth
			pods >> firestore

	igw >> net
