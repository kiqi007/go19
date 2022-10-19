### 分支介绍

- (from) runtime0: 1.19.2版本初始源码
- (name) style/{title}: 仅包含title类型对应注释的改动分支
- (into) master: 所有改动注释的并集

### 编译相关
- [代码目录](../../src/cmd/compile)
- [编译器入口](../../src/cmd/compile/internal/gc/main.go#Main)
- [解析&类型校验入口](../../src/cmd/compile/internal/noder/noder.go#LoadPackage)
#### 词法解析
- [词法解析器](../../src/cmd/compile/internal/syntax/scanner.go#scanner)
- [Token](../../src/cmd/compile/internal/syntax/tokens.go)
> [词法解析测试入口](../../src/cmd/compile/internal/syntax/scanner_test.go#TestScanner) 
#### 语法分析
- [语法解析入口](../../src/cmd/compile/internal/syntax/syntax.go:67#Parse)
- [语法解析细节](../../src/cmd/compile/internal/syntax/parser.go#fileOrNil)
- [stmt](../../src/cmd/compile/internal/syntax/nodes.go)
