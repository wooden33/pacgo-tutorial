# Step 01
本次将学习到：
* 从文件中读取信息
* 打印到标准输出
* 处理多返回值
* 处理error
* 切片的处理
* 切片的遍历
* defer的使用
* 记载error

## 总览
```
- # represents a wall
- . represents a dot
- P represents the player
- G represents the ghosts (enemies)
- X represents the power up pills
```

# Step 02: 处理玩家的输入
本次将学习到：
* 不同的终端模式
* 在Go中调用外部命令
* 向控制台告知退出信息
* 创建函数返回多个值

## 终端模式介绍

终端可以运行在三种可能的模式下：

* Cooked 模式
* Cbreak 模式
* Raw 模式

Cooked 模式是我们通常使用的模式。在这种模式下，终端接收到的每一个输入都会被预处理，这意味着系统会截取特殊字符并赋予它们特殊的含义。

注意：特殊字符包括退格键（Backspace）、删除键（Delete）、Ctrl+D、Ctrl+C、箭头键等等。

Raw 模式则相反：数据会原样传递，不进行任何预处理。

Cbreak 模式则介于两者之间。某些字符会被预处理，而某些则不会。例如，Ctrl+C 仍然会导致程序中止，但箭头键会原样传递给程序。

我们将使用 cbreak 模式，以便我们可以处理与逃逸键和箭头键对应的转义序列。

# Step 03: 创建动作

创建结构体sprite，并且接受键盘动作

# Step 04: 增加了鬼魂的移动
