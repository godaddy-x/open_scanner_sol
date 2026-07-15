# open_scanner_sol

SOL 链扫块进程，依赖 [open_scanner](../open_scanner) 框架，通过 [wallet-adapter-sol](../../github/wallet-adapter-sol) 注入链适配器。

## 启动

```bash
go run . -config resource/config_ops.yaml -mainsymbol SOL -symbol SOL
```

可选：`-initheight <slot>` 指定起始 slot。

## 结构

```
open_scanner_sol/
├── main.go              # sol.NewSolAdapter 注入 + plugin 侧载
├── plugin/register.go   # RegisterSignVerify、流水分类、handlers
├── handlers/            # SOL 专用 RPC（v1 暂无 SpeedUp/Batch）
└── resource/            # 本地 config_ops.yaml（gitignore，参考 open_scanner_eth）
```

```go
func newAdapter(data *model.OwSymbol) adapter.ChainAdapter {
    return sol.NewSolAdapter(data.Symbol, data.Name, int32(data.Decimals))
}
```

## 能力边界（相对 ETH/TRX）

| 能力 | SOL |
|------|-----|
| 主币 / SPL 建单广播 | 走 open_scanner 通用 CreateTrade / Summary |
| 扫块入账 | wallet-adapter-sol BlockScanner |
| MPC 签前 | `sol.RegisterSignVerify`（solana-ed25519） |
| SpeedUp / Cancel | 暂无（无 nonce / RBF；依赖 blockhash 过期重发） |
| Batch / Deploy | 暂无 |

## 构建

```bash
chmod +x build_release.sh && ./build_release.sh
```

产物：`open-scanner-sol-linux-amd64` 等五平台二进制。

## 本地 replace

`go.mod` 默认指向本地：

```
replace open_scanner => ../open_scanner
replace wallet-adapter-sol => ../../github/wallet-adapter-sol
replace wallet-adapter => ../../github/wallet-adapter
replace freego => ../../github/freego
```

发版前改为模块版本 require，并注释 replace。

## 配置

将 `open_scanner_eth` / `open_scanner_trx` 的 `resource/config_ops.yaml` 拷到本仓 `resource/`，确保 Mongo 中有：

- `OwSymbol`：`Symbol=SOL`，`MainSymbol=SOL`，`Curve=2`，`Decimals=9`
- `OwSymbolNode.NodeConfig` JSON：`serverAPI`（Solana RPC）、`commitment`、`broadcastAPI` 等
