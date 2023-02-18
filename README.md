# engra

一本由所有学习者共同编写的英语词根词缀词典

## Graphql

```graphql endpoint doc
http://localhost:8081/
```

```graphql
{
    vocabularies(name: "weight") {
        ...vocabularyWithChildren
    }
}

fragment vocabulary on Vocabulary {
    name
}

fragment vocabularyWithChildren on Vocabulary {
    ...vocabulary
    children {
        ...vocabulary
        children {
            ...vocabulary
            children {
                ...vocabulary
                children {
                    ...vocabulary
                }
            }
        }
    }
}
```

## 前端

提供接口和组件两种形式，组件默认展示在官方网站上。

形式为：搜索框，搜索任意单词，出现该次的构词及词根树

## 数据

数据以词根为单元采用yaml文件存储在dicts中。首批词根来源于《英语词汇的奥秘》升级版

## 迭代计划

- [x] 建立词根库
- [x] 从公开渠道爬取单词词根的数据并补充到词根库中
- [ ] 建立分级库及词频库，根据分级和词频进行优先级排序
- [ ] 按照优先级检查是否已经在词根库中，如果不再则人工补充
- [ ] 长期的人工补充工作（在前端页面查不到时引导到创建issue及PR）