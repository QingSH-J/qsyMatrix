# qsyMatrix

<p align="center">
  <a href="https://pkg.go.dev/github.com/QingSH-J/qsyMatrix"><img src="https://pkg.go.dev/badge/github.com/QingSH-J/qsyMatrix.svg" alt="Go Reference"></a>
  <a href="https://github.com/QingSH-J/qsyMatrix/actions"><img src="https://github.com/QingSH-J/qsyMatrix/actions/workflows/go.yml/badge.svg" alt="Build Status"></a>
  <a href="https://goreportcard.com/report/github.com/QingSH-J/qsyMatrix"><img src="https://goreportcard.com/badge/github.com/QingSH-J/qsyMatrix" alt="Go Report Card"></a>
  <a href="https://github.com/QingSH-J/qsyMatrix/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License"></a>
</p>

一个使用 Go 泛型实现的、类型安全、高性能的现代化矩阵运算库。

---

* [✨ 特性](#-特性)
* [🚀 安装](#-安装)
* [💡 快速上手](#-快速上手)
* [📚 API 文档](#-api-文档)
* [🤝 贡献指南](#-贡献指南)
* [📄 许可证](#-许可证)

---

## ✨ 特性

* **泛型支持**: 可在编译时确定矩阵元素类型，支持 `int`, `float64` 等多种数字类型。
* **类型安全**: 借助 Go 泛型，完全杜绝因类型不匹配导致的运行时错误。
* **丰富的 API**:
    * **构造函数**: `New`, `Zeros` (零矩阵), `Eye` (单位矩阵)。
    * **基本运算**: 矩阵加法 (`Add`)、数乘 (`ScalarMultiply`)、矩阵乘法 (`Multiply`)。
    * **矩阵变换**: 转置 (`Transpose`)。
* **清晰的错误处理**: 为维度不匹配等常见错误提供了清晰的错误信息。
* **零依赖**: 除 Go 标准库外，不依赖任何第三方库。

---

## 🚀 安装

```bash
go get [github.com/QingSH-J/qsyMatrix](https://github.com/QingSH-J/qsyMatrix)