# summarize

指定したURLのWebページ本文を取得し、OpenAI GPT-4を使って「要約+日本語翻訳」してくれるCLIツールです。しかし、遅すぎる。普通にGPTを使った方が良い。
---

![image](スクリーンショット 2025-03-25 041352.png)

## 🔧 インストール手順

1. このリポジトリをクローン

```bash
git clone https://github.com/your-name/summarize.git
cd summarize
```

2. パッケージのインストール
```bash
go mod tidy
```

3.envファイルの作成し、APIキーを設定
```bash
OPENAI_API_KEY="your-api-key"
```

4. summarize.shに実行権限を付与し、実行
```bash
chmod +x summarize.sh
./summarize.sh
```

