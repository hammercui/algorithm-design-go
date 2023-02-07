# algorithm-design-go
>golang算法，数据结构，工具库汇总
## 命名规范
#### 1 包名
必须是小写，可以允许下划线 比如net_work 
#### 2 类名
必须是小写，可以允许下划线 比如go_dialog.go ，必须是名词或动名词
#### 3 接口
但函数接口，以er后缀，整体必须是名词
#### 4 函数和方法名
动词或动词短语

## 设计模式
### 行为型
* chainOfResponsibility 责任链
* observer 观察者
* strategy 策略
* templateMethod 模板方法
* command 命令
* iterator 迭代器
* memento 备忘录
* state 状态机
* visitor 访问者
### 创造型
* abstractFactory 抽象工厂
* prototype 原型
* singleton 单例
* builder(函数式语言用的比较少) 构建者
### 函数式
* option 可选项
### 结构型
* composite 组合
* proxy 代理
* decorator 装饰器
* adapter 适配器
## 数据结构

+ list:链表
    - singleList:单向链表
    - skipList:跳表,某些场景可以取代hashMap
+ tree:树
    - binaryTree: 二叉树
    - binarySearchTree:二叉搜索树
    - 平衡二叉树（AVL）?
    - 红黑树?
+ queue:队列

## 算法
+ dp:dynamic programming动态规划问题
    - lis:最长上升子序列
    - maxArraySum:最大子序和    
    - uniquePath:最多路径问题? 
    
## 工具 tools
+ crypto:密码学
    - MD5?
    - SHA512?
    - RSA?
    - 随机
        - Mersenne twister:梅森旋转算法mt19937
+ time: 时间相关

## 语言特性 features
+ slice:切片
+ options: 可选配置项

## 乐扣 letcode
+ string_reverse