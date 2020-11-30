# 作业解答

## 结论

应该Wrap这个error，然后抛给上层。第三种错误的处理方式对使用者更友好，解释如下

## dao 层

核心代码位置：\dao\moive.go文件中的FindMoiveByName函数的24行

```go
    if err := db.QueryRow(moiveSQL, name).Scan(&moive.Id, &moive.Name, &moive.ToStar); err != nil {
        // 1. 将错误消除
        return nil, nil
        // 2. 不wrap直接返回原始错误
        return nil, err
        // 3. 使用wrap后返回, 最后选择使用第三种
        return nil, fmt.Errorf("没有找到 《%s》 这个影片: %w", name, err)
    }
```

## service 层

### 使用第一种模式

```go
// FindMoiveByName 根据影片名搜索影片
func FindMoiveByName(name string) (*model.Moive, error) {
    moive, err := dao.FindMoiveByName(name)

    if err != nil {
        return nil, err
    }

    if (moive != nil) {
        moive.Name += "-宇宙超级无敌动作影片公司"
    }
    return moive, nil
}
```

### 使用第二种模式

```go
// FindMoiveByName 根据影片名搜索影片
func FindMoiveByName(name string) (*model.Moive, error) {
    moive, err := dao.FindMoiveByName(name)
    if err != nil {
        return nil, fmt.Printf("没有找到影片 %v", err)
    }
    moive.Name += "-宇宙超级无敌动作影片公司"
    return moive, nil
}
```

### 使用第三种模式

```go
// FindMoiveByName 根据影片名搜索影片
func FindMoiveByName(name string) (*model.Moive, error) {
    moive, err := dao.FindMoiveByName(name)
    if err != nil {
        return nil, err
    }
    moive.Name += "-宇宙超级无敌动作影片公司"
    return moive, nil
}
```

--------end-----------
