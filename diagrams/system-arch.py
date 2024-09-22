from diagrams import Diagram, Cluster
from diagrams.generic.blank import Blank
from diagrams.k8s.compute import Pod
from diagrams.generic.database import SQL
from diagrams.onprem.monitoring import Prometheus
from diagrams.onprem.monitoring import Grafana

with Diagram("System Architecture", show=False):
    # User Interface
    ui = Blank("Web Frontend (Next.js)")

    # API Gateway
    api_gateway = Blank("API Gateway (Kong)")

    # Microservices
    with Cluster("Microservices"):
        auth_service = Pod("Auth Service")
        donations_service = Pod("Donations Service")

    # Database
    db = SQL("PostgreSQL")

    # Blockchain
    blockchain = Blank("Ethereum Network")

    # Monitoring
    with Cluster("Monitoring"):
        prometheus = Prometheus("Prometheus")
        grafana = Grafana("Grafana")

    # Connections
    ui >> api_gateway
    api_gateway >> [auth_service, donations_service]
    auth_service >> db
    donations_service >> db
    donations_service >> blockchain
    prometheus >> grafana
