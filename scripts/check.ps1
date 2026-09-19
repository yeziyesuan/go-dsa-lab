# W1 一键自检脚本（在有 Go 的机器上跑）
#
# 用法：在仓库根目录执行
#     powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\check.ps1
#
# 为什么要加 -ExecutionPolicy Bypass：Windows 默认禁止运行 .ps1 脚本（报
# "running scripts is disabled on this system"），加这个参数只对这一次运行生效，
# 不改你系统的全局设置。
#
# 它会依次做 5 件事，每一步失败都会明确告诉你「哪一步、怎么修」。
# 本仓库已经用 go1.22.12 通过过这 5 步（见 docs\验证记录.md），
# 所以这个脚本主要是帮你确认「你机器上的环境」是否正确。

$ErrorActionPreference = 'Continue'
$fail = 0

function Step($n, $title) {
    Write-Host ""
    Write-Host "==== [$n] $title ====" -ForegroundColor Cyan
}

function Ok($msg)   { Write-Host "  ✓ $msg" -ForegroundColor Green }
function Bad($msg)  { Write-Host "  ✗ $msg" -ForegroundColor Red; $script:fail++ }
function Hint($msg) { Write-Host "    → $msg" -ForegroundColor Yellow }

# 必须在仓库根目录：能同时看到 go.mod 和 syntax/
if (-not (Test-Path '.\go.mod') -or -not (Test-Path '.\syntax')) {
    Write-Host "请先 cd 到 go-dsa-lab 仓库根目录再运行本脚本。" -ForegroundColor Red
    exit 1
}

# ---------- 1. 环境 ----------
Step 1 "环境：go / git 是否可用"

$go = Get-Command go -ErrorAction SilentlyContinue
if ($go) {
    Ok "go 路径：$($go.Source)"
    Ok ("版本：" + (& go version))
    $proxy = (& go env GOPROXY)
    if ($proxy -match 'goproxy\.cn|goproxy\.io|mirrors') {
        Ok "GOPROXY = $proxy"
    } else {
        Hint "GOPROXY = $proxy（国内建议：go env -w GOPROXY=https://goproxy.cn,direct）"
    }
} else {
    Bad "找不到 go 命令"
    Hint "winget install --id GoLang.Go -e   然后**重开终端**；或到 https://go.dev/dl/ 下载 msi 双击安装"
    Hint "装完再来跑本脚本。下面的步骤全部跳过。"
    exit 1
}

if (Get-Command git -ErrorAction SilentlyContinue) {
    Ok ("git：" + (& git --version))
} else {
    Hint "没有 git，本步不影响编译：winget install --id Git.Git -e"
}

# ---------- 2. 格式化 ----------
Step 2 "gofmt：有没有格式不规范的 Go 文件（有输出不算致命，但应该修）"
$unformatted = & gofmt -l . 2>&1
if ($unformatted) {
    Bad "以下文件未格式化："
    $unformatted | ForEach-Object { Write-Host "    $_" }
    Hint "一条命令修好：gofmt -w ."
} else {
    Ok "全部文件格式规范"
}

# ---------- 3. 静态检查 ----------
Step 3 "go vet：编译前能查出来的错（未使用变量、动词不匹配、结构体标签等）"
& go vet ./... 2>&1 | ForEach-Object { Write-Host "    $_" }
if ($LASTEXITCODE -eq 0) { Ok "go vet 通过" } else { Bad "go vet 报错，先修它，测试结果没有意义" }

# ---------- 4. 测试 ----------
Step 4 "go test：主包 + 20 道题的表驱动测试"
if (-not (Test-Path '.\go.sum')) {
    Hint "还没有 go.sum（本仓库只用标准库，正常），若报 missing go.sum entry 就跑 go mod tidy"
}
& go test ./... 2>&1 | ForEach-Object { Write-Host "    $_" }
if ($LASTEXITCODE -eq 0) { Ok "全部测试通过" } else { Bad "有测试失败，看上面的输出定位到具体题号" }

# ---------- 5. 12 个语法示例逐个跑 ----------
Step 5 "go run：12 个语法示例能否跑起来（只看退出码）"
$dirs = Get-ChildItem -Path '.\syntax' -Directory | Sort-Object Name
foreach ($d in $dirs) {
    $out = & go run "./syntax/$($d.Name)" 2>&1
    if ($LASTEXITCODE -eq 0) {
        Ok "$($d.Name)  首行输出：$(($out | Select-Object -First 1))"
    } else {
        Bad "$($d.Name) 运行失败：$(($out | Select-Object -Last 1))"
    }
}

# ---------- 汇总 ----------
Write-Host ""
if ($fail -eq 0) {
    Write-Host "全部通过 ✅  接着做：git add -A; git commit -m 'W1: 环境 + 语法示例 + 前 20 题'" -ForegroundColor Green
    Write-Host "然后照着 docs\W1-使用说明.md 第 5 节逐条勾 W1 验收清单。" -ForegroundColor Green
} else {
    Write-Host "有 $fail 项没通过 ❌  先修红色那几项，再重跑本脚本。" -ForegroundColor Red
}
exit $fail
