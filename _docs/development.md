# 開発メモ

# 技術的な決定とその理由

## protoファイルをどこに置くか

前提: monorepo

スキーマ定義であるprotoファイルは通信する2つのサービスで共有される。
また、Bufでは `buf.yaml` が存在するディレクトリがルートとして命名規則が扱われる。
`buf.gen.yaml` は出力先を指定するため、各サービスのリポジトリに存在していれば良い。

以上から、ディレクトリ構成を以下のようにしてみる。

```text
.
├── api/proto/
    ├── buf.yaml
    ├── proto/
        ├── user/
            ├── v1/
                ├── user.proto
        ├── greet/
            ├── v1/
                ├── greet.proto
├── server/
    ├── buf.gen.yaml
    ├── go.mod
├── client/
    ├── buf.gen.yaml
    ├── src/
```
