### 分支介绍

- (from) runtime0: 1.19.2版本初始源码
- (name) style/{title}: 仅包含title类型对应注释的改动分支
- (into) master: 所有改动注释的并集

### 编译相关

- [代码目录](../../src/cmd/compile)
- [编译器入口](../../src/cmd/compile/main.go)
- [解析&类型校验入口](../../src/cmd/compile/internal/noder/noder.go#LoadPackage)

#### 词法解析(A)

- [词法解析器](../../src/cmd/compile/internal/syntax/scanner.go#scanner)
- [Token](../../src/cmd/compile/internal/syntax/tokens.go)

> [词法解析测试入口](../../src/cmd/compile/internal/syntax/scanner_test.go#TestScanner)

#### 语法分析(A)

- [语法解析入口](../../src/cmd/compile/internal/syntax/syntax.go#Parse)
- [语法解析细节](../../src/cmd/compile/internal/syntax/parser.go#fileOrNil)
- [stmt](../../src/cmd/compile/internal/syntax/nodes.go)

> [语法解析测试入口](../../src/cmd/compile/internal/syntax/parser_test.go#TestVerify)

#### 类型检查(B)

- [类型检查入口](../../src/cmd/compile/internal/noder/irgen.go#checkFiles)
- [类型检查主程序](../../src/cmd/compile/internal/types2/check.go#Files)
  -- initFiles: 同目录包名校验
  -- collectObjects: 收集所有的文件和包级别对象(package,import,var,const,type,func/method)，填充到对应容器(impMap，objMap，scope，info(顶级声明的细分部分)）
  -- packageObjects: 完成包对象的类型检查，不包含func函数体，会讲func函数体注册到delayProcess中。
- [类型检查逻辑](../../src/cmd/compile/internal/types2/decl.go)
  -- 三色法处理类型依赖关系(检查循环依赖问题)
  -- 根据obj类型分别调用(constDecl，varDecl，typeDecl&collectMethods，funcDecl做对应的检查)

##### 顶级对象的类型检查


> [编译器调试入口](../../src/cmd/compile/main.go)
> 环境：Goland 2022，Golang 1.19.2
> debug设置入参：/Users/cengqi/workspace/go19/_kiqi/add.go /Users/cengqi/workspace/go19/_kiqi/main.go
> 开启tracing日志
