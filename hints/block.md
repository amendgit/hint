1. block作为函数参数

```object-c
- (void)foo:(void(^)(void))myBlock;
```

2. typedef一个block

```object-c
typedef void(^myBlock)(void);
```

3. block作为property

```object-c
@property (nonatomic, copy) void (^myBlock)(void);
```