[![codecov](https://codecov.io/gh/rai-wtnb/accomplist-api/branch/main/graph/badge.svg?token=I3MLDYFD21)](https://codecov.io/gh/rai-wtnb/accomplist-api)
# リンク
[AccompList](https://accomplist.work/)
ログイン画面・新規登録画面以下に「テストユーザーでログイン」というボタンがあり、ユーザー登録なしにお試しいただけます。
# アプリケーション概要
### 目標・やりたいことを実現するためのモチベーションツールとして、頑張る人たちと繋がるツールとして、役立つアプリケーションです
AccompList = Accomplish + List

このアプリケーションで実現したかったこと
- 目標実現のためのモチベーション維持の手助けをすること
  - 目標をAccompListとして書き出すことで目標を明確にします。
  - 達成の記録を書くことでフィードバックを明確にします。
  - また他のユーザの達成の記録をみることで、自分でもまだ気付いていないやりたいことに気付けるかもしれないし、モチベーション維持につながるだろうという理由で、達成の記録の一覧を閲覧できるようにしました。
- 何かに励む人たちとつながることができる
  - 頑張る方達とつながることで、交流を生んだりモチベーションアップにつなげたいという目的からいいね機能やフォロー機能を考えました。

また既存のSNSは、ユーザの目的ごとに自由に使用できます。SNSの形として、全員が同じ目的(目標実現・モチベーションアップ)を持って使用するアプリケーションを作ることができないか、という考えもありました。
# 機能一覧
- S3への画像アップロード
  - プロフィール画像, 達成の記録時に登録する画像をS3へアップロード
- プロフィール
  - 名前, ツイッターアカウント, 自己紹介コメント, プロフィール画像を登録可能
- AccompListの記入
  - リストの投稿, 削除
- 達成の記録を記入
  - リスト達成時に達成した記録を画像付きで登録可能
- 検索機能
  - 関連するユーザー、達成の記録を検索可能　→　つながるきっかけに
- フォロー機能

# 使用技術
## バックエンド
- Go: 1.15.2
- Gin: 1.6.3
- gorm: 1.9.16
- aws-sdk-go: 1.35.35
## フロントエンド
[accomplist-client](https://github.com/rai-wtnb/accomplist-client)
- TypeScript: 4.0.3
- React: 16.13.1
- Next.js: 9.5.4
- Tailwindcss: 1.9.1
## インフラ
- 構成図
![accomplist-infra](https://user-images.githubusercontent.com/55418247/102662879-1021a280-41c3-11eb-8d1c-5071b954c4a8.png)
- AWS
  - ECS / ECR / ALB / EC2 / VPC / RDS(PostgreSQL) / S3 / Route53 / ACM / CloudWatch / SSM
- Terraform: 0.12.5
  - IaCに取り組みました
- Docker
  - Docker: 20.10.0
    - ボリュームでのコンテナ間データ共有
  - docker-compose: 1.27.4
    - ローカル開発環境
    - docker networkを使用したAPI,クライアント間のやりとり
- CircleCI/CD
    - 自動テスト
    - 自動デプロイ
