# 🏦 Cloud Native Banking Infrastructure (Mini Bank)

![Google Cloud](https://img.shields.io/badge/Google_Cloud-%234285F4.svg?style=for-the-badge&logo=google-cloud&logoColor=white)
![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)
![Terraform](https://img.shields.io/badge/terraform-%235835CC.svg?style=for-the-badge&logo=terraform&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

## 📖 プロジェクト概要 (Overview)
本プロジェクトは、**Google Cloud Platform (GCP)** を活用して構築した、**クラウドネイティブなバンキング・インフラストラクチャ**のポートフォリオです。

金融システムに求められる「高可用性」と「拡張性」を意識し、**GKE (Google Kubernetes Engine)** を中心としたマイクロサービスアーキテクチャを採用しました。また、すべてのインフラリソースは **Terraform** を用いてコード化 (IaC) し、再現性と管理効率を高めています。



## 📸 Infrastructure & Deployment Proof (デモ・証跡)

Cloud resources are currently managed via Terraform. Below are the evidences of successful deployment.
*(コスト最適化のためリソースは停止中です。以下は構築時の証跡スクリーンショットです。)*

### 1. Web Application (On GKE)
!<img width="1470" height="923" alt="Image" src="https://github.com/user-attachments/assets/882e71ca-2633-4d84-9b29-5ceb69777f3b" />
<img width="1470" height="923" alt="Image" src="https://github.com/user-attachments/assets/72de5f6d-40e0-475a-a30d-9d640995526b" />
> Actual running instance on Google Kubernetes Engine with Japanese UI.

### 2. Infrastructure Provisioning (Terraform)
![Terraform Apply](file:///Users/bin/Library/Mobile%20Documents/com~apple~CloudDocs/port/%E1%84%89%E1%85%B3%E1%84%8F%E1%85%B3%E1%84%85%E1%85%B5%E1%86%AB%E1%84%89%E1%85%A3%E1%86%BA%202025-11-27%2015.49.46.png)
> **"Apply complete!"**: All infrastructure components (VPC, GKE) were automatically provisioned using Terraform.

### 3. Kubernetes Pod Status & Troubleshooting
![Kubectl Status](file:///Users/bin/Library/Mobile%20Documents/com~apple~CloudDocs/port/%E1%84%89%E1%85%B3%E1%84%8F%E1%85%B3%E1%84%85%E1%85%B5%E1%86%AB%E1%84%89%E1%85%A3%E1%86%BA%202025-11-27%2015.49.14.png)
> **Problem Solving**: Diagnosed `InvalidImageName` error during initial deployment and successfully transitioned pods to `Running` state by fixing the image tag manifest.



## 🏗 アーキテクチャ (Architecture)

Google Cloud 上に構築されたインフラ構成図です。

graph LR
    User((User)) -->|HTTPS| LB[Cloud Load Balancer]
    subgraph "Google Cloud Platform (Tokyo Region)"
        subgraph "VPC Network"
            subgraph "GKE Cluster (Autopilot)"
                LB --> Pod1[Backend App (Go)]
                LB --> Pod2[Backend App (Go)]
            end
        end
    end

    🛠 技術スタック (Tech Stack)
Cloud Provider: Google Cloud Platform (GCP)

Orchestration: Google Kubernetes Engine (GKE Autopilot)

Infrastructure as Code: Terraform

Backend: Go (Golang) - 軽量かつ高速なAPI処理

Frontend: Vanilla JS

Containerization: Docker, Google Artifact Registry



🚀 実装のポイント (Key Features)
1. Terraformによるインフラのコード化 (IaC)
手動操作 (ClickOps) によるミスを排除するため、VPCネットワーク構築からGKEクラスタのプロビジョニングまで、インフラ構成をすべて Terraform で定義しました。

モジュール構成: provider.tf, network.tf, gke.tf に分割し、保守性を向上。

2. GKEによるマイクロサービス運用
アプリケーションをDockerコンテナ化し、Kubernetes 上にデプロイしました。

Deployment: 複数のレプリカ (Replicas) を稼働させ、耐障害性を確保。

Service (LoadBalancer): 外部トラフィックを適切に分散。

3. クロスプラットフォーム・ビルド戦略 (トラブルシューティング)
ローカル開発環境 (Apple Silicon/Mシリーズ) とクラウド本番環境 (Linux/AMD64) のアーキテクチャの違いによるデプロイエラーを解決しました。

課題: ARM64アーキテクチャでビルドされたイメージが、GKE (AMD64) 上で exec format error を引き起こした。

解決: Docker Buildx を活用し、--platform linux/amd64 フラグを用いてクラウド環境互換のイメージをビルド・デプロイするパイプラインを確立。



💻 ローカルでの実行方法 (How to Run)
Docker環境があれば、以下のコマンドでローカルでも動作確認が可能です。

Bash

# 1. イメージのビルド
docker build -t minibank-portfolio .

# 2. コンテナの起動
docker run -p 8080:8080 minibank-portfolio
ブラウザで http://localhost:8080 にアクセスしてください。



🔮 今後の展望 (Future Roadmap)
[ ] データベース統合: Cloud Spanner (または Cloud SQL) を導入し、トランザクションの整合性を担保。

[ ] イベント駆動アーキテクチャ: Cloud Pub/Sub と Cloud Functions (Python) を連携させ、高額送金時のリアルタイム通知機能を実装。

[ ] GitOpsの導入: ArgoCDを活用した継続的デリバリー (CD) パイプラインの構築。



👨‍💻 作成者 (Author)
Role: Cloud Infrastructure Engineer / Backend Developer

Focus: GCP, Kubernetes, Terraform, Go
